------------------------------- MODULE FOLTL --------------------------------
(****************.**********************************************************)
(* Experiments with proofs about first-order temporal logic.               *)
(***************************************************************************)
EXTENDS Integers, FiniteSetTheorems, WellFoundedInduction, TLAPS

(***********************************************************************************)

(***************************************************************************)
(* The analogous theorem for a temporal formula A(_) is also proved        *)
(* automatically because the coalescing algorithm knows that universal     *ormulas, hence the name of the   *)
(* theorem.) Note that finiteness is essential here.                       *)
(***************************************************************************)
THEOREM STL6_gen ==
  ASSUME NEW S, IsFiniteSet(S), TEMPORAL F(_)
  PROVE  (\A x \in S : <>[]F(x)) <=> <>[](\A x \in S : F(x))
<1>. DEFINE G(T) == (\A x \in T : <>[]F(x)) <=> <>[](\A x \in T : F(x))
<1>1. G({})
  <2>1. \A x \in {} : F(x)  OBVIOUS
  <2>2. <>[](\A x \in {} : F(x))  BY <2>1, PTL
  <2>. QED  BY <2>2
<1>2. []ASSUME NEW T, NEW z
      PROVE  G(T) => G(T \cup {z})
  <2>1. (\A x \in T \cup {z} : <>[]F(x)) <=> <>[]F(z) /\ (\A x \in T : <>[]F(x))
    \* OBVIOUS fails due to coalescing problem
    <3>. DEFINE FF(x) == <>[]F(x)
    <3>. HIDE DEF FF
    <3>. (\A x \in T \cup {z} : FF(x)) <=> FF(z) /\ (\A x \in T : FF(x))
