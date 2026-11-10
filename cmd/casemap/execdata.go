// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// JaCoCo's exec file format (org.jacoco.core.data.ExecutionDataWriter).
const (
	blockHeader  = 0x01
	blockSession = 0x10
	blockExec    = 0x11
	execMagic    = 0xC0C0
	execVersion  = 0x1007
)

// ClassData is one class's probe array.
type ClassData struct {
	ID     uint64
	Name   string
	Probes []bool
}

// ExecData maps a class id to its probes.
type ExecData map[uint64]*ClassData

// Merge ORs one class's probes in, keeping the name.
func (e ExecData) merge(c *ClassData) error {
	have, ok := e[c.ID]
	if !ok {
		e[c.ID] = &ClassData{ID: c.ID, Name: c.Name, Probes: append([]bool(nil), c.Probes...)}
		return nil
	}
	if len(have.Probes) != len(c.Probes) {
		return fmt.Errorf("class %s: %d probes against %d", c.Name, len(c.Probes), len(have.Probes))
	}
	for i, p := range c.Probes {
		if p {
			have.Probes[i] = true
		}
	}
	return nil
}

// ReadExec reads an exec file; several JVMs may have appended to it.
func ReadExec(path string) (ExecData, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readExec(bufio.NewReader(f))
}

func readExec(r *bufio.Reader) (ExecData, error) {
	out := ExecData{}
	for {
		t, err := r.ReadByte()
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
		switch t {
		case blockHeader:
			var h [4]byte
			if _, err := io.ReadFull(r, h[:]); err != nil {
				return nil, err
			}
			if binary.BigEndian.Uint16(h[:2]) != execMagic || binary.BigEndian.Uint16(h[2:]) != execVersion {
				return nil, fmt.Errorf("exec header %x: not JaCoCo 0.8.x data", h)
			}
		case blockSession:
			if _, err := readUTF(r); err != nil {
				return nil, err
			}
			var skip [16]byte
			if _, err := io.ReadFull(r, skip[:]); err != nil {
				return nil, err
			}
		case blockExec:
			var id [8]byte
			if _, err := io.ReadFull(r, id[:]); err != nil {
				return nil, err
			}
			name, err := readUTF(r)
			if err != nil {
				return nil, err
			}
			probes, err := readBools(r)
			if err != nil {
				return nil, err
			}
			if err := out.merge(&ClassData{ID: binary.BigEndian.Uint64(id[:]), Name: name, Probes: probes}); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unknown exec block type 0x%02x", t)
		}
	}
}

func readUTF(r *bufio.Reader) (string, error) {
	var l [2]byte
	if _, err := io.ReadFull(r, l[:]); err != nil {
		return "", err
	}
	b := make([]byte, binary.BigEndian.Uint16(l[:]))
	if _, err := io.ReadFull(r, b); err != nil {
		return "", err
	}
	return string(b), nil
}

func readBools(r *bufio.Reader) ([]bool, error) {
	n, err := readVarInt(r)
	if err != nil {
		return nil, err
	}
	out := make([]bool, n)
	var b byte
	for i := 0; i < n; i++ {
		if i%8 == 0 {
			if b, err = r.ReadByte(); err != nil {
				return nil, err
			}
		}
		out[i] = b&1 != 0
		b >>= 1
	}
	return out, nil
}

func readVarInt(r *bufio.Reader) (int, error) {
	v, shift := 0, 0
	for {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		v |= int(b&0x7f) << shift
		if b&0x80 == 0 {
			return v, nil
		}
		shift += 7
	}
}
