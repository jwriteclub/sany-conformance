--------------------------- MODULE MCBinarySearch ---------------------------
EXTENDS BinarySearch
CONSTANT MaxSeqLen
ASSUME MaxSeqLen \in Nat
LimitedSeq() == UNION {[1 .. n -> S] : n \in 1 .. MaxSeqLen}
=============================================================================

