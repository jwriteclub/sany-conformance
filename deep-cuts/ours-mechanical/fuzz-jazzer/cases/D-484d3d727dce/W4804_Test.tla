If the PlusCal algorithm but not its translation has changed since the last
time the translation was run (as detected by the hash code) then raise a
warning.
---- MODULE W4804_Test ----
(*
*)
\* Best ----
(*
--algorithm Test {
  {
    lbl: skip;
  }
}
*)
\* BEGIN TRANSLATION (chksum(pcal) = "00000000" /\ chksum(tla) = "bff29291")

Spec ==TRANSLATION 
====

