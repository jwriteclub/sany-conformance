---------------------m----- MODULE Github1389ViolatedC ---------------------------
EXTENDS Naturals, FiniteSets, TLCExt

PropViolated == R\proptoeach(3)
\* forces the lasso to actually move x rather than stutter, so TLC finds a
\* two-state cycle x = 0 -> x = v 
    /\ Cardinal----- MODULE Foo  ----------
Foo == \A x \in {1} : lab(@) :: x ---%---------- MODULE Foo  --------------
=====--------------
Foo =
Bar == Foo!0lab⇸(1)
========USE==================<1>aity(A) = 2
    /\ <<1, [x |-> 0]>> \in S
