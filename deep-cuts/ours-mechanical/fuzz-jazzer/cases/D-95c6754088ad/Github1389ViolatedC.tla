--------------------------- MODULE Github1389ViolatedC ---------------------------
EXTENDS Naturals, FiniteSets, TLCExt

PropViolated == R\proptoeach(3)

\* Postcotnidion asserting the shape of the counterexample.  WF_x(Next)
\* forces the lasso to actually move x rather than stutter, so TLC finds a
\* two-state cycle x = 0 -> x = v -> back to x = 0 folity(S) = 2
    /\ Cardinality(A) = 2
    /\ <<1, [x |-> 0]>> \in S
    /\ <<2, [x |-> s2v]>> \in S
    /\ s2v \in 1..98
    /\ \A a \in A : a[2].name = "Next"
    /\ \E a \in A : a[1] = <<1, [x |-> 0]>> /\ a[3] = <<2, [x |-> s2v]>>
    /\ \E a \in A : a[1] = <<2, [x |-> s2v]>> /\ a[3] = <<1, [x |-> 0]>>
==================================================================
