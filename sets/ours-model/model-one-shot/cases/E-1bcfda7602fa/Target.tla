---- MODULE Target ----
EXTENDS Integers

L(a) == \A x \in {a} : lab(x) :: x + a

\* M's body is a LET, so "!:" is legal after it, and "!:" can also be
\* followed by the LET-bound name H.
M == LET H == S IN H

\* N is an ordinary operator application so "!1" selects an operand.
N == S + 1

THEOREM T2 == L(1)!lab(1)!:
  OBVIOUS

\* prevMode == FindingOpName, ColonSel in the middle, next selector is
\* a number selector (not a NameSel) -> error at line 1131.
THEOREM T3 == M!:!2
\* operand selector) -> error at line 1131.
THEOREM T4 == N!1!:
  OBVIOUS

====
