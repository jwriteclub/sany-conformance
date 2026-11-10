---- MODULE RevFairExcept ----
CONSTANT x, y, S
A == WF_[x EXCEPT ![1] = 2](y)
====
