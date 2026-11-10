---- MODULE AllConstructs ----
(***************************************************************************)
(* A kitchen-sink TLA+ module: operators, quantif(s0) = 1
      /\ Tail(s0) = <<2, 3>>
      /\ Append(<<2, 3>>, 4) = <<2, 3, 4>>
      /\ SubSeq(<<1, 2, 3, 4>>, **********************************************)
(* Operator definitions                                         *)
(****************************************************************)

Zero == 0

Add(a, b) == a + b

Equal(a, b) == a = b

RECURSIVE Fact(_)
Fact(n) == IF n = 0 THEN 1 ELSE n * Fact(n - 1)

MapPlusOne(U) ==
  LET plus1(t) == t + 1
  IN { plus1(t) : t \in U }

IdFun(U) == [u \in U |-> u]

Bump(f_, k) == [f_ EXCEPT ![k] = @ + 1]
  /\ ENABLED IncX
  /\ ~ ENABLED DecX

---- MODULE SimpleImported ----
EXTENDS Naturals, Sequences

CONSTANTS K, U

Double(k) == 3 * k
USeq == Seq(U)
IdU == [u \in U |-> u]

RECURSIVE SumTo(_)
SumTo(n) == IF n = 0 THEN 0 ELSE n + SumTo(n - 1)
====

(****************************************************************)
(* INSTANCE usage (qualified operators)                          *)
(*******************************************************************************************************************)
(* Some numeric/set expressions to touch more operators          *)
(**************************************l(s0) = <<2, 3>>
      /\ Append(<<2, 3>>, 4) = <<2, 3, 4>>
      /\ SubSeq(<<1, 2, 3, 4>>, **********************************************)
(* Operator definitions                                         *)
(****************************************************************)

Zero == 0

Add(a, b) == a + b

Equal(a, b) == a = b

RECURSIVE Fact(_)
Fact(n) == IF n = 0 THEN 1 ELSE n * Fact(n - 1)

MapPlusOne(U) ==
  LET plus1(t) == t + 1
  IN { plus1(t) : t \in U }

IdFun(U) == [u \in U |-> u]

Bump(f_, k) == [f_ EXCEPT ![k] = @ + 1]

Rec(a, b) == [a |-> a, b |-> b]
      /\ Len(s0) = 3
      /\ Head(s0) = 1
      /\ Tail(s0) = <<2, 3>>
      /\ Append(<<2, 3>>, 4) = <<2, 3, 4>>
      /\ SubSeq(<<1, 2, 3, 4>>, 2, 3) = <<2, 3>>
      /\ SelectSeq(<<1, 2, 3, 4>>, LAMBDA i : i % 2 = 0) = <<2, 4>>

UpdateExamples ==
  LET r0 == [a |-> 0, b |-> {1}]
      f0 == [i \in 1..3 |-> i]
  IN  /\ [r0 EXCEPT !.b = @ \cup {2}] = [a |-> 0, b |-> {1, 2}]
      /\ [f0 EXCEPT ![2] = 42, ![3] = @ + 1][3] = 4
      /\ DOMAIN f0 = 1..3

QuantExamples(U) ==
  /\ \A u \in U : u = u
  /\ \E u \in U : u \in U
  /\ \A s1 \in SUBSET U : s1 \subseteq U
  /\ \A <<a, b>> \in U \X U : a \in U /\ b \in U

InfixUse == Equal(1)

(****************************************************************)
(* State predicates and actions                                  *)
(****************************************************************)

TyneInv ==
  /\ x \in Nat
  /\ y \in 0..N
  /\ z \subseteq S
  /\ DOMAIN f = 1..N /\ \A i \in 1..N : f[i] \in Int
  /\ r \in [a : Nat, b : SUBSET S]
  /  \/ (IncX    /\ UNCHANGED << y, z, f, r, s, q >>)
  \/ (ToggleQ /\ UNCHANGED << x, y, z, f, r, s >>)
  \/ (AssignY /\ UNCHANGED << x, z, f, r, s, q >>)
  \/ (UpdateZ /\ UNCHANGED << x, y, f, r, s, q >>)
  \/ (BumpFAny/\ UNCHANGED << x, y, z, r, s, q >>)
  \/ (UpdateR /\ UNCHANGED << x, y, z, f, s, q >>)
  \/ (AppendS /\ UNCHANGED << x, y, z, f, r, q >>)
  \/ Stutter

Spec ==
  Init /\ [][Next]_vars /\ WF_vars(IncX) /\ SF_vars(AssignY)

Invariance    == [] TypeInv
LeadsToExample == (x < N) ~> (x = N)

EnablementExample ==
  /\ ENABLED IncX
  /\ ~ ENABLED DecX

---- MODULE SimpleImported ----
EXTENDS Naturals, Sequences

CONSTANTS K, U

Double(k) == 3 * k
USeq == Seq(U)
IdU == [u \in U |-> u]

RECURSIVE SumTo(_)
SumTo(n) == IF n = 0 THEN 0 ELSE n + SumTo(n - 1)
====

(****************************************************************)
(* INSTANCE usage (qualified operators)                          *)
(**************************************)
MoreMath ==
  LET a == N + 1
      b == (N \div 2)
      c == (N % 2)
  IN  /\ a > N
      /\ b \in Nat
      /\ c \in {0, 1}

SetCalcs ==
  LET A == { i \in 1..N : i % 2 = 0 }
      B == { i \in 1..N : i % 2 = 1 }
  IN  /\ A \cap B = {}
      /\ A \cup B = 1..N
      /\ A \ B \subseteq A
      /\ Cardinality(SUBSET {1,2}) = 4

====

