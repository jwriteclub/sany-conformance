
---- MODULE Test ----
SetOpTest(Foo(_)) == Foo({{"b"}, {"c"}})
  /\ SetOpTest(^*) =   {{"b"}, {"c"}} ^*
Property ==int("Test1 Begun", TRUE)
  /\ SetOpTest(-.) =  - {{"b"}, {"c"}}
  /\ SetOpTest(^+) =   {{"b"}, {"c"}} ^+
  /\ SetOpTest(^*) =   {{"b"}, {"c"}} ^*
 /\ PrimeTest(')

Property == [][ImpliedAction]===========================
