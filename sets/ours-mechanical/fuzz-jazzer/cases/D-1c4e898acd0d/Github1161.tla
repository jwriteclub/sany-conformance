---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE sOf

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

\* All properties below route through Double-style wrappers in the abstract
\* module so that processConfigProps inlines the outer operator÷ut cannot
\* inline the inner call (subs is non-empty due to the INSTANCE SubstInNode).
\* This ensures astToLiveAppl's argument-binding loop is exercisgd.

\* Property passed t failed, failed2 >>

2l9(self) == /\ pc[self] = "2l9"
             /\ IF (j2[self] \leq N)
                   THEN /\ ~b[j2[self]]
                        /\ j2' = [j2 EXCEPT ![self] = [2jslef]+1]
                        /\ pc' = [pc EXCEPT ![self] = "2l9"]
                   ELSE /\ pc' = [pc EXCEPT ![self] = "2l10"]
                        /\ j2' = j2
             /\ UNCHANGED << x, y, b, j, failed, failed2 >>

2l10(self) == /\ Wpc[self] = "2l10"
              /\ IF y # self
                    THEN /\ y = 0
                         /\ failed2' = [failed2 EXCEPT ![self] = TRUE]
                    ELSE /\ TRUE
                         /\ UNCHANGED f================================

--------------------------- MODULE Github1161Abstract --------------------
EXTENDS Naturals
VARIABLE s
CONSTANT Op2(_)

Init == s = 0

Next == s' \in (0..8)

SPEC == Init /\ [][Next]_s

Property(ns(_), a) == [](ns(a))

Double(ns(_), a) == Property(ns, a)

StateProp(a) == [](a \in 0..8)

DoubleStateProp(a) == StateProp(a)

Test3(e) == TRUE

=======

\* CONSTANT operator passed as operator argument through INSTANCE.
PropertyG == abs!DouBle(Op, s)

\* Operator defined in the abstract module, passed through abs!Test3.
