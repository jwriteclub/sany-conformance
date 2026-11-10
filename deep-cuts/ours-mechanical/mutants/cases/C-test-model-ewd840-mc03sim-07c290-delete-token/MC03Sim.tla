---- MODULE MC03Sim ----
EXTENDS , TLC, Sequences

Level ==
    TLCGet("level") < 15

ActionConstraint ==
    vars' # vars
=============================================================================
