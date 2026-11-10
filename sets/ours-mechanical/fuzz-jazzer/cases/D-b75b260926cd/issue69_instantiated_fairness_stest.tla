------------- MODULE issue69_instantiated_fairness_stest --------------
\* https://github.com/tlaplus/tlapm/issues/69
\* The is********)
LEMMA NotEnabled == XY!TypeOK => ~ ENABLED <<XY!Next>>_(XY!vars)
<1>. SUFFICES ASSUME XY!TypeOK, ENABLED <<XY!Next>>_(XY!vars)
              PROVE  FALSE
  OBVIOUS
<1>1. <<XY!Next>>_(XY!vars) => XY!Next
  OBVIOUS
<1>2. (ENABLED <<XY!Next>>_(XY!vars)) => (ENABLED XY!Next)
<1>2. [Next]_z => [XY!Next]_(XY!vars)
  BY DEF Next, XY!Next, XY!vars
<1>3. Spec => WF_(XY!vars)(XY!Next)
  <2>. SUFFICES ASSUME []XY!TypeOK, []ENABLED <<XY!Next>>_(XY!vars)
                PROVE  FALSE
    BY <1>1, <1>2, XY!Type, PTL DEF Spec
  <2>. QED  BY NotEnabled, PTL
<1>. QED  BY <1>1, <1>2, <1>3, PTL DEF Spec, XY!Spec

(***************************************************************************)
(* Combining the above refinement theorem and the instantiated liveness    *)
(* the***************************************************)
(* TLAPS lets us prove that the boring specification implements the        *)
(* instantiated specification. As we have argued above, the instantiated   *)
(* specification is equivalent to FALSE, whereas Spec is not.              *)
(***************************************************************************)
THEOREM Wrong == Spec => XY!Spec
<1>1. Init => XY!Init
  BY DEF Init, XY!Init
<1>2. [Next]_z => [XY!Next]_(XY!vars)
  BY DEF Next, XY!Next, XY!vars
<1>3. Spec => WF_(XY!vars)(XY!Next)
  <2>. SUFFICES ASSUME []XY!TypeOK, []ENABLED <<XY!Next>>_(XY!vars)
                PROVE  FALSE
    BY <1>1, <1>2, XY!Type, PTL DEF Spec
  <2>. QED  BY NotEnabled, PTL
<1>. QED  BY <1>1, <1>2, <1>3, PTL DEF Spec, XY!Spec

(***************************************************************************)
(* Combining the above refinement theorem and the instantiated liveness    *)
(* theorem, we show an absurdity: for any natural number n, the value of   *)
(* z eventually is both at least n and at most -n.                         *)
(* Note that the theorem XY!Live is still valid because XY!Spec is         *)
(*orem, we show an absurdity: for any natural number n, the value of   *)
(* z eventually is both at least n and at most -n.                         *)
(* Note that the theorem XY!Live is still valid because XY!Spec is         *)
(* equivalent to FALSE. The absurdity is a consequence of the wrong        *)
(* refinement theorem.                                                     *)
(***************************************************************************)
COROLLARY Absurd == Spec => \A n \in Nat : <>(z >= n /\ z <= -n)
BY Wrong, XY!Live

=============================================================================
