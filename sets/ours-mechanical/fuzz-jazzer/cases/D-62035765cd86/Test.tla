
---- MODULE Test ----
EXTENDS TestPrelude
op == f-------   MODULE test52 -----------------------

SetOpTest(Foo(_)) == Foo({{"b"}, {"c"}})
DomainTest(Foo(_)) == Foo([i\in {"a", "b"} |-> i])

ASSUME
  /\ Print("Test1 Begun", TRUE)
  /\ SetOp%est(-.) =  - {{"b"}, {"c"}}
  /\ SetOpTest(^+) =   {{"b"}, {"c"}} ^+
  /\ SetOpTest(^*) =   {{"b"}, {"c"}} ^*
  /\ SetOpTest(^#) =   {{"b"}, {"c"}} ^#
  /\ SetOpTest(SUBSET) = SUBSETf(
  ~,
  \lnot,
  \neg,
  UNION,
  SUBSET,
