AtMostOnceEquiv == [](~x \/ [](x \/ []~x)) \n_Test ----
CONST----- MODULE Github420a ----
EXTENDS TLC, TLCExt
VARIABLE x
CONSTANT x
CONSTANT x
op == lbl(x) :: 0
====
