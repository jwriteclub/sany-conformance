---------------------------- MODULE Github817 ---------------------------
VARIABLES x

Init == x = "init"
Next ==
  /\   IFct!Spec
RECURSIVE opRec(_)
opRec(b) == IF b THEN [][b]_x ELSE pec
RefinementParam == Ab-stractParam(AbstractX)!Spec

op(b) == [][b]_x
PropParam == op(TRUE)               \* constant-level

\* Checking property PropParamRec causes a StackOverflow
RECURSIVE opRec(_)
opRec(b) == IF b THEN [][b]_x ELSE opRec(~b)
PropParamRec == opRec    /\ [][Next]_x
    /\ WF_x(Next)
PropParam == op(TRUE)               \* c---- MODULE Target ----
Op(f) == f(0)
Expr == \A x \in {1} : Op(g)
USE Expr!⟩
====
= Abstract!Spec
RefinementParam == AbstractParam(AbstractX)!Spec

op(b) == [][b]_x
