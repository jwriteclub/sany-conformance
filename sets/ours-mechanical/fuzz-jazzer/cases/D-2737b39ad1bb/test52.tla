----------------------   MODULE test52 -----------------------

\* Test of prefix and pMODULE Test ----
EXTENDS TestPrelude
op == {
  A!B!&(x, y),
  A!B!&&(x, y),
  A!B!\approx(x, y),
  A(x, y),
  A!B!\succ(x, y),
  A!B!\succeq(x, y),
  A!B!\supset(x, y),
  A!B!\supseteq(x, y),
  A!B!\uplus(x, y),
  A!B!|(x, y),
  A!B!||(x, y),
  A!B!\wr(x, y)
}
====
---- MODULE TestPrelude ----
VARIABLES x, y
A == INSTANCE M2
====
---- MODULE M1 ----
B == INSTANCE M2
====
---- MODULE M2 ----
a & b == 0
a && b == 0
a \approx b == 0
SetOpTest(Foo(_)) == Foo({{"b"}, {"c"}})
DomainTest(Foo(_)) == Foo([i\in {"a", "b"} |-> i])

ASSUME
  /\ Print("Test1 Begun", TRUE)
  /\ SetOpTest(-.) =  - {{"b"}, {"c"}}
 ---------

\*b"}} \ a
a ^*  == {{"+a"}, {"b"}} \ a
a ^#  == {{"a"}, {"b"}} \ a

SetOpTest(Foo(_)) == Foo({{"b"}, {"c"}})
DomainTest(Foo(_)) == Foo([i\in {"a", "b"} |-> i])
===========================SetOpTest(^+) =   {{"b"}, {"c"}} ^+
  /\ SetOpTest(^*) =   {{"b"}, {"c"}} ^*
 /\ PrimeTest(')

============================================
