<!-- Copyright 2026 Christopher O'Connell -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

# Results for manyan

[manyan](https://github.com/soasme/manyan) is an MIT-licensed TLA+ implementation in C by soasme, with its own front end and model checker. This page compares it with stock SANY on this corpus, and with TLC on manyan's own configured test inputs. The differences below are compatibility observations against SANY's recorded behaviour, not judgements on manyan.

What was compared:

- **manyan** 0.10.0 at [`1bbb2c5`](https://github.com/soasme/manyan/tree/1bbb2c5ce19c855684f3543d50075ff5b2650cab), built from source.
- **SANY**: the recorded output of the unmodified tlaplus `dc6a944` build in each manifest (see the [README](../README.md)), including its recorded crashes.
- **TLC**: the 1.8.0 pre-release build `d8fb231` that manyan's own test runner pins, SHA-256 `e1f11b75c4b5bc4478ba0ac5cb6ae28e40a01b0016bc33e0aa97da22dfdbcd7f`, with CommunityModules `202611060952`, SHA-256 `a3b27c40f1fb5085f30c6ead37e918457f8cfc62c1e50a1dc668ef592ad2b5fb`. This is a different tlaplus build from the SANY reference.
- **The corpus**: all 9,453 cases in 31 manifests, deep cuts included. The manifests are byte-identical to those in this repository's initial public release (`04fc94bd`); `identities.json` lists their SHA-256s.

Measured on 2026-11-13 on macOS 15.5 arm64, one process per case.

## Front end

**manyan agrees with stock SANY on 8,430 of 9,424 comparable accept/reject verdicts (89.45%), and differs on 994 (10.55%).** The same 994 differences appear against the recorded [prudence results](results-prudence.md).

| Reference / manyan | accept / accept | reject / reject | reference rejects, manyan accepts | reference accepts, manyan rejects | agreement |
|---|---:|---:|---:|---:|---:|
| Stock SANY | 3,881 | 4,549 | 769 | 225 | 8,430 / 9,424 (89.45%) |
| prudence, complete calls | 3,881 | 4,558 | 769 | 225 | 8,439 / 9,433 (89.46%) |

Only accept or reject is compared here, not error positions or messages. The SANY side is the manifests' recorded output; the prudence side is [`results/prudence-b801ed2/results.jsonl.gz`](../results/prudence-b801ed2/), whose SHA-256 is recorded in `identities.json`. Neither reference was rerun. Counts are cases, including related and duplicate inputs.

Every input file's hash was checked against its manifest before the run. Each case used its original root filename, support files and library links, and no other modules were made visible.

How the 9,453 cases reduce to 9,424:

- manyan gave a verdict on 9,451 cases.
- `B-test-model-bug156-te-9a905e` (TE.tla) was not run, because of its resource needs on the measuring machine.
- `E-550917c970f7` ended with SIGSEGV, no output and no verdict; the cause was not diagnosed. Its 63,202-byte source holds 63,171 consecutive opening parentheses. SANY also crashes on it, and prudence reports incomplete.
- The 28 cases where SANY crashed are not scored for anyone. manyan accepted 11, rejected 16 and signalled on one.

prudence's 9,433 pairs exclude its 19 incomplete calls and TE. Its nine extra pairs are SANY-crash cases, not SANY matches.

### Differences by class

| Stock SANY outcome | manyan outcome | cases |
|---|---|---:|
| semantic error | accept | 621 |
| parse error | accept | 139 |
| missing module | accept | 8 |
| module cycle | accept | 1 |
| accept | reject | 225 |

Of the 225 cases manyan rejects, it reports 197 as parse errors, 20 as semantic errors and eight as missing modules. Of the parse diagnostics, 136 begin `Unexpected character` and name the single byte `e2` (the first byte of a UTF-8 encoded symbol), 45 say `Encountered`, 15 report a precedence conflict and one an unterminated string. These are groups of diagnostics, not a claim that each group has one cause.

Examples, each checked against the source:

| Case | Observation |
|---|---|
| `A-err-e4200-use` | `USE DEF DoesNotExist`: SANY rejects the unknown operator. manyan accepts; its proof handling examines HIDE but skips this USE body. |
| `U-04c61bc862db` | `THEOREM ASSUME IF , PROVE TRUE`: SANY rejects the malformed statement. manyan's theorem-ASSUME branch skips it without parsing the expression. |
| `B-tests-specs-ap1-a1fc54` | `EXTENDS Apalache`: SANY cannot find this module in the case. manyan supplies a built-in Apalache module and accepts. This is a difference in module resolution environment. |
| `S-11fabfc9980e` | A proof on the theorem's line holds `INSTANCE x` in module x. SANY reports a dependency cycle; manyan skips the proof and accepts. |
| `U-df29a2200669` | `B == 1 ∈ Nat`: SANY accepts. manyan rejects the first UTF-8 byte of the membership symbol. |
| `B-fast-language-wftrue-test-c67c8e` | SANY accepts `WF_TRUE(TRUE)`. manyan rejects the bare TRUE subscript, where its parser expects another expression form. |
| `B-org-lamport-tlatools-test-model-github14-cf613b` | SANY accepts `SUBSET {0, 1} \cup SUBSET {2}`. manyan reports a SUBSET/union precedence conflict. |

The relevant code at the pinned commit is in
[`manyan.c`](https://github.com/soasme/manyan/blob/1bbb2c5ce19c855684f3543d50075ff5b2650cab/manyan.c):
proof skipping at lines 3337–3379 and 3437–3455; Unicode lexing at 2099–2105
and 2275–2281; fairness subscripts at 2638–2657 and 2769–2773; precedence at
2427, 2441 and 3109–3142; built-in module lookup at 3888–3905.
[`manyan-std.c`](https://github.com/soasme/manyan/blob/1bbb2c5ce19c855684f3543d50075ff5b2650cab/manyan-std.c#L1093)
registers the Apalache module text.

### How the front-end verdict was read

manyan has no front-end-only command. For this comparison a separate source copy replaces the one line `if (o->typecheck) return typecheck(root);` with an acceptance marker and a successful return. That point, `manyan.c:20605`, comes after the unchanged root parse, filename check, dependency loading, name resolution and level checks, and before type inference, TLC-specific numeric restrictions, native binding, configuration evaluation, assumptions and state exploration.

The rest of the command-line start-up is unchanged. An empty configuration is supplied, because manyan reads a configuration before it parses the module. Exit 0 with the marker means accept; exit 150 means reject. Signals, timeouts, resource stops and other exits give no verdict. This measures manyan's own front end, including its own handling of proofs. The TLC comparison below uses the unmodified build.

Both builds used Apple clang 17.0.0 with `-O2 -std=c11 -Wall -Wextra -Wpedantic`, `manyan-std.c` and `-lm`, without profile-guided optimisation. No performance claim is made.

## Checking against TLC

| Observation | matching pairs / pairs with the observation |
|---|---:|
| exit status and outcome class | 778 / 778 |
| generated, distinct and queued state counts | 468 / 468 |
| printed counterexample traces | 217 / 217 |
| report text, after normalisation | 778 / 778 |

The inputs are every `tests/spec` root input at the pinned manyan commit (779), its two `tests/tla/MC*.tla` models and `examples/DieHard`: 782 in all. Sibling `.files` modules, embedded configurations, `.in` inputs and named options were staged with the helpers in manyan's own test runner. Both engines ran with one worker, including the one input that asks for four. Symbolic and type-inference inputs target Apalache and are not included.

Four of the 782 have no comparable pair, and none is credited:

- manyan's runner skips TLC on three by design: `config/stdin` uses a configuration path TLC does not support, and `errors/stack_eval` and `errors/stack_parse` use manyan's `-stack` option.
- On `tests/tla/MCWorkers`, manyan went over the 320 MiB per-process threshold (347,505,264 bytes observed) and TLC reached the five-second limit.

The 778 finished pairs include 282 successes, 123 assumption violations, 18 deadlocks, 85 safety violations, 46 liveness violations and two assertion violations. The other 222 report configuration, front-end, evaluation or other errors, and are not counted as successful searches. For example, on `DieHard` both report a safety violation with 73 generated, 14 distinct and one queued state and the same seven-state trace; on `MCBFS` both succeed with 92,321 generated, 57,825 distinct and zero queued states.

Settings for this run: Java 17.0.21+9 (Temurin), UTF-8 output, 128 MiB heap, 96 MiB metaspace, 4 MiB Java stack, one active processor and one parallel GC thread (the flags are in `identities.json`). Each engine ran in its own writable copy of the same staged inputs. Each engine had a five-second wall limit and a 320 MiB per-process memory threshold, and each batch a 600-second, 512 MiB limit on its process group; hitting a limit gives no verdict. The memory limits are best-effort macOS footprint monitors, not hard allocation limits.

Report text is compared after the noise filtering in manyan's runner, which among other things collapses parse and semantic failures, and with TLC's generic warning hint dropped from both sides. It is not byte-for-byte output parity. Counts and traces are only credited when both sides print them. Traces are compared as ordered text, including state values and action labels; that they are shortest is not checked.

### An earlier run with other settings

A first full run, kept in `initial/`, used C-locale output, SerialGC and a 1 MiB Java stack. There 777 of 778 finished pairs had equal exits:

- `modules/recursive_deep` (a finite 400-step recursion): TLC hit StackOverflowError; manyan succeeded.
- `stdlib/json_to_json`: TLC printed `?` where manyan printed `é`.
- `expressions/except_missing`: repeated EXCEPT warnings differed.

The final settings (UTF-8 output and a larger stack) remove these differences; all three match there. The earlier results are kept rather than discarded. No general claim about warning equivalence is made.

## Limits

- **The corpus has no `.cfg` files**, in any manifest or case directory, so no corpus case is model checked. The TLC comparison uses manyan's own configured test inputs, chosen by manyan's authors; agreement on them does not establish general checking compatibility.
- **Front-end comparison is accept/reject only.** Error positions, messages and warnings are not compared.
- **Skipped and stopped inputs are not credited:** TE.tla in the front end, the three runner skips and the MCWorkers resource stop in checking, the SIGSEGV case, and the 28 SANY crashes.
- **The front-end verdict comes from a one-line modification** of manyan, described above, since it has no front-end-only mode.
- **The cases are not a random sample of TLA+ files**, and many are related: mutants are small edits of upstream specs.
- **One machine, one run of each setting.**

## Overlap between manyan's tests and this corpus

The pinned manyan checkout has 1,308 tracked `.tla` files with 1,250 distinct contents. Against every TLA+ root, support file and library in the corpus, 46 match byte for byte (34 distinct hashes), and 154 match after a module projection (132 distinct). The projection keeps the text from the first module header to the final standalone terminator, normalises line endings and strips trailing spaces and tabs; module names, comments, blank lines and indentation are kept. This is text overlap, not behavioural equivalence, and it does not find adapted copies.

| Source in manyan | root inputs | overlap after projection |
|---|---:|---:|
| `tests/spec`, excluding `tlc/` | 664 | 0 |
| `tests/spec/tlc` | 115 | 111 |
| `tests/symbolic/apalache` | 226 | 0 |
| `tests/typecheck/apalache` | 71 | 0 |

[`tests/spec/tlc/README.md`](https://github.com/soasme/manyan/blob/1bbb2c5ce19c855684f3543d50075ff5b2650cab/tests/spec/tlc/README.md) says those models come from the tlaplus v1.7.4 test-model suite, so most of them are modules this corpus already has, with their configurations. Of the four not matched, `test220` differs internally from the corpus module, and `tlcget_action`, `tlcget_generated_post` and `tlcget_generated_violation` are not in the corpus. No manyan input was added to this corpus.

## Files

This run's results are in [`results/manyan-1bbb2c5/`](../results/manyan-1bbb2c5/). Read the `.gz` files with `gzip -dc`.

| File | Contents |
|---|---|
| `frontend.jsonl.gz` | one row per corpus case: manyan's exit and verdict, the SANY and prudence reference classes, input hashes |
| `checking.jsonl.gz` | one row per checking input: both engines' exit, class, state counts, trace and output hashes |
| `summary.json` | the counts on this page |
| `initial/` | the same three files for the earlier checking run |
| `identities.json` | hashes of the manyan source, both builds, the TLC and CommunityModules jars, manyan's runner, the prudence results and every manifest; Java version and flags |
| `overlap.tsv`, `overlap-counts.json` | the overlap matches and counts above |
| `resource-receipts.json` | peak memory and exit status of each batch |
| `SHA256SUMS` | hashes of all of the above; check with `shasum -a 256 -c SHA256SUMS` |

`overlap-counts.json` records the corpus revision as `3b07ab72`, the pre-release revision of this repository; its manifests are the ones published here.
