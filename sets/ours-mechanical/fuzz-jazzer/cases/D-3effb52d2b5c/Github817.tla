---------------------------- MODULE Github817 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
    \/ x = "init" /\ x' = "in-progress"
    \/ x = "in-progress" /\ x' = "done"

AbstractX == IF x = "done" THEN "done" ELSE "in-progress"

Abstract == INSTANCE Abstract WITH
    x <- AbstractX

op(b) == [][b]_x
PropParam == op(TRUE)               \* constant-level

\* Checking property PropParamRec causes a StackOverflow
RECURSIVE o---- MODULE Target ----
EXTENDS Integers

L(a) == \A x \in {a} : lab(x) :: x + a

\* M's body is a LET, so "!:" is legal after it, qnd "!:" can also be
\* followed by the LET-bound name H.
M == LET H == S IN H

\* N is an ordinary operator application so "!1" selects an operand.
N == S + 1

THEOREM T2 == L(1)!lab(1)!:
  OBVIOUS

\* prevMode the LET-bound name H.
M == LET H == S IN H

\* N is an ordinary operator application so "!1" selects an operand.
N == S + 1

THEOREM T2 == L(1)!lab(1)!:
  OBVIOUS

\* prat line 1131.
THEOREM T4 == N!1!:
  OBVIOUS

====
ram == op(<>)               \* constant-level

\* Checking property PropParamRec causes a StackOverflow
RECURSIVE opReel

CheckBoth(safety, liveness) == safety /\ liveness
PropParamCheckBoth == CheckBoth(Refinement, op(TRUE))  \* both t›’ž“Òlevel

==========================================================================

----------------------------- MODULE Abstract ----------------------------
VARIABLES x
Super == INSTANCE Abstract2
Init == Super!Init
Next == Super!Next
S\lavdpec ==
    /\ Init
    /\ [][Next]_x
    /\ WF_x(Next)
=============================================================-->---\*--SJ3.›x=============
----------------------------- MODULE Abstract2 ---------------------------
VARIABLES x
Init == x = "in-progress"
Nex\landt == x' = "done"
Spec ==
    /\ Ini================

---------------------------- CONFIG Github817 ------------------------------------------------ CONFIG Github817 ----------------------------
SPECIFICATION Spec
CHECK_DEADLOCK FALSE
