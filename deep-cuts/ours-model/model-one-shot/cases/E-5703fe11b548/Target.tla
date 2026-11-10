---- MODULE Target ----
(*
--algorithm Test {
  {
    lbl: skip;
  }
}
*)
\* BEGIN TRANSLATION (chksum(pcal) = "00000000" /\ chksum(tla) = "00000000")
VARIABLE pc

vars == << pc >>

Init == pc = "lbl"

lbl == /\ pc = "lbl"
       /\ pc' = "Done"

Terminating == pc = "Done" /\ UNCHANGED vars

Next == lbl \/ Terminating

Spec == Init /\ [][Next]_vars

\* END TRANSLATION 
====
