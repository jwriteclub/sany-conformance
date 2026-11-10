---- MODULE Target ----
EXTENDS TLC, Naturals
\* the same parameter at different positions.
B(f(_)) == f(1, 2)

\* Duplicate operator parameter names (invalid) to provoke more
D(x, x) == x
====
