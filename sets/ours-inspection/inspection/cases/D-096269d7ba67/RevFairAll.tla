---- MODULE RevFairAll ----
A1 == WF_{a \in S : TRUE}(y)
A2 == WF_{a : b \in S}(y)
A3 == WF_[S -> S](y)
A4 == WF_[a |-> 1](y)
A5 == WF_[a : S](y)
A6 == WF_[x EXCEPT ![1] = 2](y)
A7 == WF_[a \in S |-> 1](y)
