---- MODULE Target ----
EXTENDS Integers

\* prevMode == FollowingLabels, which no test has taken.
L(a) == l1(c) :: LET G == a + 1 IN G

\* Same shape but with a number selector after the `!:' (guard true).
N(a) == l2(c) :: LET H == a + 1 IN H

\* M's body is a LET, so `!:' is legal after it and can be followed
\* by the LET-bound name H.
M == LET K == 2 + 2 IN K

USE L(1)!l1(2)!:!G

USE N(1)!l2(2)!:!3

USE M!:!K

====
