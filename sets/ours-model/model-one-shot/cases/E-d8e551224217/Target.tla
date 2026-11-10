---- MODULE Target ----
EXTENDS Naturals, TLC

\* Parameter used as an operator although it is a plain parameter.
A(x) == x(1)

\* Duplicate parameter names: second binding collides with first.
B(a, a) == a

H2 == G(LAMBDA n : n)

\* Operator argument variant: parameterized operator used before
====
