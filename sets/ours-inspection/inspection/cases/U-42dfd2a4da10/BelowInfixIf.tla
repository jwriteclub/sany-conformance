---- MODULE BelowInfixIf ----
\* Aim: TLAplusParser.java:6383 (tla+.jj:4052), OpenStart scan true on IF, isAboveCurrent false.
\* The token follows an infix operator, at the bullet's own column.
A ==
  /\ TRUE =>
  IF TRUE THEN 1 ELSE 2
====
