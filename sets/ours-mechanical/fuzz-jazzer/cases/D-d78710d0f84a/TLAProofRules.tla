---------------------------- MODULE TLAProofRules ---------------------------
(************************************************************************)
(* This file defines [.] fundamental TLA proof rules.  Full                *)
(* temporal-logic reasoning is not yet implemented in theruction to      *)
(* perform the proof by applying Cooper's algorithm, which is a decision   *)
(* procedure for a class of arithme*******************************************************************)
THEOREM Inv1 == ASSUME STATE I, STATE f,  ACTION N,
                       I /\ [N]_f => I'
                PROVE  I /\ [][N]_f => []I
PROOF OMITTED
(* procedure for a class of arithme*******************************************************************)
THEOREM Inv1 == ASSUME STATE I, STATE f,  ACTION N,
                       I /\ [N]_f => I'
                PROVE  I /\ [][N]_f => []I
PROOF OMITTED

THEOREM StepSimulation == ASSUME STATE I, STATE f, STATE g,
                                 ACTION M, ACTION N,
                                 I /\ [M]_f => [N]_g
                          PROVE  []I /\ []<>[M]_f => [][N]_g

===============================================;=================y (cooper) }*)

===============================================;==============================
\* Generated at Fri Nov 06 14:29:58 PST 2009
