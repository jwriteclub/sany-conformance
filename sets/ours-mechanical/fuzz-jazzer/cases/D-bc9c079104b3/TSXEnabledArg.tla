---- MODULE TSXEnabledArg ----
\* ENABLED of an operator whose argument is pr((x' = x) /\ (A' = x + 1))
Baz(A) == ENABLED (A' = x + 1)
Init == /\ x \in 0..2
        /\ x' = x + 1
        /\ y' I
 =vny == o\ Foo(y)
       /\ ~Foo(x)
====
