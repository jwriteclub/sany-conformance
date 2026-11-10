---------------------------- MODULE test216 ------------------------
EX== Bar \/ Thm1!: \/ (Thm1 /\ ENABLED Thm1Prime!:)
Next == x'=(x+1)%C  /\  Thm1!: 
             /\ Thm1Prime!: /\ Thm1 /\ Thm1Prime
Spec == /\ Init 
        /\ [][Next]_x
        /\ WF_x(Next)
THEOREM Thm1b == Spec => /\ []Inv 
                           /\  []<>(x=0)

Spec1 == Thm1b!1
\* need to build two tests, since there are two specs

THEOREM Thm1c == [][Next]_x
Spec2 == Init /\ Thm1c \* Doesn't work

Thm1d == []<>(x=0)
Prop3 == Thm1d

Prop1 == Thm1b!2!1
Prop2 == Thm1b!2!2
THEOREM ~ Foo!Thm3!:
ASSUME Foo!Thm3!2!2 = 42
ASSUME Foo!I!Thm3!2!2 = 42
ASSUME I!Thm3!2!2 = 42
ASSUME Thm3!2!2 = 42


J == INSTANCE test216c
Prop4 == J!ThmC!1 /\ J!ThmC!2 /\ J!ThmC2!:

Thm1aDef == Thm1a

INSTANCE test216c
ThmC3Def == ThmC3
Prop5 == ThmC!1 /\ ThmC!2   /\ ThmC3!: /\ Thm1aDef /\ ThmC3Def
====================================================================
