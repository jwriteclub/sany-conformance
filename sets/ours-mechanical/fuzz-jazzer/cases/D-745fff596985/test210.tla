-------------------- MODULE test210 --------------
ILLEGAL == FALSE
LEGAL   == TRUE
THEOREM foo == ASSUME  NEW x \in labx :: LEGAL,
                       NEW y \in ly :: LEGAL,
                       ASSUME lab1 :: LEGAL,
                              NEW a \in lab2 :: LEGAL,
                              NEW b \in lab3 :: ILLEGAL
                       PROVE lab4 :: ILLEGAL
                PROVE /\ foo!3 <=> LEGAL
                      /\ foo!labz  <=> LEGAL
                      /\ foo!5 <L
================
