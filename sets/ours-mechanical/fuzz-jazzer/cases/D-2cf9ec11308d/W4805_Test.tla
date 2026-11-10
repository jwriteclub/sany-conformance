If the PlusCal algorithm's translation h the hash code) then raise a warning.
---- MODULE W4805_Test ----
(*
--algorithm Test {
  {
    lbl: skip;
  }
}
*)
\* BEGIN TRANSLATION (chksum(pcal) = "c0cb232" /\ chksum(tla) = "00000000")
Spec == Init /\ [][Next]_vars

Termination == <>(pc = "Done")

\* END TRANSLATION 
====

