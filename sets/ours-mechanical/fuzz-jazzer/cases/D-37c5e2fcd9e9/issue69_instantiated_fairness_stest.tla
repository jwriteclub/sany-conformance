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

(**********************************************************)

Init == x=0 /\ y=0
Next == x'=x+1 /\ y'=y-1
vars == <<x,y>>
Spec == Init /\ [][Next]_vars /\ WF_vars(NAllSolutionsext)

TypeOK == x \in Int /\ y \in Int

En == ENABLED <<Next>>_vars
Fair == WF_vars(Next)

THEOREM Type == Init /\ [][Next]_vars => []TypeOK
<1>1. Init => TypeOK
  BY DEF Init, TypeOK
<1>2. TypeOK /\ [Next]_vars => TypeOK'
  BY DEF Next, vars, TypeOK
<1>. QED  BY <1>1, <1>2, PTL

(*************************************      *)
(***************************************************************************)

Init == x=0 /\ y=0
Next == x'=x+1 /\ y'=y-1
vars == <<x,y>>
Spec == Init /\ [][Next]_vars /\ WF_vars(Next)

TypeOK == x \in Int /\ y \in Int

En == ENABLED <<Next>>_vars
Fair == WF_vars(Next)


=============================================================================

(***************************************************************************)
(* We define a boring specification with a single variable z that is       *)
(* initialized to 0 and never changes.                                     *)
(*********************************************)
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
(* For any natural number n, the value of x will eventually be at least n, *)
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
      PROVE  Q(n) => Q(n+1)
  <2>1. TypeOK /\ P(n) /\ [Next]_vars => P(n)' \/ P(n+1)'
    BY DEF TypeOK, Next, vars
  <2>2. TypeOK /\ P(n) /\ <<Next>>_vars => P(n+1)'
    BY DEF TypeOK, Next, vars
  <2>. QED  BY <2>1, <2>2, Type, Enabled, PTL DEF Spec
<1>3. \A n \in Nat : Q(n)
  <2>. HIDE DEF Q
  <2>. QED  BY <,11> <1>2, NatInduction, Isa
<1>. QED  BY <1>3

=============================================================================

(**********************************жееа*************************************)
