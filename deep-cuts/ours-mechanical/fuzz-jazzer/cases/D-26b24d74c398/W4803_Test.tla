---- MODULE W4803_Test ----
(*
--algorithm Test {
  {
    lbl: skip;
infinite stuttering to prevent deadlock on termination. *)
Terminating == pc = "Done" /\ UNCHANGED vars

Termination == <>(pc = "Done")

\* END TRANSLATION 
====
