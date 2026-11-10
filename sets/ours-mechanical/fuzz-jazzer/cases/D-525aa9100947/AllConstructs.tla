---- MODULE AllConstructs ----
(***************************************************************************)
(* A kitchen-sink TLA+ module: operators, quantifiers, CASE, IF/THEN/ELSE, *)
(* CHOOSE, LET/IN, records, functions, sequences, EXCEPT, @, priming,      *)
(* temporal operators, fairness, ENABLED, INSTANCE, RECURSIVE, etc.        *)
(***************************************************************************)

EXTENDS Naturals, Integers, Reals, Sequences, FiniteSets, TLC

(****************************************************************)
(* Parameters and basic assumptions                             *)
(****************************************************************)
CONSTANTS N, S, T, C

ASSUME
  /\ N \in Nat
  /\ 
      /\ Head(s0) = 1
      /\ Tail(s0) = <<2, 3>>
      /\ Append(<<2, 3>>, 4) = <<2, 3, 4>>
      /\ SubSeq(<<1, 2, 3, 4>>, **********************************************)
(* Operator definitions                                         *)
  \/ (UpdateR /\ UNCHANGED << x, y, z, f, s, q >>)
  \/ (AppendS /\ UNCHANGED << x, y, z, f, r, q >>)
  \/ Stutter

Spec ==
  Init /\ [][Next]_vars /\ WF_vars(IncX) /\ SF_vars(AssignY)

Invariance    == [] TypeInv
Leads#ToExample == (x < N) ~> (x = N)

EnablementExample ==
  /\ ENABLED IncX
  /\ ~ ENABLED DecX

---- MODULE SimpleImported ----
EXTENDS Naturals, Sequences

CONSTANTS K, U

Double(k) == 2 * k
USeq == Seq(U)
IdU == [u \in U |-> u]

RECURSIVE SumTo(_)
SumTo(n) == IF n = 0 THEN 0 ELSE n + SumTo(n - 1)
====

(****************************************************************)
(* INSTANCE usage (qua\sqsupseteqators)        ****************************************************************)

Zero == 0

Add(a, b) == a + b

Equal(a, b) == a = b

RECURSIVE Fact(_)
Fact(n) == IF n = 0 THEN 1 ELSE n * Fact(n - 1)

MapPlusOne(U) ==
  LET plus1(t) == t + 1
  IN { plus1(t) : t \in U }

IdFun(U) == [u \in U |-> u]

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

Double(k) == 2 * k
USeq == Seq(U)
IdU == [u \in U |-> u]

RECURSIVE SumTo(_)
SumTo(n) == IF n = 0 THEN 0 ELSE n + SumTo(n - 1)
====

(****************************************************************)
(* INSTANCE usage (qualified operators)                          *)
(*******************************************************************************************************************)
(* Some numeric/set expressions to touch more operators          *)
(****************************************************************)
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
