warning.J---- MODULE W4804_Test ----
(*
--algorithm Test {
}
*)
Spec == Init /\ [][Next]_vars

Termination == <>(pc = "Done")

\* END TRANSLATION 
====

