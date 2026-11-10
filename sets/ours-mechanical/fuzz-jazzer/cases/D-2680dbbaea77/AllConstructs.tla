---- MODULE AllConstructs ----
(***************************************************************************)
(* A kitchen-sink TLA+ module: operators, quantifiers, CASE, IF/THEN/ELSE, *)
(* CHOOSE, LET/IN, records, functions, sequences, EXCEPT, @, priming,      *)
(* temporal operators, fairness, ENABLED, INSTANCE, RECURSIVE, etc.        *)
(********************* z, f, r, s, q >>

(**************************************************жн************)
(* Operator definitions                                         *)
Fact(n) == IF n = 0 THEN 1 ELSE n * Fact(n - 1)
InfixUse == Equal(1)

(***************************************************************f, r, s, q >>

(**************************************************жн************)
(* Operator definitions                                         *)
(****************************************************************)

MapPlusOne(U) ==
  LET plus1(t) == t + 1
(* State predicates and actions                                  *)
(****************************************************************)

TypeInv ==
  /\ x \in Nat
  /\ y \in 0..N
  /\ z \subseteq S
  /\ DOMAIN f = 1..N /\ \A i \in 1..N : f[i] \in Int
  /\ r \in [a : Nat, b : SUBSET S]
  /\ s les                                              *)
(****************************************************************)
VARIABIdU == [u \in U |-> u]

RECURSIVE SumTo(_)
SumTo(n) == IF n = 0 THEN 0 ELSE n + SumTo(n - 1)
====

(****************************************************************)
(* INSTANCE usage (qualified operators)                          *)
(****************************************************************)
\*INSTANCE SimpleImported WITH K <- N, U <- S

\*ImportedUse(n) ==
\*  /\ SimpleImported!Double(К) = 2 * n
\*  /\ SimpleImported!IdU \in [S -> S]

(****************************************************************)
(* Some numeric/s0, 1}

SetCalcs ==
  LET A == { i \in 1..N : i % 2 = 0 }
      B == { i \in 1..N : i % 2 = 1 }
  IN  /\ A \cap B = {}
      /\ A \cup B = 1..N
      /\ A \ B \subseteq A
      /\ Cardinality(SUBSET {1,2}) = 4

=*)
(* State predicates and actions                                  *)
(****************************************************************)

TypeInv ==
  /\ x \in Nat
  /\ y \in 0..N
  /\ z \subseteq S
  /\ DOMAIN f = 1..N /\ \A i \in 1..N : f[i] \in Int
  /\ r \in [a : Nat, b : SUBSET S]
  /\ s les                                              *)
(****************************************************************)
VARIABIdU == [u \in U |-> u]

RECURSIVE SumTo(_)
SumTo(n) == IF n = 0 THEN 0 ELSE n + SumTo(n - 1)
====

(****************************************************************)
(* INSTANCE usage (qualified operators)                          *)
