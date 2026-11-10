// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// World is the closed set of application classes plus, on demand, the JDK
// classes above them. JDK method bodies are never analysed: a JDK class is a
// source of signatures that an application class may override.
type World struct {
	Classes  map[string]*Class
	AppNames []string  // sorted
	Methods  []*Method // application methods by ID
	Missing  map[string]bool
	// Shadowed lists classes present in more than one input directory; the
	// first directory wins, as on a class path.
	Shadowed []string

	lib       map[string]*zip.File
	closers   []io.Closer
	subtypes  map[string][]*Class // concrete application subtypes, sorted by name
	ancestors map[string][]string // sorted, self included
	selects   map[string][]*Method
	callbacks map[string][]*Method
	jdkSigs   map[string][]string
	lambdas   map[string]*Class
	constants *constantFields
	deadCache map[*Method]*deadInfo
	shapes    map[*Method]map[int]*BranchShape
}

// NewWorld loads every class under dirs and indexes the JDK's jmods.
func NewWorld(dirs []string, jdkHome string) (*World, error) {
	w := &World{
		Classes:   map[string]*Class{},
		Missing:   map[string]bool{},
		lib:       map[string]*zip.File{},
		subtypes:  map[string][]*Class{},
		ancestors: map[string][]string{},
		selects:   map[string][]*Method{},
		callbacks: map[string][]*Method{},
		jdkSigs:   map[string][]string{},
		lambdas:   map[string]*Class{},
	}
	for i, d := range dirs {
		if err := w.loadDir(d, i); err != nil {
			return nil, err
		}
	}
	if err := w.indexJDK(jdkHome); err != nil {
		return nil, err
	}
	sort.Strings(w.AppNames)
	sort.Strings(w.Shadowed)
	for _, n := range w.AppNames {
		for _, m := range w.Classes[n].Methods {
			m.ID = len(w.Methods)
			w.Methods = append(w.Methods, m)
		}
	}
	for _, n := range w.AppNames {
		c := w.Classes[n]
		if c.IsAbstract() {
			continue
		}
		for _, a := range w.Ancestors(c) {
			w.subtypes[a] = append(w.subtypes[a], c)
		}
	}
	return w, nil
}

// Close releases the jmod files.
func (w *World) Close() {
	for _, c := range w.closers {
		c.Close()
	}
}

func (w *World) loadDir(dir string, origin int) error {
	var paths []string
	err := filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if d.Name() == "META-INF" {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(p, ".class") && d.Name() != "module-info.class" {
			paths = append(paths, p)
		}
		return nil
	})
	if err != nil {
		return err
	}
	sort.Strings(paths)
	for _, p := range paths {
		data, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		c, err := ParseClass(data)
		if err != nil {
			return fmt.Errorf("%s: %w", p, err)
		}
		if _, dup := w.Classes[c.Name]; dup {
			w.Shadowed = append(w.Shadowed, c.Name)
			continue
		}
		c.App = true
		c.Origin = origin
		w.Classes[c.Name] = c
		w.AppNames = append(w.AppNames, c.Name)
	}
	return nil
}

// indexJDK maps class names to jmod entries. A jmod is a zip after a four byte header.
func (w *World) indexJDK(home string) error {
	mods, err := filepath.Glob(filepath.Join(home, "jmods", "*.jmod"))
	if err != nil {
		return err
	}
	if len(mods) == 0 {
		return fmt.Errorf("no jmods under %s", home)
	}
	sort.Strings(mods)
	for _, path := range mods {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		w.closers = append(w.closers, f)
		info, err := f.Stat()
		if err != nil {
			return err
		}
		z, err := zip.NewReader(io.NewSectionReader(f, 4, info.Size()-4), info.Size()-4)
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		for _, e := range z.File {
			if strings.HasPrefix(e.Name, "classes/") && strings.HasSuffix(e.Name, ".class") {
				name := strings.TrimSuffix(strings.TrimPrefix(e.Name, "classes/"), ".class")
				if _, dup := w.lib[name]; !dup {
					w.lib[name] = e
				}
			}
		}
	}
	return nil
}

// Class returns a class by name, reading the JDK on demand. An unknown name
// yields a recorded placeholder with no supertypes; the caller is not told apart
// because the only consequence is that nothing is inherited through it.
func (w *World) Class(name string) (*Class, error) {
	if c, ok := w.Classes[name]; ok {
		return c, nil
	}
	if e, ok := w.lib[name]; ok {
		r, err := e.Open()
		if err != nil {
			return nil, err
		}
		defer r.Close()
		data, err := io.ReadAll(r)
		if err != nil {
			return nil, err
		}
		c, err := ParseClass(data)
		if err != nil {
			return nil, fmt.Errorf("jdk %s: %w", name, err)
		}
		w.Classes[name] = c
		return c, nil
	}
	c := &Class{Name: name, Missing: true, MethodBy: map[string]*Method{}, FieldBy: map[string]*Field{}}
	if name != "java/lang/Object" {
		c.Super = "java/lang/Object"
	}
	w.Classes[name] = c
	w.Missing[name] = true
	return c, nil
}

func (w *World) mustClass(name string) *Class {
	c, err := w.Class(name)
	if err != nil {
		// A corrupt JDK image cannot be worked around, and silently treating the
		// class as absent would make the analysis unsound.
		fatalf("%v", err)
	}
	return c
}

func (c *Class) supers() []string {
	out := make([]string, 0, len(c.Interfaces)+1)
	if c.Super != "" {
		out = append(out, c.Super)
	}
	return append(out, c.Interfaces...)
}

// Ancestors lists the class itself and every supertype, sorted.
func (w *World) Ancestors(c *Class) []string {
	if a, ok := w.ancestors[c.Name]; ok {
		return a
	}
	seen := map[string]bool{}
	var walk func(x *Class)
	walk = func(x *Class) {
		if seen[x.Name] {
			return
		}
		seen[x.Name] = true
		for _, s := range x.supers() {
			walk(w.mustClass(s))
		}
	}
	walk(c)
	out := make([]string, 0, len(seen))
	for n := range seen {
		out = append(out, n)
	}
	sort.Strings(out)
	w.ancestors[c.Name] = out
	return out
}

// ConcreteSubtypes lists instantiable application classes at or below a type.
func (w *World) ConcreteSubtypes(name string) []*Class { return w.subtypes[name] }

// ResolveMethod follows JVMS 5.4.3.3: the class chain, then superinterfaces.
func (w *World) ResolveMethod(class, sig string) *Method {
	for c := w.mustClass(class); ; c = w.mustClass(c.Super) {
		if m := c.MethodBy[sig]; m != nil {
			return m
		}
		if c.Super == "" {
			break
		}
	}
	best := w.interfaceCandidates(w.mustClass(class), sig)
	for _, m := range best {
		if !m.IsAbstract() {
			return m
		}
	}
	if len(best) > 0 {
		return best[0]
	}
	return nil
}

// interfaceCandidates finds the maximally specific superinterface methods.
func (w *World) interfaceCandidates(c *Class, sig string) []*Method {
	var found []*Method
	for _, n := range w.Ancestors(c) {
		a := w.mustClass(n)
		if !a.IsInterface() {
			continue
		}
		if m := a.MethodBy[sig]; m != nil && !m.IsPrivate() && !m.IsStatic() {
			found = append(found, m)
		}
	}
	var out []*Method
	for _, m := range found {
		overridden := false
		for _, o := range found {
			if o != m && contains(w.Ancestors(o.Class), m.Class.Name) {
				overridden = true
				break
			}
		}
		if !overridden {
			out = append(out, m)
		}
	}
	return out
}

func contains(sorted []string, s string) bool {
	i := sort.SearchStrings(sorted, s)
	return i < len(sorted) && sorted[i] == s
}

// Select follows JVMS 5.4.6 for an object of exact class k: the first
// non-abstract declaration up the class chain, otherwise the maximally specific
// non-abstract superinterface methods.
func (w *World) Select(k *Class, sig string) []*Method {
	key := k.Name + " " + sig
	if v, ok := w.selects[key]; ok {
		return v
	}
	var out []*Method
	for c := k; ; c = w.mustClass(c.Super) {
		if m := c.MethodBy[sig]; m != nil && !m.IsStatic() && !m.IsAbstract() {
			out = []*Method{m}
			break
		}
		if c.Super == "" {
			break
		}
	}
	if out == nil {
		for _, m := range w.interfaceCandidates(k, sig) {
			if !m.IsAbstract() {
				out = append(out, m)
			}
		}
	}
	w.selects[key] = out
	return out
}

// JDKSignatures lists the overridable instance methods the JDK supertypes of k declare.
func (w *World) JDKSignatures(k *Class) []string {
	if v, ok := w.jdkSigs[k.Name]; ok {
		return v
	}
	seen := map[string]bool{}
	for _, n := range w.Ancestors(k) {
		a := w.mustClass(n)
		if a.App || a.Lambda {
			continue
		}
		for _, m := range a.Methods {
			if m.IsStatic() || m.IsPrivate() || m.Access&accFinal != 0 || m.Name == "<init>" || m.Name == "<clinit>" {
				continue
			}
			seen[m.Sig()] = true
		}
	}
	out := make([]string, 0, len(seen))
	for s := range seen {
		out = append(out, s)
	}
	sort.Strings(out)
	w.jdkSigs[k.Name] = out
	return out
}

// Callbacks lists the application methods that the JDK can call on an object of
// exact class k: those whose exact signature a JDK supertype declares and which
// selection picks for k. The hierarchy decides, not a list of names.
func (w *World) Callbacks(k *Class) []*Method {
	if v, ok := w.callbacks[k.Name]; ok {
		return v
	}
	out := []*Method{}
	for _, sig := range w.JDKSignatures(k) {
		for _, m := range w.Select(k, sig) {
			if m.Class.App {
				out = append(out, m)
			}
		}
	}
	w.callbacks[k.Name] = out
	return out
}

// LambdaClass is the modelled runtime class of a lambda implementing ifaces.
func (w *World) LambdaClass(ifaces []string) *Class {
	name := "lambda$" + strings.Join(ifaces, "+")
	if c, ok := w.lambdas[name]; ok {
		return c
	}
	c := &Class{Name: name, Super: "java/lang/Object", Interfaces: ifaces, Lambda: true, MethodBy: map[string]*Method{}, FieldBy: map[string]*Field{}}
	w.lambdas[name] = c
	w.Classes[name] = c
	return c
}

// ResolveField follows JVMS 5.4.3.2: the class, its superinterfaces, its superclass.
func (w *World) ResolveField(class, nameDesc string) *Field {
	var walk func(c *Class) *Field
	walk = func(c *Class) *Field {
		if f := c.FieldBy[nameDesc]; f != nil {
			return f
		}
		for _, i := range c.Interfaces {
			if f := walk(w.mustClass(i)); f != nil {
				return f
			}
		}
		if c.Super != "" {
			return walk(w.mustClass(c.Super))
		}
		return nil
	}
	return walk(w.mustClass(class))
}

// MethodByRef finds an application method named "class.name:desc".
func (w *World) MethodByRef(ref string) (*Method, error) {
	dot := strings.Index(ref, ".")
	colon := strings.Index(ref, ":")
	if dot < 0 || colon < dot {
		return nil, fmt.Errorf("bad method reference %q", ref)
	}
	c, ok := w.Classes[ref[:dot]]
	if !ok || !c.App {
		return nil, fmt.Errorf("no application class for %q", ref)
	}
	m := c.MethodBy[ref[dot+1:]]
	if m == nil {
		return nil, fmt.Errorf("no method %q", ref)
	}
	return m, nil
}
