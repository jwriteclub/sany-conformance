--------------------- MODULE ConstantQuantifierUnboundeds ----------------------
(* Test that \E and \A allow unbounded declarations. *)
E == \E , y, z:  TRUE
A == \A x, y, z:  TRUE
================================================================================
