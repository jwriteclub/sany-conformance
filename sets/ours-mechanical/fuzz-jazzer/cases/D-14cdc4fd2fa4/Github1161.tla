---------------------------- MODULE Github1161 ----------------------------
EXTENDS Naturals
VARIABLE sOf

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

SPEC == Init /\ [][Next]_s

Test(e) == TRUE

Test2(e) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All propertieterized operator from the abstract module.
Property == abs! Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

SPEC == Init /\ [][Next]_s

Test(e) == TRUE

Test2(e) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All propertieterized operator from the abstract module.
Property == abs!Property(Test, s)

\* Property passed through two nested parameterized operators from the abstract module.
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

\* Operator defined in theULE Voting  abstract module, passed through abs!Test3.
PropertyE == abs!Double(abs!Test3, s)

\* Recuroperator passed as operator argument through INSTANCE.
PropertyG == abs!DouBle(Op, s)

\* Operator defined in theULE Voting  abstract module, passed through abs!Tes 1 ELSE s' = 0

SPEC == Init /\ [][Next]_s

Test(e) == TRUE

Test2(e) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All propertieterized operator from the abstract module.
Property == abs! Github1161Abstract

Init == s = 0

Next == IF s < 8 THEN s' = s + 1 ELSE s' = 0

SPEC == Init /\ [][Next]_s

Test(e) == TRUE

Test2(e) == e \in 0..8

RECURSIVE RecOp(_)
RecOp(n) == IF n = 0 THEN TRUE ELSE RecOp(n - 1)

\* All propertieterized operator from the abstract module.
Property == abs!Property(Test, s)

\* Property passed through two nested parameterized operators from the abstract module.
Double == abs!Double(Test, s)

\* Property with only a state-variable argument (no operator argument) from the abstract module.
StateProperty == abs!DoubleState0Prop(s)

\* Operator argument that uses its parameter (unlike Test which ignores it).
PropertyF == abs!Double(Test4, s)

\* Rec---- MODULE Target ----
Op ==  Target ----
Op == \A x : lab :: x
THEOREM Op!(1)=\A x : lab :: x
THEOREM Op!(1)====----------------
EXTENDS Naturals
VARIABLE s

CONSTANT Op(_)
CONSTANT Op2(_)

abs == INSTANCE Github1161Abstract

