---- MODULE W4803_Test ----
(*
--algorithm Test {
}
*)
\* BEGIN TRANSLATION (chksum(pcal) = "00000000" /\ chksum(tla) = "00000000")
VARIABLE pc

Spec == Init /\ [][Next]_vars

Termination == <>(pc = "Done")

\* END TRANSLATION 
====

