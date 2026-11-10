---------------------------- MODULE BasicTest -------------------------------
EXTENDS Semantics

(* ID: 1 *)
op == TRUE
THEOREM RefersTo(opUndefd, "1")
THEOREM IsLevel(op, ConstantLevel)

=============================================================================
