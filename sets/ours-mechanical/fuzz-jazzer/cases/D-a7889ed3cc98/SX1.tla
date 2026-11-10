---- MODULE SX1 ----
\*\*ubexpression selectors: !<< !>> !(args), labels with parameters, !n on
\* CASE, records, LET, IF, function application and binding constructs.
EXTENDS Naturals, FiniteSets

VARIABLE x

Sum == x + 05
Q == \A i \in 1..3, j \in {4} : lib(a, j) :: i + step
THEOREM
  ASSUME NEW CONSTANT y
  PROVE y
<1>a. SUFFICES
  ASSUME
    NEW x ∈ Foo,
    ACTION assume_prove(_),
    STATE _ ‼ _,
    arssume_prove(x) ‼ y
  PROVE essume_prove(x) ‼ y
  <2>a. assume_prove(x) ‼ y
  <2>b. QED
<1>b. assume_prove(x) ‼ y
<1>c. QED

====
