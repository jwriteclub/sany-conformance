\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Helper of Del01: TLA translation checksum is wrong, PlusCal checksum is right
---- MODULE Del01a ----
EXTENDS Naturals
(* --algorithm a {
  variable v = 0;
  { v := v + 1; }
} *)
\* BEGIN TRANSLATION (chksum(pcal) = "c932f2a7" /\ chksum(tla) = "0")
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
