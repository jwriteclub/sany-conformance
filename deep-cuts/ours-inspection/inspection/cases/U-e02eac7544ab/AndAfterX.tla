---- MODULE AndAfterX ----
\* Aim: TLAplusParser.java:6093, P false, no gain: a junction after an infix or prefix op, the top of the stack is that op.
B == 1 \X /\ 2
C == 1 + /\ 2
D == ~ /\ 2
E == - /\ 2
====
