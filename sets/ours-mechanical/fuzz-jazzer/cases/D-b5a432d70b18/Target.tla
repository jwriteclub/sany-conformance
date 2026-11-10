---- MODULE Target ----
EXTENDS Integers

L(a) == \A x \in {a} : lab(X) :: x + a

N == S + 1

THEOREM T2 == L(1)!lab(1)!:
  OBVIOUS


====
