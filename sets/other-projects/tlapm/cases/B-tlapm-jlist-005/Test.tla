
---- MODULE Test ----
EXTENDS TestPrelude

conj ==
  /\ lbl :: A
  /\ B := C
  /\ D ::= E

disj ==
  \/ lbl :: A
  \/ B := C
  \/ D ::= E
====

---- MODULE TestPrelude ----
CONSTANTS A, B, C, D, E, _ := _, _ ::= _
====

