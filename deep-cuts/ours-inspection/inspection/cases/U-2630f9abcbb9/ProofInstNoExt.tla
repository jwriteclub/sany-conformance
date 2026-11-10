---- MODULE ProofInstNoExt ----
CONSTANT C
THEOREM TRUE
<1> INSTANCE PinM
<1> INSTANCE PinF
<1> QED
====
(*@FILE PinM*)
---- MODULE PinM ----
CONSTANT C
D == 1
THEOREM TI == TRUE
====
(*@FILE PinF*)
---- MODULE PinF ----
E == 1
THEOREM TJ == TRUE
====
