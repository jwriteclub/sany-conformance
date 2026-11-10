// Copyright 2026 Christopher O'Connell
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const maxListedPerSet = 40

type tally struct{ cases, match, orderNorm, mismatch, incomplete, errors, timeouts, oracle, falseComplete, kClass, kPos, kText int }

// caseRecord is one case as written to results.jsonl and progress.jsonl.
// "policy" is present only on a case whose result is ResultOrderNormalised, and
// "acquisition_visits" only on a case scored under an acquisition override, so every
// other row is byte for byte what it was before.
func caseRecord(c *Case) map[string]any {
	r := map[string]any{"set": c.Set, "id": c.ID, "result": c.Result, "kind": c.Kind, "false_complete": c.FalseComplete,
		"sany": c.Sany, "prudence": c.Prudence}
	if c.Policy != nil {
		r["policy"] = c.Policy
	}
	if c.Acquisition != nil {
		r["acquisition_visits"] = c.Acquisition
	}
	return r
}

func writeOutputs(out string, cases []*Case, sets []string, live bool, jar, stage string) error {
	// Machine-readable: one JSON object per case.
	jf, err := os.Create(filepath.Join(out, "results.jsonl"))
	if err != nil {
		return err
	}
	jw := bufio.NewWriter(jf)
	enc := json.NewEncoder(jw)
	for _, c := range cases {
		if err := enc.Encode(caseRecord(c)); err != nil {
			return err
		}
	}
	if err := jw.Flush(); err != nil {
		return err
	}
	if err := jf.Close(); err != nil {
		return err
	}

	bySet := map[string]*tally{}
	var total tally
	for _, c := range cases {
		t := bySet[c.Set]
		if t == nil {
			t = &tally{}
			bySet[c.Set] = t
		}
		for _, x := range []*tally{t, &total} {
			x.cases++
			if c.FalseComplete {
				x.falseComplete++
			}
			switch c.Result {
			case "match":
				x.match++
			case ResultOrderNormalised:
				x.orderNorm++
			case "mismatch":
				x.mismatch++
				switch c.Kind {
				case "class":
					x.kClass++
				case "position":
					x.kPos++
				case "text":
					x.kText++
				}
			case "incomplete":
				x.incomplete++
			case "timeout":
				x.timeouts++
			case "oracle-unavailable":
				x.oracle++
			default:
				x.errors++
			}
		}
	}

	var b strings.Builder
	source := "the recorded stock-SANY output in each manifest.jsonl"
	if live {
		source = "stock SANY, rerun now with " + jar
	}
	fmt.Fprintf(&b, "<!-- Copyright 2026 Christopher O'Connell -->\n<!-- SPDX-License-Identifier: Apache-2.0 -->\n\n")
	b.WriteString("# SANY and prudence, case by case\n\n")
	stageName := "default (full semantic and level analysis)"
	if stage != "" {
		stageName = "`" + stage + "` only"
	}
	fmt.Fprintf(&b, "Java side: %s.\nPrudence side: `prudence.Check`, the library call, stage **%s**, with support modules read from the case directory.\n\n", source, stageName)
	if rescoredFrom != "" {
		fmt.Fprintf(&b, "Rescored from `%s`: the recorded SANY and prudence answers there, with no new run of either.\n\n", rescoredFrom)
	}
	if no3d {
		b.WriteString("**Rule 3d is off** (`-no-3d`): every nested-module error stays where rule 3c puts it.\n\n")
	}
	b.WriteString(header)
	b.WriteString("\n## Summary\n\n| set | cases | match | match (order-normalised) | mismatch | prudence incomplete | errors | timeouts | oracle unavailable | false complete |\n|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	for _, s := range sets {
		t := bySet[s]
		fmt.Fprintf(&b, "| `%s` | %d | %d | %d | %d | %d | %d | %d | %d | %d |\n", s, t.cases, t.match, t.orderNorm, t.mismatch, t.incomplete, t.errors, t.timeouts, t.oracle, t.falseComplete)
	}
	fmt.Fprintf(&b, "| **all** | %d | %d | %d | %d | %d | %d | %d | %d | %d |\n\n", total.cases, total.match, total.orderNorm, total.mismatch, total.incomplete, total.errors, total.timeouts, total.oracle, total.falseComplete)
	b.WriteString("**Match (order-normalised)** is counted apart from match: the case agrees in class, count, every position, module and text, except that one or more 4332 extra-parameter lists name the same parameters in a different order (rule 8). It is not a match. Each such case is listed below with both raw texts, and its results.jsonl row carries `policy` with the rule's name, the errors it was needed for and the SHA-256 of each side's text.\n\n")
	b.WriteString(acquisitionFootnote(cases))
	b.WriteString("**Timeouts** are their own outcome: the per-case deadline passed before prudence answered. They are never counted as a match or as prudence incomplete.\n\n")
	fmt.Fprintf(&b, "Mismatches by kind, all sets: class %d, position %d, text only %d.\n\n", total.kClass, total.kPos, total.kText)

	b.WriteString("**Oracle unavailable**: SANY itself crashed on the case (recorded or live), so there is no verdict to compare with. It gets no match credit whatever prudence said, and is not counted as false complete.\n\n")
	b.WriteString("**False complete** is counted apart from the columns before it: prudence's public Completion is `complete`, it reports no error, and SANY reports one (not a SANY crash). It says prudence claimed to have finished and missed an error. Prudence's Completion and Level are in `results.jsonl` for every case.\n\n")
	b.WriteString("## Mismatches\n\n")
	if total.mismatch == 0 {
		b.WriteString("None.\n")
	}
	for _, s := range sets {
		listed, hidden := 0, 0
		var sb strings.Builder
		for _, c := range cases {
			if c.Set != s || c.Result != "mismatch" {
				continue
			}
			if listed == maxListedPerSet {
				hidden++
				continue
			}
			listed++
			fmt.Fprintf(&sb, "#### `%s` (%s)\n\nSANY:\n```\n%s\n```\nprudence:\n```\n%s\n```\n\n", c.ID, c.Kind, Show(c.Sany), Show(c.Prudence))
		}
		if listed == 0 {
			continue
		}
		fmt.Fprintf(&b, "### `%s`\n\n%s", s, sb.String())
		if hidden > 0 {
			fmt.Fprintf(&b, "%d more in `results.jsonl`.\n\n", hidden)
		}
	}
	b.WriteString("## Matches that needed rule 8 (order-normalised)\n\n")
	n := 0
	for _, c := range cases {
		if c.Result != ResultOrderNormalised || c.Policy == nil {
			continue
		}
		n++
		fmt.Fprintf(&b, "#### `%s/%s`\n\nPolicy: %s.\n\n", c.Set, c.ID, c.Policy.Name)
		for _, d := range c.Policy.Diffs {
			fmt.Fprintf(&b, "- error %d at %s%d:%d\n  - SANY (`%s`): `%s`\n  - prudence (`%s`): `%s`\n", d.Index, moduleLabel(d.Module), d.Line, d.Col, d.SanySHA256[:12], d.Sany, d.PrudenceSHA256[:12], d.Prudence)
		}
		b.WriteString("\n")
	}
	if n == 0 {
		b.WriteString("None.\n\n")
	}
	b.WriteString("## Errors\n\n")
	n = 0
	for _, c := range cases {
		if c.Result == "error" {
			n++
			fmt.Fprintf(&b, "- `%s/%s`: %s\n", c.Set, c.ID, Show(c.Prudence))
		}
	}
	if n == 0 {
		b.WriteString("None.\n")
	}
	b.WriteString("\n## Timeouts\n\n")
	n = 0
	for _, c := range cases {
		if c.Result == "timeout" {
			n++
			fmt.Fprintf(&b, "- `%s/%s`: %s\n", c.Set, c.ID, Show(c.Prudence))
		}
	}
	if n == 0 {
		b.WriteString("None.\n")
	}
	return os.WriteFile(filepath.Join(out, "report.md"), []byte(b.String()), 0o644)
}

func moduleLabel(m string) string {
	if m == "" {
		return ""
	}
	return m + " "
}

const header = `## What is compared, and what is not

Per case, each side is reduced to a **class** and a list of **errors**.

Classes (also ` + "`missing-module`" + ` and ` + "`module-cycle`" + `, below): ` + "`ok`" + ` (accepted), ` + "`parse-error`" + ` (lexical or syntax rejection), ` + "`semantic-error`" + ` (rejected after parsing), ` + "`crash`" + `, ` + "`missing-module`" + ` and ` + "`module-cycle`" + ` (a module SANY could not find, or a circular dependency; rule 1), and, for prudence only, ` + "`incomplete`" + ` (it stopped before deciding; its semantic and level analysis is not finished).
Each error is a line, a column and a message text.

A case is a **match** only if the class, the number of errors, every line and column, and every message text agree.
A case is **prudence incomplete** if prudence did not decide and SANY did not reject it at parsing (or only failed to find a module). This is silence, not disagreement.
Everything else is a **mismatch**, with a kind: *class*, *position* (same class, different places) or *text* (same places, different words).

**Normalisation. This is the whole list.**

1. SANY side, parse errors (recognised by the text ***Parse Error***, Lexical error, Lexical {error or Error in fairness expression; SANY pushes the last one on its parser error list and then prints its usual 4003): one error per record SANY's parser logged. TLAplusParser.java:113-116 (pinned dc6a944) prints each entry of its PErrors list with its own println, and an entry starts with one of these whole lines: ` + "`***Parse Error***`" + ` (a ParseException, :85), ` + "`Lexical {error: EOF reached, ...`" + ` or ` + "`Lexical error at line L, column C.  Encountered: ...`" + ` (:93-97), ` + "`@ used in !.@`" + ` (:5323), ` + "`Error in fairness expression at L: C`" + ` (:5536, :5555) or ` + "`Couldn't properly parse expression-- incomplete expression at ...`" + ` (OperatorStack.java:447); the last three are pushed without stopping the parse, so a ParseException record can follow them. Each record runs to the next one or to ` + "`Fatal errors`" + `. Its position is the first of ` + "`Encountered \"...\" at line L, column C`" + `, ` + "`Lexical error at line L, column C`" + `, ` + "`Error in fairness expression at L: C`" + `, or any ` + "`line L, column C`" + ` in its text, and none (rule 5) if its text names none, as ` + "`@ used in !.@`" + ` does. Its message is its text, residual stack trace included, minus the ` + "`***Parse Error***`" + ` line. Output from after the ` + "`Parsing file`" + ` line up to ` + "`Fatal errors`" + ` that is not exactly a sequence of such records stays one error, positioned by the first such pattern in the output. On the prudence side the same ` + "`***Parse Error***`" + ` prefix is dropped from its message.
   Class ` + "`missing-module`" + `: SANY printed ` + "`Cannot find source file for module`" + ` and no parse error; one error with no position (0:0) and just SANY's one sentence as text (the "*** Errors: N" tail that repeats it is cut). Class ` + "`module-cycle`" + ` is the same for SANY's ` + "`Circular dependency among .tla files`" + `, with the sentence and the cycle path as text, cut before the "*** Errors" tail. Prudence's diagnostics with code ` + "`semantic/4220`" + ` map to missing-module and ` + "`semantic/4222`" + ` to module-cycle (older patterns ` + "`module/...not-found`" + `, ` + "`module/...cycle`" + ` are kept); with neither, its answer there is "incomplete".
2. SANY side, semantic errors (rewritten after the 9,453-case run): SANY (tla2sany/drivers/SANY.java, ` + "`frontEndSemanticAnalysis`" + `) shares one Errors object across all modules and, after each ` + "`Semantic processing of module X`" + ` line, prints the whole accumulated list again under ` + "`Semantic errors:`" + ` / ` + "`*** Errors: N`" + `; ` + "`Errors.addMessage`" + ` already ignores a message equal to one it has (same code, location and text). So the entries of every ` + "`*** Errors`" + ` section are read, SANY's own between-list lines (` + "`Semantic processing of module`" + `, ` + "`Semantic errors`" + `, ` + "`Fatal errors`" + `, ` + "`Parsing file`" + `) end the entry in progress and are never part of a message, and an entry is one distinct (location line, text): repeats are dropped. A section order of Errors then Warnings (the abort path "Semantic errors detected before the unexpected exception") is read the same way; the outcome is semantic-error whenever any ` + "`*** Errors`" + ` section exists, whatever follows it. Each error carries the module its location line names (` + "`of module M`" + `), its starting ` + "`line L, col C`" + ` and its text. Entries under ` + "`*** Warnings`" + ` are not compared. An entry at ` + "`Unknown location`" + ` (e.g. E4221) is kept with no position (0:0, rule 5). An entry whose location is SANY's never-set sentinel (` + "`line 2147483647, col 2147483647 to line -2147483648 ...`" + `) is also kept with no position and no module, instead of being glued onto the previous message.
3. Prudence side: every diagnostic of severity error. Its start position is put in SANY's terms (rule 3a, 3b below); informational diagnostics (including its "incomplete" note) are not errors.
   3a. For a lexical or syntax error (not a semantic one, whose messages name other places), if the message text holds a SANY-style position (Encountered ... at line L, column C, else the first line L, col C), that position is used.
   3b. Otherwise the line is prudence's line plus one, and the column is rebuilt from the source line of the file the diagnostic is in (its Path, read from the case directory; the root file if Path is the root or unreadable) and prudence's ByteOffset with SANY's tab rule: JavaCC's SimpleCharStream has tabSize = 8, one column per UTF-16 unit, and a tab advances to the next multiple of eight (checked in tla2sany/parser/SimpleCharStream.java, column--; column += tabSize - (column % tabSize)). Prudence's own Column is per byte by design and is not used.
   3c. Module and line for a semantic error: the prudence diagnostic's file (its Path) is scanned for ` + "`---- MODULE X ----`" + ` headers and for the closing ` + "`====`" + ` token (four or more equals signs anywhere on a line, as SANY's lexer takes it), and the error is attributed to the innermost module whose lines contain it. SANY numbers lines from the file start for the first top-level module of a file, but from the module's own header line for a later top-level module in the same file (seen in the corpus: prudence 451:38 is SANY 66:38 of module Common, which starts at line 386); a module nested in another counts as its top-level ancestor, in line base and in name (an error in Inner nested in Target is printed "of module Target", case E-42f14edda2fd). Prudence's line is put into that numbering. Module names are compared as well as line and column, when both sides have one. A prudence error in a file prudence's Path points to but the harness cannot read, or in no module, has no module and its line is unchanged.
   3c'. Module headers and ends, as SANY's lexer reads them: a header is BEGIN_MODULE (tla+.jj:948, ` + "`\"----\" (\"-\")* (\" \")* \"MODULE\"`" + `: spaces only, and the name may follow MODULE directly, as in ` + "`---- MODULEOFt ----`" + `), spaces or tabs, one name token BeginModule accepts (an identifier, or a keyword isFieldNameToken turns into one; not a number, not a word cut by the ` + "`WF_`" + `/` + "`SF_`" + ` tokens), spaces or tabs, and SEPARATOR (four or more dashes), all on one line; a line ends at \n, \r\n or a lone \r; a ==== after ` + "`\\*`" + ` on a line is part of a comment and ends nothing; a module header is found anywhere on a line while no module is open (SANY skips arbitrary text, comments included, up to the next header), but only outside ` + "`\\*`" + ` comments inside a module.
   3d. Nested modules reached by name (EXTENDS or INSTANCE) are extracted as units of their own, so SANY prints an error inside nested module N twice: in the enclosing module (rule 3c numbering) and "of module N" with lines counted from N's header. Prudence reports both at the same root-file bytes. Where unambiguous, the second of exactly two identical prudence errors (same byte offset, same text) inside N is put in N's own unit: module N, line minus (N's header line - 1). Unambiguous means: N is nested directly in a top-level module, N's name is declared once in the file, N and its parent are closed by ====. Counts, columns, byte spans and the other errors are untouched. Otherwise (a module declared more than once, an unclosed module, nesting deeper than one level, three or more identical errors) nothing is moved, and the prudence side carries a note "nested-module attribution ambiguous (rule 3d)" with the reason, so the case stays a visible mismatch if the numbers differ.
   3e. When prudence reports the unit that emitted an error (its Diagnostic.Unit: the module, and SANY's 1-based line and UTF-16, tab-expanded column in that unit's own text, set for a monolith-extracted unit), those are the error's module, line and column, exactly as SANY prints "line L, col C ... of module N". Rules 3a-3c are not applied to it and rule 3d never moves it or pairs it. An error with no Unit is projected by 3a-3d as before. A Unit with no module or a line or column below 1 makes the case a prudence error. Such an error is shown with "(unit)" and its row carries ` + "`\"unit\": true`" + `.
4. Both sides: runs of whitespace in a message become one space. Errors are sorted by line, column, text, module.
5. Where SANY's error has no line and column (missing module, cycle, 4003 "Could not parse module" / "An unexpected problem ... validate", 4221 at "Unknown location", entries at In module M with no coordinates such as the "In module --TLA+ BUILTINS--" entries, and the never-set sentinel), the position is not compared. Each such SANY error is matched to a prudence error with the same text, wherever prudence puts it; if none has that text it is paired with a prudence error no positioned SANY error claims and counts as a text mismatch; if there is none, it is a position (count) mismatch. The remaining errors are compared in order. All entries SANY prints are kept, including the two of 4003 plus 4220/4222 (cycle and missing-module cases). An "In module M" line on its own starts a new SANY entry (it is a location, as SANY's Location.toString prints it for coordinates 0,0,0,0), it is not a message continuation. Parse errors keep one error per parser record (rule 1): the 4003 "Could not parse module" that SANY prints after every parse error is not added.
6. A case where SANY itself crashed (recorded status crash, or a live JVM failure) is ` + "`oracle-unavailable`" + `, its own outcome in the summary: never a match (before, a prudence crash against it was scored as a match), never a false complete. False complete means prudence completes with no error where SANY rejects.
7. Nothing else, apart from the one exception in rule 8. No message rewriting, no position tolerance, no dropping of cases.
8. One approved exception: the order of the names in SANY's LABEL_PARAMETER_UNNECESSARY (4332) extra-parameter list. SANY builds the text ` + "`Label <name> declares extra parameter(s)  <names>`" + ` in tla2sany/semantic/Generator.java ` + "`formalParamsEqual`" + ` (pinned dc6a944, lines 7092-7099) by iterating a HashSet of the label's residual formal parameters; the loop only appends names to the string, and the method's result is discarded by its one caller (line 6933). Two errors that already agree in position and module and differ in text are equal under this rule only if both texts are exactly that template (after rule 4) with the same label, each list names distinct parameters, and the two lists hold the same names, none missing, added or repeated. No other 4332 text (` + "`Illegal parameter %s of label`" + `), no repeated-name list and no other message, position, module, count or class is relaxed. A case whose only differences are such lists is counted as **match (order-normalised)**, never as a match; its row keeps both sides as recorded and adds the rule's name and the raw texts.

**Not compared:** the syntax tree, the semantic graph and UIDs, warning codes, error codes, end positions of ranges, and anything SANY prints on stderr. SANY's level errors are semantic errors in its output, so they are compared like any other. A match here says the two agree on accept or reject and on where and why the first-line complaints are, not that the trees are the same.

Line one of each manifest.jsonl describes the environment, not a case, and is skipped. What SANY did is a fact about SANY, not a statement that it is right.
`
