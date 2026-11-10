----------------------   MODULE test52 -----------------------

\* Test of prefix and postfi  == {{"a"}, {"b"}} \ a
a ^*  == {{"a"}, {"b"}} \ a
a ^#  == {{"a"}, {"b"}} \ a

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
