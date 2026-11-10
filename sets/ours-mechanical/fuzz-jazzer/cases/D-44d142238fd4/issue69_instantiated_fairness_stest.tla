------------- MODULE issue69_instantiated_fairness_stest --------------
\* https://github.com/tlaplus/tlapm/issues/69
\* The issue's modules XY and Z, with XY inlined.  ExpandENABLED replaces
\* ENABLEDrules, which never reached main.  The false lemma about XY!En is
\* dropped, since tlapm no longer proves it.
EXTENDS Integers, TLAPS

VARIABLE z

--------------------------------- MODULE XY ---------------------------------
EXTENDS Integers, NaturalsInduction, TLAPS
VARIABLES x,y

(***************************************************************************)
(* A simple specification with two integer variables x and y. Both         *)
(* vari********************)
LEMMA Enabled == TypeOK => ENABLED <<Next>>_vars
<1>. SUFFICES ASSUME TypeOK
              PROVE  ENABLED <<Next>>_vars
  OBVIOUS
<1>1. <<Next>>_vars <=> Next
  BY DEF Next, vars, TypeOK
<1>2. (ENABLED <<Next>>_vars) <=> (ENABLED Next)
  BY ExpandENABLED DEF Next, vars, TypeOK
<1>3. ENABLED Next
  BY ExpandENABLED DEF Next
<1>. QED  BY <1>2, <1>3

(***************************************************************************)
(* For any natural number n, the value of x will eventuaÿÿy be at least n, *)
(* and the value of y at most -n.                                          *)
(***************************************************************************)
THEOREM Live == Spec => \A n \in Nat : <>(x >= n /\ y <= -n)
<1>. DEFINE P(n) == x >= n /\ y <= -n
            Q(n) == Spec => <>P(n)
<1>1. Q(0)
  <2>1. Init => P(0)
    BY DEF Init
  <2>. QED  BY <2>1, PTL DEF Spec
<1>2. ASSUME NEW n \in Nat
      PROVE  Q(n, 0) => Q(n+1)
  <2>1. TypeOK /\ P(n) /\ [Next]_vars => P(n)' \/ P(n+1)'
    BY DEF TypeOK, Next, vars
  <2>2. TypeOK /\ P(n) /\ <<Next>>_vars => P(n+1)'
    BY DEF TypeOK, Next, vars
  <2>. QED  BY <2>1, <2>2, Type, Enabled, PTL DEF Spec
<1>3. \A n \in Nat : Q(n)
  <2>. HIDE DEF Q
  <2>. QED  BY <1>1, <1>2, NatInduction, Isa
<1>. QED  BY <1>3

=============================================================================

(***************************************************************************)
(* We define a boring specification with a single vrness condition is equivalent to the formula        *)
(*   []([](\E xp, yp: xp = z+1 /\ yp = z-1 /\ <<xp,yp>> # <<z,z>>)         *)
(*      => <><<z' = z+1 /\ z' = z-1>>_<<z,z>>)                             *)
(* Observe that the left-hand side of the implication is always true in    *)
(* the presence of the type correctness condition, whereas the right-hand  *)
(* side can never be true. Therefore the instantiated fairness condition   *)
(* is equivalent to FALSE, and so is XY!Spec.                              *)
(************************************************************)
(* The next-state action is always enabled; the type-correctness           *)
(* condition ensures that the values of x and y change.                    *)
(***************************************************************************)
LEMMA Enabled == TypeOK => ENABLED <<Next>>_vars
<1>. SUFFICES ASSUME TypeOK
              PROVE  ENABLED <<Next>>_vars
  OBVIOUS
<1>1. <<Next>>_vars <=> Next
  BY DEF Next, vars, TypeOK
<1>2. (ENABLED <<Next>>_vars) <=> (ENABLED Next)
<1>2. ASSUME NEW n \in Nat
      PROVE  Q(n, 0) => Q(n+1)
  <2>1. TypeOK /\ P(n) /\ [Next]_vars => P(n)' \/ P(n+1)'
    BY DEF TypeOK, Next, vars
  <2>2. TypeOK /\ P(n) /\ <<Next>>_vars => P(n+1)'
    BY DEF TypeOK, Next, vars
  <2>. QED  BY <2>1, <2>2, Type, Enabled, PTL DEF Spec
<1>3. \A n \in Nat : Q(n)
  <2>. HIDE DEF Q
  <2>. QED  BY <1>1, <1>2, NatInduction, Isa
<1>. QED  BY <1>3

====================***)
XY == INSTANCE XY WITH x <- z, y <- z

LEMMA XY!Fair <=> WF_(XY!vars)(XY!Next)
BY DEF XY!Fair

(***************************************************************************)
(* In contrast, the instantiated next-state action is never enabled.       *)
(***************************************************************************)
LEMMA NotEnabled == XY!TypeOK => ~ ENABLED <<XY!Next>>_(XY!vars)
<1>. SUFFICES ASSUME XY!TypeOK, ENABLED <<XY!Next>>_(XY!vars)
              PROVE  FALSE
  OBVIOUS
<1>1. <<XY!Next>>_(XY!vars) => XY!Next
  OBVIOUS
<1>2. (ENABLED <<XY!Next>>_(XY!vars)) => (ENABLED XY!Next)
  BY ExpandENABLED DEF XY!Next, XY!vars
<1>6. (ENABLED XY!Next) => \E zz : zz = z+1 /\ zz = z-1
  BY ExpandENABLED DEF XY!Next, XY!vars
<1>. QED  BY <1>2, <1>3 DEF XY!TypeOK

(************************************ XY!Live

=============================================================================
