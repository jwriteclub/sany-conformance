\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: SpecObj.java:1000 RemoveConditional_EQUAL_IF: this module has a PlusCal-only checksum divergence and Del01a a TLA-only one; each keeps its own warning, a mutant turns one into the 'both changed' warning (two mutants, one per clause)
\* Aim: SpecObj.java:1031 RemoveConditional_EQUAL_ELSE: the 'PlusCal algorithm ... has changed' warning for this module must appear; mutant drops it. Helper: Del01a
---- MODULE Del01 ----
EXTENDS Naturals, Del01a
(* --algorithm a {
  variable v = 0;
  { v := v + 1; }
} *)
\* BEGIN TRANSLATION (chksum(pcal) = "0" /\ chksum(tla) = "69e30512")
VARIABLES pc, v

vars == << pc, v >>

Init == (* Global variables *)
        /\ v = 0
        /\ pc = "Lbl_1"

Lbl_1 == /\ pc = "Lbl_1"
         /\ v' = v + 1
         /\ pc' = "Done"

(* Allow infinite stuttering to prevent deadlock on termination. *)
Terminating == pc = "Done" /\ UNCHANGED vars

Next == Lbl_1
           \/ Terminating

Spec == Init /\ [][Next]_vars

Termination == <>(pc = "Done")

\* END TRANSLATION
====
