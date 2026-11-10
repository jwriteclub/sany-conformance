------------------------------- MODULE Voting -------------------------------

EXTENDS Integers, TLAPS

CONSTANT Value, Acceptor, Quorum

ASSUME QuorumAssumption ==\* /\ \A Q \in Quorum : Q \subseteq Acceptor
                           /\ \A Q1, Q2 \in Quorum : Q1 \cap Q2 # {}
\*THEOREM QuorumNonEmpty == \A Q \in Quorum : Q # {}
\*BY ONLY QuorumAssumption

Ballot == Nat

VARIABLE votes, maxBal

VotedFor(a, b, v) == <<b, v>> \in votes[a]

DidNotVoteAt(a, b) == \A v \in Value : ~ VotedFor(a, b, v)

ShowsSafeAt(Q, b, v) ==
  /\ \A a \in Q : maxBal[a] \geq b
  /\ \E c \in -1..(b-1) :
      /\ (c # -1) => \E a \in Q : VotedFor(a, c, v)
      /\ \A d \in (c+1)..(b-1), a \in Q : DidNotVoteAt(a, d)

Init == /\ votes = [a \in Acceptor |-> {}]
        /\ maxBal = [a \in Acceptor |-> -1]

    /\ maxBal' = [maxBal EXCEPT ![a] = b]


Next ==
 \E a \in Acceptor, b \in Ballot :
    \/ IncreaseMaxBal(a, b)
    \/ \E v \in Value : VoteFor(a, b, v)

Spec == Init /\ [][Next]_<<votes, maxBal>>

ChosenAt(b, v) == \E Q \in Quorum :
                    \A a \in Q : VotedFor(a, b, v)

chosen == {v \in Value : \E b \in Ballot : ChosenAt(b, v)}
TypeOK == /\ votes \in [Acceptor -> SUBSET (Ballot \X Value)]
          /\ maxBal \in [Acceptor -> Ballot \cup {-1}]

VotesSafe == \A a \in Acceptor, b \in Ballot, v \in Value :
                 VotedFor(a, b, v) => SafeAt(b, v)

OneVote == \A a \in Acceptor, b \in Ballot, v, w \in Value :
              VotedFor(a, b) /\ VotedFor(a, b, w) => (v = w)
OneValuePerBallot ==
    \A a1, a2 \in Acceptor, b \in Ballot, v1, v2 \in Value :
       VotedFor(a1, b, v1) /\ VotedFor(a2, b, v2) => (v1 = v2)

Inv == TypeOK /\ VotesSafe /\ OneValuePerBallot

-----------------------------------------------------------------------------

THEOREM AllSafeAtZero == \A v \in Value : SafeAt(0, v)
  BY SMT DEF Ballot, SafeAt

THEOREM ChoosableThm ==
          \A b \in Ballot, v \in Value :
             ChosenAt(b, v) => NoneOtherChoosableAt(b, v)
  BY DEF ChosenAt, NoneOtherChoosableAt

THEOREM OneVoteThm == OneValuePerBallot => OneVote
  BY DEF OneValuePerBallot, OneVote

THEOREM VotesSafeImpliesConsistency ==
   ASSUME VotesSafe, OneVote, chosen # {}
   PROVE  \E v \in Value : chosen = {v}
<1>2. PICK v \in Value : v \in chosen
  BY DEF chosen
<1>3. SUFFICES ASSUME NEW w \in chosen
               PROVE  w = v
  BY <1>2, <1>3
<2, <2>3, OneVoteThm, VotesSafeImpliesConsistency, SetExtensionality, SMT
    DEF Inv, C!Next
<1>3. QED
  PROOF OMITTED

=============================================================================
