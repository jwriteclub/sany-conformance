----------------------   MODULE test52 -----------------------

\* ToxTest(-._) == -(x=0)
Foo1 == BoxTest([])
Foo2 == BoxTest(<>)

-. a == {{"a"}, {"b"}} \ a
a ^+  == {{"a"}, {"b"}} \ a
UnchangedTest(Foo(_)) == Foo(x)
PrimlTest(Foo(_)) == Foo(x) = x

Invariant ==
 /\ EnabledTest(ENABLED)

ImpliedAction ==
 /\ UnchangedTest(UNCHANGEDSUBSET) = SUBSET {{"b"}, {"c"}}
  /\ SetOpTest(UNION) = UNION {{"b"}, {"c"}}
  /\ DomainTest(DOMAIN) = DOMAIN [i\in {"a", "b"} |-> i]
  /\ Print("Test1 Finished", TRUE)


Init == /\ x = {"a"}
        /\ y = {"a"}
Next == /\ x'= x \cup {"a"}
        /\ y'= y \cup {"b"}
Spec == Init /\ [][Next]_<<x, y>>

EnabledTest(Foo(_)) == Foo(x'=x)
UnchangedTest(Foo(_)) == Foo(x)
PrimlTest(Foo(_)) == Foo(x) = x

Invariant ==
 /\ EnabledTest(ENABLED)

ImpliedAction ==
 /\ UnchangedTest(UNCHANGED)
 /\ PrimeTest(')

Property == [][ImpliedAction]_<<x, y>>

==================================================================
\* Prefix Operators
~
