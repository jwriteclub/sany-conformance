------------------------ MODULE FIFO -------------------------
CONSTANT Message
VARIABLES in, out
Inner(q) == INSTANCE InnerFIFO 
Spec == \EE q : Inner(q, 0)!Spec
==============================================================
