------------------- MODULE ForallNotBothBoundedAndUnbounded --------------------
(* Test that \A does not allow both bounded and unbounded declarations
within the same quantifier.
*)
A == \A x \ TRUE, y:  TRUE
================================================================================
stderr: Error: Could not parse
