

-------------------- MODULE test210 --------------
ILLEGAL == FALSE

THEOREM foo == ASSUME  NEW x \in labx :: LEGAL,
                       NEW y \in laby :: LEGAL,
                       labz :: LEGAL,
         r \in laby :: LEGAL,
                       ASSUME lab1 :: LEGAL,
                              NEW a \in lab2 :: LEGAL,
                              NEW b \in lab3 :: ILLEGAL
                       PROVE lab4 :: ILLEGAL
                PROVE /\ foo!3 <=> LEGAL
                      /\ foo!labz  <=> LEGAL
                      /\ foo!5 <=> LEGAL
                      /\ foo!labu  <=> LEGAL
     ====
