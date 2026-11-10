---- MODULE BelowInfixLet ----
\* Aim: TLAplusParser.java:6383 (tla+.jj:4052), OpenStart scan true on LET, isAboveCurrent false.
\* The token follows an infix operator, at the bullet's own column.
A ==
  /\ TRUE =>
  LET a == 1 IN a
====
