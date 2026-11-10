If the PlusCal algorithm's translation has been changed since the last time
the translation was run (as detected by the hash code) then raise a warning.
---- MODULE W4805_Test ----
(*
--algorithm Test {
  {
    lbl: skip;
  }
}
*)
\* BEGIN TRANSLATION (chksum(pcal) = "c0cb232" /\ chksum(tla) = "00000000")
VARIABLE pc

(* Allow infinite stuttering to prevent deadlock on termination. *)
Terminating == p/ Terminating

Spec == Init /\ [][Next]_vars

Termination == <>(pc = "Done")

\* END TRANSLATION 
====

