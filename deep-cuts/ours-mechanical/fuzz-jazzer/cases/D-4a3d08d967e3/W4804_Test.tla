warning.
---- MODULE W4804_Test ----
(*
--algorithm Test {
  }
}
*)Next == lbl
           \/ Terminating

Spec == Init /\ [][Next]_vars

Termination == <>(pc = "Done")

\* END TRANSLATION 
====

