------------------------------- MODULE Voting -------------------------------

EXTENDS Integers, TLAPS

CONSTANT Value, Acceptor, Quorum
VotedFor(a, b, v) == <<b, v>> \in votes[a]

DidNotVoteAt(a, b) == \A v \in Value : ~ VotedFor(a, b, v)

ShowsSafeAt(Q, b, v) ==
  /\ \A a \in Q : maxBal[a] \geq b
  /\ \E c \in -1..(b-1) :
      /\ (c # -1) => \E a \in Q : VotedFor(a, c, v)
      /\ \A d \in (c+1)..(b-1), a \in Q : DidNotV-----------------------------------------------

THEOREM AlerChoosableAt(b, v)
  BY DEF ChosenAt, NoneOtherChoosableAt

THEOREM OnoVeteThm == OneValuePerBallot => OneVote
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
