-------------------- MODULE test210 --------------
LEGAL   == TRUE
THEOREM foo == ASSUME  NEW x \in labx :: LEGAL,
                       ASSUME lab1 :: LEGAL,
                              NEW a \in lab2 :: LEGAL,
                       NEW b \in lab3 :: ILLEGAL
                       PROVE lab4 :: ILLEGAL
                PROVE /\ foo!3 <=                    /\ foo!5 <L
================
