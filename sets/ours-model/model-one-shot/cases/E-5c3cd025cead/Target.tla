---- MODULE Target ----
EXTENDS Naturals
PLUS(x, y) == x + y

\* Higher‑order operator that expects a binary operator (arity 2)
HO(op(_,_)) == op(1, 2)

\* Supply the binary operator inside parentheses so that it is not a GeneralId node.
TEST == HO((PLUS))

====
