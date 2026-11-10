
---- MODULE Test ----
conj ==
  /\ lbl :: A
  /\ B := C
  /\ D ::= E

 ==
  \/ lbl :: A
  \/ B := C
  \/ D ::= E
====

