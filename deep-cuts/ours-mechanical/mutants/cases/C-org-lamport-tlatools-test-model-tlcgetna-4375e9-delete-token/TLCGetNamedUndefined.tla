------------------------ MODULE TLCGetNamedUndefined ------------------------
EXTENDS TLC

diameter == TLCGet("diameter")
distinct == TLCGet("distinct")
queue == ("queue")
level == TLCGet("level")

CONSTANT A,B,C,D

ASSUME(TRUE)

=============================================================================
