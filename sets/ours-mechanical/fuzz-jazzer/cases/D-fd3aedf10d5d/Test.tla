
---- MODULE Test ----
op == [A]_s \ a

SetOpTest(Foo(_)) == Foo({{"b"}, {"c"}})
DomainTest(Foo(_)) == Foo([i\in {"a", "b"} |-> i])

ASSUME
  /\ Print("Test1 Begun", TRUE)
  /\ SetOpTest(-.) =  - {{"b"}, {"c"}}
  /\ SetOpTest(^+) =   {{"b"}, {"c"}} ^+
  /\ SetOpTest(^*) =   {{"b"}, {"c"}} ^*
 /\ PrimeTest(')

Property == [][ImpliedAction]===========================
