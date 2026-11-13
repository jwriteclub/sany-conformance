<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# SANY and prudence, case by case

Java side: the recorded stock-SANY output in each manifest.jsonl.
Prudence side: `prudence.Check`, the library call, stage **default (full semantic and level analysis)**, with support modules read from the case directory.

**Rule 3d is off** (`-no-3d`): every nested-module error stays where rule 3c puts it.

## What is compared, and what is not

Per case, each side is reduced to a **class** and a list of **errors**.

Classes (also `missing-module` and `module-cycle`, below): `ok` (accepted), `parse-error` (lexical or syntax rejection), `semantic-error` (rejected after parsing), `crash`, `missing-module` and `module-cycle` (a module SANY could not find, or a circular dependency; rule 1), and, for prudence only, `incomplete` (it stopped before deciding; its semantic and level analysis is not finished).
Each error is a line, a column and a message text.

A case is a **match** only if the class, the number of errors, every line and column, and every message text agree.
A case is **prudence incomplete** if prudence did not decide and SANY did not reject it at parsing (or only failed to find a module). This is silence, not disagreement.
Everything else is a **mismatch**, with a kind: *class*, *position* (same class, different places) or *text* (same places, different words).

**Normalisation. This is the whole list.**

1. SANY side, parse errors (recognised by the text ***Parse Error***, Lexical error, Lexical {error or Error in fairness expression; SANY pushes the last one on its parser error list and then prints its usual 4003): one error per record SANY's parser logged. TLAplusParser.java:113-116 (pinned dc6a944) prints each entry of its PErrors list with its own println, and an entry starts with one of these whole lines: `***Parse Error***` (a ParseException, :85), `Lexical {error: EOF reached, ...` or `Lexical error at line L, column C.  Encountered: ...` (:93-97), `@ used in !.@` (:5323), `Error in fairness expression at L: C` (:5536, :5555) or `Couldn't properly parse expression-- incomplete expression at ...` (OperatorStack.java:447); the last three are pushed without stopping the parse, so a ParseException record can follow them. Each record runs to the next one or to `Fatal errors`. Its position is the first of `Encountered "..." at line L, column C`, `Lexical error at line L, column C`, `Error in fairness expression at L: C`, or any `line L, column C` in its text, and none (rule 5) if its text names none, as `@ used in !.@` does. Its message is its text, residual stack trace included, minus the `***Parse Error***` line. Output from after the `Parsing file` line up to `Fatal errors` that is not exactly a sequence of such records stays one error, positioned by the first such pattern in the output. On the prudence side the same `***Parse Error***` prefix is dropped from its message.
   Class `missing-module`: SANY printed `Cannot find source file for module` and no parse error; one error with no position (0:0) and just SANY's one sentence as text (the "*** Errors: N" tail that repeats it is cut). Class `module-cycle` is the same for SANY's `Circular dependency among .tla files`, with the sentence and the cycle path as text, cut before the "*** Errors" tail. Prudence's diagnostics with code `semantic/4220` map to missing-module and `semantic/4222` to module-cycle (older patterns `module/...not-found`, `module/...cycle` are kept); with neither, its answer there is "incomplete".
2. SANY side, semantic errors (rewritten after the 9,453-case run): SANY (tla2sany/drivers/SANY.java, `frontEndSemanticAnalysis`) shares one Errors object across all modules and, after each `Semantic processing of module X` line, prints the whole accumulated list again under `Semantic errors:` / `*** Errors: N`; `Errors.addMessage` already ignores a message equal to one it has (same code, location and text). So the entries of every `*** Errors` section are read, SANY's own between-list lines (`Semantic processing of module`, `Semantic errors`, `Fatal errors`, `Parsing file`) end the entry in progress and are never part of a message, and an entry is one distinct (location line, text): repeats are dropped. A section order of Errors then Warnings (the abort path "Semantic errors detected before the unexpected exception") is read the same way; the outcome is semantic-error whenever any `*** Errors` section exists, whatever follows it. Each error carries the module its location line names (`of module M`), its starting `line L, col C` and its text. Entries under `*** Warnings` are not compared. An entry at `Unknown location` (e.g. E4221) is kept with no position (0:0, rule 5). An entry whose location is SANY's never-set sentinel (`line 2147483647, col 2147483647 to line -2147483648 ...`) is also kept with no position and no module, instead of being glued onto the previous message.
3. Prudence side: every diagnostic of severity error. Its start position is put in SANY's terms (rule 3a, 3b below); informational diagnostics (including its "incomplete" note) are not errors.
   3a. For a lexical or syntax error (not a semantic one, whose messages name other places), if the message text holds a SANY-style position (Encountered ... at line L, column C, else the first line L, col C), that position is used.
   3b. Otherwise the line is prudence's line plus one, and the column is rebuilt from the source line of the file the diagnostic is in (its Path, read from the case directory; the root file if Path is the root or unreadable) and prudence's ByteOffset with SANY's tab rule: JavaCC's SimpleCharStream has tabSize = 8, one column per UTF-16 unit, and a tab advances to the next multiple of eight (checked in tla2sany/parser/SimpleCharStream.java, column--; column += tabSize - (column % tabSize)). Prudence's own Column is per byte by design and is not used.
   3c. Module and line for a semantic error: the prudence diagnostic's file (its Path) is scanned for `---- MODULE X ----` headers and for the closing `====` token (four or more equals signs anywhere on a line, as SANY's lexer takes it), and the error is attributed to the innermost module whose lines contain it. SANY numbers lines from the file start for the first top-level module of a file, but from the module's own header line for a later top-level module in the same file (seen in the corpus: prudence 451:38 is SANY 66:38 of module Common, which starts at line 386); a module nested in another counts as its top-level ancestor, in line base and in name (an error in Inner nested in Target is printed "of module Target", case E-42f14edda2fd). Prudence's line is put into that numbering. Module names are compared as well as line and column, when both sides have one. A prudence error in a file prudence's Path points to but the harness cannot read, or in no module, has no module and its line is unchanged.
   3c'. Module headers and ends, as SANY's lexer reads them: a header is BEGIN_MODULE (tla+.jj:948, `"----" ("-")* (" ")* "MODULE"`: spaces only, and the name may follow MODULE directly, as in `---- MODULEOFt ----`), spaces or tabs, one name token BeginModule accepts (an identifier, or a keyword isFieldNameToken turns into one; not a number, not a word cut by the `WF_`/`SF_` tokens), spaces or tabs, and SEPARATOR (four or more dashes), all on one line; a line ends at \n, \r\n or a lone \r; a ==== after `\*` on a line is part of a comment and ends nothing; a module header is found anywhere on a line while no module is open (SANY skips arbitrary text, comments included, up to the next header), but only outside `\*` comments inside a module.
   3d. Nested modules reached by name (EXTENDS or INSTANCE) are extracted as units of their own, so SANY prints an error inside nested module N twice: in the enclosing module (rule 3c numbering) and "of module N" with lines counted from N's header. Prudence reports both at the same root-file bytes. Where unambiguous, the second of exactly two identical prudence errors (same byte offset, same text) inside N is put in N's own unit: module N, line minus (N's header line - 1). Unambiguous means: N is nested directly in a top-level module, N's name is declared once in the file, N and its parent are closed by ====. Counts, columns, byte spans and the other errors are untouched. Otherwise (a module declared more than once, an unclosed module, nesting deeper than one level, three or more identical errors) nothing is moved, and the prudence side carries a note "nested-module attribution ambiguous (rule 3d)" with the reason, so the case stays a visible mismatch if the numbers differ.
   3e. When prudence reports the unit that emitted an error (its Diagnostic.Unit: the module, and SANY's 1-based line and UTF-16, tab-expanded column in that unit's own text, set for a monolith-extracted unit), those are the error's module, line and column, exactly as SANY prints "line L, col C ... of module N". Rules 3a-3c are not applied to it and rule 3d never moves it or pairs it. An error with no Unit is projected by 3a-3d as before. A Unit with no module or a line or column below 1 makes the case a prudence error. Such an error is shown with "(unit)" and its row carries `"unit": true`.
4. Both sides: runs of whitespace in a message become one space. Errors are sorted by line, column, text, module.
5. Where SANY's error has no line and column (missing module, cycle, 4003 "Could not parse module" / "An unexpected problem ... validate", 4221 at "Unknown location", entries at In module M with no coordinates such as the "In module --TLA+ BUILTINS--" entries, and the never-set sentinel), the position is not compared. Each such SANY error is matched to a prudence error with the same text, wherever prudence puts it; if none has that text it is paired with a prudence error no positioned SANY error claims and counts as a text mismatch; if there is none, it is a position (count) mismatch. The remaining errors are compared in order. All entries SANY prints are kept, including the two of 4003 plus 4220/4222 (cycle and missing-module cases). An "In module M" line on its own starts a new SANY entry (it is a location, as SANY's Location.toString prints it for coordinates 0,0,0,0), it is not a message continuation. Parse errors keep one error per parser record (rule 1): the 4003 "Could not parse module" that SANY prints after every parse error is not added.
6. A case where SANY itself crashed (recorded status crash, or a live JVM failure) is `oracle-unavailable`, its own outcome in the summary: never a match (before, a prudence crash against it was scored as a match), never a false complete. False complete means prudence completes with no error where SANY rejects.
7. Nothing else, apart from the one exception in rule 8. No message rewriting, no position tolerance, no dropping of cases.
8. One exception: the order of the names in SANY's LABEL_PARAMETER_UNNECESSARY (4332) extra-parameter list. SANY builds the text `Label <name> declares extra parameter(s)  <names>` in tla2sany/semantic/Generator.java `formalParamsEqual` (pinned dc6a944, lines 7092-7099) by iterating a HashSet of the label's residual formal parameters; the loop only appends names to the string, and the method's result is discarded by its one caller (line 6933). Two errors that already agree in position and module and differ in text are equal under this rule only if both texts are exactly that template (after rule 4) with the same label, each list names distinct parameters, and the two lists hold the same names, none missing, added or repeated. No other 4332 text (`Illegal parameter %s of label`), no repeated-name list and no other message, position, module, count or class is relaxed. A case whose only differences are such lists is counted as **match (order-normalised)**, never as a match; its row keeps both sides as recorded and adds the rule's name and the raw texts.

**Not compared:** the syntax tree, the semantic graph and UIDs, warning codes, error codes, end positions of ranges, and anything SANY prints on stderr. Those are covered by a separate, deeper XML-level comparison. SANY's level errors are semantic errors in its output, so they are compared like any other. A match here says the two agree on accept or reject and on where and why the first-line complaints are, not that the trees are the same.

Line one of each manifest.jsonl describes the environment, not a case, and is skipped. What SANY did is a fact about SANY, not a statement that it is right.

## Summary

| set | cases | match | match (order-normalised) | mismatch | prudence incomplete | errors | timeouts | oracle unavailable | false complete |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| `deep-cuts/inspection-pit-survivors` | 4 | 4 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `deep-cuts/inspection` | 146 | 144 | 0 | 0 | 0 | 0 | 0 | 2 | 0 |
| `deep-cuts/fuzz-jazzer` | 103 | 93 | 0 | 0 | 0 | 0 | 0 | 10 | 0 |
| `deep-cuts/grammar-sentences` | 373 | 373 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `deep-cuts/mutants` | 2937 | 2936 | 0 | 0 | 0 | 0 | 0 | 1 | 0 |
| `deep-cuts/model-agents` | 22 | 22 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `deep-cuts/model-one-shot` | 178 | 174 | 0 | 0 | 0 | 0 | 0 | 4 | 0 |
| `deep-cuts/model-pit-survivors` | 4 | 4 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `deep-cuts/model-scoring` | 3 | 3 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `tlapm` | 677 | 677 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `tlx` | 587 | 587 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `tree-sitter-tlaplus` | 357 | 357 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `inspection-pit-survivors` | 68 | 68 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `inspection` | 42 | 42 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `fuzz-jazzer` | 1132 | 1125 | 1 | 0 | 0 | 0 | 0 | 6 | 0 |
| `grammar-sentences` | 15 | 14 | 0 | 0 | 0 | 0 | 0 | 1 | 0 |
| `mutants` | 63 | 63 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `model-agents` | 19 | 19 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `model-one-shot` | 163 | 161 | 0 | 0 | 0 | 0 | 0 | 2 | 0 |
| `model-openstart` | 23 | 23 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `model-pit-survivors` | 23 | 23 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `model-scoring` | 5 | 5 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `tlaplus-communitymodules` | 80 | 80 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `tlaplus-examples` | 433 | 433 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `tlaplus-inline-java` | 52 | 52 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `tlaplus-other` | 296 | 296 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `tlaplus-rfcs` | 475 | 474 | 0 | 0 | 0 | 0 | 0 | 1 | 0 |
| `tlaplus-semantic-corpus` | 27 | 27 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `tlaplus-semantic-error-corpus` | 121 | 120 | 0 | 0 | 0 | 0 | 0 | 1 | 0 |
| `tlaplus-syntax-corpus` | 355 | 355 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| `tlaplus-test-models` | 670 | 670 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| **all** | 9453 | 9424 | 1 | 0 | 0 | 0 | 0 | 28 | 0 |

**Match (order-normalised)** is counted apart from match: the case agrees in class, count, every position, module and text, except that one or more 4332 extra-parameter lists name the same parameters in a different order (rule 8). It is not a match. Each such case is listed below with both raw texts, and its results.jsonl row carries `policy` with the rule's name, the errors it was needed for and the SHA-256 of each side's text.

**Acquisition-visit overrides**: 1 case(s) scored under a recorded per-case `Options.MaxAcquisitionVisits` (not an outcome; each is counted in its normal column above).

- `tlaplus-test-models/B-test-model-bug156-te-9a905e`: `Options.MaxAcquisitionVisits` = 3000000 (applied; scored match). TE.tla's module acquisition needs more than the default 1,000,000 visits per pass (root reviewer)

**Timeouts** are their own outcome: the per-case deadline passed before prudence answered. They are never counted as a match or as prudence incomplete.

Mismatches by kind, all sets: class 0, position 0, text only 0.

**Oracle unavailable**: SANY itself crashed on the case (recorded or live), so there is no verdict to compare with. It gets no match credit whatever prudence said, and is not counted as false complete.

**False complete** is counted apart from the columns before it: prudence's public Completion is `complete`, it reports no error, and SANY reports one (not a SANY crash). It says prudence claimed to have finished and missed an error. Prudence's Completion and Level are in `results.jsonl` for every case.

## Mismatches

None.
## Matches that needed rule 8 (order-normalised)

#### `fuzz-jazzer/D-54cb9f0ad55d`

Policy: label-extra-parameter-multiset (coord352).

- error 5 at SubrefLL_test 7:23
  - SANY (`e23735fb0ec2`): `Label lab declares extra parameter(s) x i z`
  - prudence (`8655cc8a5b80`): `Label lab declares extra parameter(s) i x z`

## Errors

None.

## Timeouts

None.
