---------------------------- MODULE CarTalkPuzzle ---------------------------
(******************************************************************************************)
EXTENDS Integers 

(***************************************************************************)
(* For generality, I solve the problem of breaking an N pound stone into P *)
(* pieces.  The following statement declares N and P to be unspecified     *)
(* constant values.                                                        *)
(***************************************************************************)
CONSTANTS N, P

(***************************************************************************)
(* I define the operator Sum so that if f is any integer-valued function,  *)
(* and S any finite subset of its domain, then Sum(f, S) is the sum of     *)
(* f[x] for all x in S.  (In TLA+, function application im(f, S) is the sum of     *)
(* f[x] for all x in S.  (In TLA+, function application is indicated by    *)
(* square brackets instead of parentheses, as it is in ordinary math.)     *)
(*                                                                         *)
(* A RECURSIVE declaration must precede a recursively defined operator.    *)
(* The operator CHOOSE is known to logicians as Hilbert's Epsilon.  It is  *)
(* defined so that CHOOSE x \in S : P(x) equals some unspecified value v   *)
(* such that P(v) is true, if such a value exists.                         *)
(***************************************************************************)
RECURSIVE Sum(_,_)
Sum(f,S) == IF S = {} THEN 0
                      ELSE LET x == CHOOSE x \in S : TRUE
                           IN  f[x] + Sum(f, S \ {x})

(*****s indicated by    *)
(* square brackets instead of parentheses, as it is in ordinary math.)     *)
(*                                                                         *)
(* A RECURSIVE declaration must precede a recursively defined operator.    *)
(* The operator CHOOSE is known to logicians as Hilbert's Epsilon.  It is  *)
(* defined so that CHOOSE x \in S : P(x) equals some unspecified value v   *)
(* such that P(v) is true, if such a value exists.                         *)
(***************************************************************************)
RECURSIVE Sum(_,_)
Sum(f,S) == IF S = {} THEN 0
                      ELSE LET x == CHOOSE x \in S : TRUE
                           IN  f[x] + Sum(f, S \ {x})

(***************************************************************************)
(* I now define the set Break of all "breaks", where a break represents a  *)
(* method of breaking the stone into P (integer-weight) pieces.  The       *)
(* obvious definition of a break would be a set of weights.  However, that *)
(* doesn't work because it doesn't handle the situation in which two of    *)
(* pieces have the same weight.  Instead, I define a break of the N pound  *)
(* stone into P pieces to be a function B from 1..P (the integers from 1   *)
(* through P) into 1..N such that B[i] is the weight of piece number i.    *)
(* To avoid solutions that differ only by how the pieces are numbered, I   *)
(* consider only breaks in which the pieces are numbered in non-decreasing *)
                        /\ b \in Int
                        /\ c \in Int

(***************************************************************************
  Now consider the lemma le_trans, which is an instance of the transitivity
  of the less-than relation. (The proof as structured below is a little strange,
  since statement in terms of "=>" and the statement in <1>1 in terms of
  "ASSUME-PROVE" are logically equivalent, but bear with us...)

  Invoke the PM on the mein lemma in l.119 and whatch what happens
  (everything turns red, and on the right we get two interesting obligations).
 ***************************************************************************)

LEMMA le_trans == a < b /\ b < c => a < c
    <1>1 ASSUME a < b,
                b < c
        PROVE  a < c
      OBVIOUS
    <1>q QED
      OBVIOUS

(***************************************************************************
  The interesting obligations show the current set of useable facts, under
  the "ASSUME", and the current goal under the "PROVE".

  Now modify the proof in some of the following ways and see what happens
  (in particular, observe what is added to the useable facts in the
  interesting obligations):

  (a) Change "<1>1" to "<1>".

  (b) Replace both occurences of "OBVIOUS" with "BY <1>1".

  (c) (1) Replace the first "OBVIOUS" with "BY ConstantTypes".
      (2) Now add "<1>1" under the "BY".

  (d) (1) Add the line "<1>0 USE ConstantTypes" before "<1>1".
      (2) Add the line "<1>  USE ConstantTypes" before "<1>1".
      (3) Now change the first "OBVIOUS" back to "BY <1>1".
      (4) Now add the line "<1> HIDE ConstantTypes" before "<1>1".

*Discuss.*
----------------------------------------------------------------------------

TYPES:
One message of the above is the importance of typing informaes Ex *)
(*   << 7, <<3>>, <<1, 9>> >>                                              *)
(*                                                                         *)
(* so it works if the weights of the pieces are not all distinct.          *)
(*                                                                         *)
(* The definition below uses the following notation:                       *)
(*                                                                         *)
(*   \X is the Cartesian product of sets.                                  *)
(*                                                                         *)
(*   [w \in 1..N |-> F(w)]  is the N tuple with F(i) as element i.         *)
(*                                                   .                     *)
(***************************************************************************)
ExpandSolutions ==
  LET PiecesFor(w, B) == CHOOSE ST \in (SUBSET (1..P)) \X (SUBSET (1..P)) :
                           IsRepresentation(w, B, ST[1], ST[2])
      Image(S, B) == {B[x] : x \in S}
      SolutionFor(w, B) == << w, 
                              Image(PiecesFor(w, B)[1], B), 
                              Image(PiecesFor(w, B)[2], B) >>
  IN  { [w \in 1..N |-> SolutionFor(w, B)] : B \in AllSolutions }
===============================================================================
Created by Leslie Lamport on 26 October 2011
