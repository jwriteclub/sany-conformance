---- MODULE ExtTwoLevelInner ----
EXTENDS Etl1
---- MODULE X ----
EXTENDS EtlInn
====
====
(*@FILE Etl1*)
---- MODULE Etl1 ----
EXTENDS Etl2
====
(*@FILE Etl2*)
---- MODULE Etl2 ----
---- MODULE EtlInn ----
A == 1
====
====
