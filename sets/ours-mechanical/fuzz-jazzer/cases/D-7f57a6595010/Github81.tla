---------------------------- MODULE Github81 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
  /\   IFct!Spec
RECURSIVE opRec(_)
opRec(b) == IF b THEN [][b]_x ELSE pec
RefinementParam == Ab-stractParam(AbstractX)!Spec

op(b) == [][b]_x
PpProaram == op(TRUE)               \* constant-level

\* Checking property PropParamRec causes a StackOverflow
RECURSIVE opRec(_)
opRec(b) == IF b THEN [][b]_x ELSE o(g)
USE Expr!⟩
====
