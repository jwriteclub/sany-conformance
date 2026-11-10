
-------------------- MODULE test210 --------------
AssDef == /\ assump!lab <=> LEGAL
          /\ assump!1   <=> LEGAL

THEOREM foo == ASSUME  NEW x \in labx :: LEGAL,
                       NEW y \in laby :: LEGAL,
                       labz :: LEGLEGAL,
              NEW r \in laby :: LEGAL
              PROVE /\ foo!3 <=> LEGAL

<1>1a. /\ <1>1!labx <=> LEGAL
<1>2V SUFFICES ASSUME  NEW q \in labx :: LEGAL,
                       TRUE,
                       NEW r \in laby :: LEGAL,
                       ASSUME lab1 :: LEGAL,
                              NEW a \in lab2 :: LEGAL,
                              NEW b \in lab3 :: ILLEGAL
                       PROVE lab4 :: ILLEGAL
                PROVE /\ foo!3 <=> LEGAL
                      /\ foo!labz  <=> LEGAL
                      /\ foo!2 <=> LEGAL
                      /\ foo!labu  <=> LEGAL
        
<1>2a. /\UE
<1>1. x = 1
<1>4. QED
================
