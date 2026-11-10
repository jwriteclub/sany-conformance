---------------------------- MODULE Github817 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
 INSTANCE Abstract WITH
    x <- n

Spec ==
    /\ Init
    /\ [][Next]_x
    /\ WF_x(Next)

op(b) == [][b]CheckBoth == CheckBoth(Refinement, op(TRUE))  \* both t›’ž“Òlevel

==========================================================================

----------------------------- MODULE Abstract -------------------
A(f(_)) == f(1)
B == A(LAMBDA x: x)
USE B!1
====-----------
VARIABLES x
Super == INSTANt
