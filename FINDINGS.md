<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Oddities in SANY, for the TLA+ team

Things we found while running this corpus against SANY at tlaplus `dc6a944`, each with a way to reproduce it. We are not claiming that any of them is a bug.

To reproduce, run the commands from a case's directory; `scripts/setup reference` builds the jar. Line numbers are at that commit. `TP` is `src/tla2sany/parser/TLAplusParser.java` and `TM` is `TLAplusParserTokenManager.java`, both under `tlatools/org.lamport.tlatools`.

## Java crashes, by signature

SANY crashes on 28 of the 9,453 cases. The table gives the smallest case for each signature; run `java -cp tla2tools.jar tla2sany.SANY <root>` on it.

| exception | thrown in | how it surfaces | cases | smallest, from the repository root | bytes |
|---|---|---|---:|---|---:|
| `util.WrongInvocationException`: Defined more recursive operators than were declared | `Generator.endOpDefNode` | `FrontEndException`, logged, with a stack trace | 6 | `deep-cuts/ours-model/model-one-shot/cases/E-f68eb4c7f7e2` | 81 |
| `NullPointerException`: `this.facts[i]` is null | `UseOrHideNode.factCheck` | same | 3 | `sets/ours-mechanical/fuzz-jazzer/cases/D-275fbcebbf60` | 38 |
| `ArrayIndexOutOfBoundsException` | `Generator.processLetIn` | same | 3 | `sets/ours-model/model-one-shot/cases/E-b1f53d508dac` | 1821 |
| `ArrayIndexOutOfBoundsException` | `Generator.processAction` | same | 2 | `sets/ours-mechanical/grammar-sentences/cases/S-582b422bd805` | 57 |
| `ArrayIndexOutOfBoundsException` | `Generator.genIdToSelector` | same | 1 | `sets/ours-mechanical/fuzz-jazzer/cases/D-2fecd0489510` | 49 |
| `ArrayIndexOutOfBoundsException`, negative index | not printed | logged as the only message, no trace, while parsing | 3 | `sets/ours-mechanical/fuzz-jazzer/cases/D-29d6d46361df` | 153 |
| `ArrayIndexOutOfBoundsException`, positive index | not printed | same | 3 | `sets/ours-mechanical/fuzz-jazzer/cases/D-0b19d850de35` | 61 |
| `IllegalArgumentException`: fromIndex(N) > toIndex(-1) | not printed | same | 3 | `deep-cuts/ours-mechanical/fuzz-jazzer/cases/D-484d3d727dce` | 359 |
| `NumberFormatException` on a long number | not printed | same | 3 | `sets/ours-mechanical/fuzz-jazzer/cases/D-5ba7194164bc` | 55 |
| `StackOverflowError` | no stable frame | escapes, exit 1. Overflows at `-Xss512k`, the default 1 MiB, and `-Xss16m` | 1 | `sets/ours-model/model-one-shot/cases/E-550917c970f7` | 63202 |

`XMLExporter` also fails on some specs that SANY accepts:

- `sets/ours-model/model-one-shot/cases/E-e0456e574085` (a `DEFINE` step named in `BY DEF`) throws `IllegalArgumentException: unsupported kind: 37 in xml export` and exits 1. So does `deep-cuts/ours-model/model-one-shot/cases/E-72e9c889699d`.
- `sets/tla-project/tlaplus-test-models/cases/A-fix-sany-stringformfeedescape` and 8 others have U+000C in a string. They exit 7, and the exporter says why: XML 1.0 cannot hold that character.
- 11 upstream Farsite modules and mutants of them, such as `sets/tla-project/tlaplus-other/cases/B-org-lamport-tla-toolbox-uitest-farsite-d-b17fc2`, run out of memory at a 15.7 GiB heap.

## SANY against tlx

tlx (tlaplus-x/tlx, `bec5985`) is a Rust implementation with its own parser. It is not a reference, so a disagreement says only that the two differ. They disagree on accept or reject on 302 cases: SANY accepts and tlx rejects on 155, and the other way round on 147.

`data/elements-extra.tsv` lists all of them, in the rows starting `tlx`, under 55 signatures. A signature is SANY's verdict, tlx's verdict, SANY's first code, and tlx's message with names and numbers blanked. To reproduce, use `tools/tlxcheck`, a wrapper over tlx's `Program::load`; build it with `cargo build --release --locked` after `scripts/regenerate`.

The largest signatures, with an example case where there is one:

- SANY rejects at parsing with code 4003, tlx accepts: 32 cases (`A-syn-jlist-006`).
- tlx `syntax error: missing _ (in in)`: 27 (`B-org-lamport-tla-toolbox-uitest-farsite-a-6ba330`).
- SANY code 4310 (a prime in an invariant), tlx accepts: 21, all mutants.
- tlx `missing _ (in label_as)`: 17.
- tlx `unexpected _ (in bound_infix_op)`: 15 (`A-sem-implicitproofstest`).
- tlx `unexpected _ (in operator_definition)`: 15 (`A-syn-let-in-010`).
- SANY code 4005, tlx accepts: 14.
- SANY code 4202, tlx accepts: 13.

tlx's `corpus/external.json` records `Disalg-ICS-NJU/PaxosStore-TLA` as having no licence. That record is out of date: at the commit tlx pins, `0188955e70`, the repository's root `LICENSE` is MIT. tlx's 4 wrappers for it, in `sets/other-projects/tlx`, extend modules that tlx does not vendor, so SANY reports `Cannot find source file`, and that is what is recorded. With upstream's 8 modules from `experiment/TPaxos` and tlapm's `TLAPS` and `NaturalsInduction` beside them, SANY accepts all 4.

## `ℕ` is not `Nat`

RFC 5's table in tlaplus/rfcs (`accepted_rfcs/rfc5_unicode/tla-unicode.csv:81`) lists `nat_number_set,Nat,ℕ,U+2115`. tree-sitter-tlaplus's `grammar.js:600` has `choice('Nat', 'ℕ')  // From Naturals standard module`. But SANY's `Naturals.tla` defines only `Nat`, and `tla+.jj:975` and `:1325` treat `ℕ` as an ordinary identifier.

So with `EXTENDS Naturals`, `A == ℕ` gives `Unknown operator: 'ℕ'`, code 4200 (`deep-cuts/ours-inspection/inspection/cases/U-381e0d0d96f2`). `ℕ == Nat` followed by `A == ℕ` is accepted, and tlx rejects that as a duplicate definition (`deep-cuts/ours-inspection/inspection/cases/U-ee44cabf8907`). Upstream's syntax corpus treats it as a plain name (`test/tla2sany/corpus/unicode/numbers-unicode.txt:6-7`, `ℕ ≜ {}`); only the test translator maps it to `NAT_NUMBER_SET` (`TlaPlusParserOutputTranslator.java:412-413`).

## The memoised `jj_ntk` never went stale

`TP` has 130 sites that read `switch ((jj_ntk==-1)?jj_ntk():jj_ntk)`. In principle the memo could disagree with the next token, because actions in `tla+.jj` change tokens after they are read: 7 set `kind = IDENTIFIER` and 9 call `belchDEF()`.

A jar with all 130 sites replaced by `jj_ntk()` gives the same stock output as the reference on all 9,453 cases, apart from the varying stack trace of `E-550917c970f7` (`results/diff-ntk.tsv.gz`). To reproduce, make that replacement in `TP`, compile it against the jar, add the three classes with `jar uf`, and run `scripts/verify --jar`.

An earlier run with counters patched in, over 8,721 cases, found 0 stale reads in 15.6 million compute reads and 15.4 million memo reads. Both arms ran at 48 sites, only the compute arm at 54, only the memo arm at 27, and neither at 1. In 509 further runs, models wrote specs meant to make the memo go stale, and it was dead at all 16 mutation points; that is why `model-ternary-hunt` is empty. This is evidence, not proof. A reading of the 16 points agrees, but it was not exhaustive.

## Dead lexer arms

An exact simulation shows that 542 missed outcomes of `jjMoveNfa_*` in `TM` cannot be taken, and neither can 93 in the literal DFA.

For the NFA, the methods were parsed into an interpreter and every ordered active-state list was explored. All 65,536 chars reduce to between 5 and 79 classes per method. Entry points were logged from every literal prefix (509), each followed by 209 chars, in all 6 lexical states: 638,286 inputs. On the 538 lines with data, what the simulator can reach is exactly what the corpus covers, with nothing extra on either side.

The dead arms by shape:

- 224 arms in a state that is never active for that char range (for example `TM:4093`, `if (curChar == 69)` in state 105);
- 136 false sides of `if (kind > N)`;
- 88 `switch` labels;
- 48 conditions only ever seen one way;
- 43 false sides of `&& kind > N`;
- 3 `while (i != startsAt)`.

For the DFA, the sweep was exhaustive: all 508 prefixes `jjstrLiteralImages` needs, each followed by every ASCII char, in all 6 states. `TM:5120`, `++jjround == 0x7fffffff`, needs about 2 GiB of input and was not tried.

To see a miss: `go run ./cmd/casemap who TLAplusParserTokenManager.java:4093` reports 0 cases. `data/staticreach/generated-missed.tsv` lists them all. The simulator was a one-off and is not in this repository, so "dead" rests on the cross-check above.

## Unreachable syntax-error arms

`TP` has 78 `jj_consume_token(-1)` lines in productions. `tools/grammar-trace/run` patches a copy of the parser so that every decision and arm reports itself, then parses every `.tla` file in this tree: 14,673 files, 43 arms run.

Of the 35 that do not run, 27 cannot. Each is the `default` of a choice behind a guard that already tested the same token, so a syntax error lands on the guard's arm instead. `data/grammar-guards.tsv.gz` lists what can reach each one and which of those tokens have no case label. They are in `PrefixOpToken`, `NEPrefixOpToken`, `InfixOpToken`, `PostfixOpToken`, `Prelude`, `SomeFixDecl`, `Assumption`, `NewSymb` (twice), `Theorem`, `Proof`, `ParenthesesExpression`, `OpenExpression`, `OpOrExpr`, `SomeQuant`, `SomeTQuant`, `BraceCases` (three times), `FairnessExpr`, `JuncItem`, `Expression`, `ExtendableExpr`, `PrimitiveExp` (twice), `BangExt` and `OpenStart`.

The other 8 are `else` sides of `if (jj_2_k ...)`. 6 follow the one-way decisions below, and 2 (`TP:3421`, `:7209`) were not classified.

Of the 74 `LOOKAHEAD` decisions, 68 see their scan go both ways. 6 go only one way (`TP:3649`, `4979`, `5709`, `6152`, `6392`, `6822`), each because an earlier test already answered it.

`OpenStart()` at `TP:7226` has no caller in `src` or `test`. Only its lookahead routine `jj_3R_94` runs, at `TP:6067` and `:6383`. `sets/ours-model/model-openstart` and the `Below*` specs in `sets/ours-inspection/inspection` reach all eight alternatives of that scan, with its predicate both ways, at both sites.

Separately, `beginsProof` checks whether the second character of a token's image is `+` or `*`, whatever the token's kind. So `THEOREM TRUE "+x"` enters `Proof` on a string: `LaProofPlusString`, in `sets/ours-inspection/inspection/cases/U-4dabf3cce6b7`. The spec is rejected either way.

## The plain command

- **Exit 0 on a semantic rejection.** Without `-error-codes`, `tla2sany.SANY` exits 0 when it rejects a spec at the semantic stage: 3,250 cases here. With it, those exit 4. A parse failure exits 255 either way: 2,096 cases. Example: `sets/tla-project/tlaplus-semantic-error-corpus/cases/A-err-e4200`.
- **Codes are not printed.** Neither mode prints an error code. Codes are available only from `Errors.getMessages()`.
- **Each JVM leaves an empty directory.** `SimpleFilenameToStream` creates `tlc-<random>` in `java.io.tmpdir` (`FilenameToStream.java:185`) and registers its files for deletion, but not the directory itself (`SimpleFilenameToStream.java:228`). Run any case and count them.

## What depends on where it runs

No verdict moved with platform, JVM flags, stack depth or charset in these runs. `results/depends-on.tsv` lists every case whose output moved at all.

- **File system.** `sets/ours-mechanical/fuzz-jazzer/cases/D-5ac9249fe910` imports `FiniteSetS`, and a `FiniteSets.tla` sits beside it. On ext4 SANY says `Cannot find source file`. On APFS it says `File name 'FiniteSetS' does not match the name 'FiniteSets'`. `sets/ours-mechanical/fuzz-jazzer/cases/D-f82fa4953e79` imports `InTegers`. With a jar, both file systems say the file cannot be found, since lookup inside a jar is case-sensitive. On APFS with a class directory on the class path, SANY also finds `Integers.tla`.
- **Charset.** SANY rewrites a nested module to a temporary file in the default charset. With `LC_ALL=C`, `sets/ours-mechanical/fuzz-jazzer/cases/D-027940a247ed` gets `Lexical error at line 9, column 5. Encountered: " " (32), after : "?"`, where UTF-8 gives `Could not find declaration or definition of symbol '≃'`. 31 cases change their first position this way, and all are still rejected. 113 more only print `?` for a non-ASCII character. With `-Dfile.encoding=ISO-8859-1`, 143 messages change and no verdict does.
- **Version.** v1.7.4 prints an older banner, and predates modules nested in a file and Unicode syntax: `results/diff-rel174.tsv.gz`.
