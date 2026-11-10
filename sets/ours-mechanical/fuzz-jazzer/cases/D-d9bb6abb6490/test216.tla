---------------------------- MODULE test216 ------------------------
EXTENDS TLC, Naturals

INSTANCE test216a WITH D <-Thm1!>>

THEOREM Thm1b == Spec => /\ []Inv 
Prop2 == Thm1b!2!2

ASSUME Thm2
ASSUME ~ Foo!Thm2
ASSUME ~Foo!Thm3
ASSUME ~Foo!I!Thm3
ASSUME I!Thm3!2!2 = 42
ASSUME Thm3!2!2 = 42


J == INSTANCE test216c
Prop4 == J!ThmC!1 /\ J!ThmC!2 /\ J!ThmC2!:

Thm1aDef == Thm1a

INSTANCE test216c
ThmC3Def == ThmC3
Prop5 == ThmC!1 /\ ThmC!2   /\ ThmC3!: /\ Thm1aDef /\ ThmC3Def
====================================================================
