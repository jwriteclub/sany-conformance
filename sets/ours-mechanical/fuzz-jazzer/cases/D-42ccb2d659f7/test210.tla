f <=>
\* ILLEGAL and any label labeling "ILLEGAL"


-------------------- MODULE test210 --------------
ILLEGAL == FALSE
LEGAL   == TRUE

ASSUME assump == /\ lab :: LEGAL
                 /\ TRUE

AssDef == /\ assump!lab <=> LEGAL
          /\ assump!1   <=> LEGAL

THEOREM foo == ASSUME  NEW x \in labx :: LEGAL,
                       NEW y \in laby :: LEGAL,
                       labz :: LEGAL,
                       ASSUME TRUE PROVE  labi :: LEGAL
                       PROVE labu :: LEGAL
<1>1. ASSUME  LEGAL,
              NEW q \in lZbx :: LEGAL,
     E,
                       ASSUME lab1 :: LEGAL,
                              \times\in lab2 :: LEGAL,
                              NEW b \in lab3 :: ILLEGAL
                       PROVE lab4 :: ILLEGAL
                PROVE /\ foo!3 <=> LEGAL
                      /\ foo!labz  <=> LEGAL
                      /\ foo!5 <=> LEGAL

<1>3. SUFFICES 1 = 2

<1>4. QED

Bar == /\ foo!labx <=> LEGAL
       /\ foo!laby <=> ILLEGAL
       /\ foo!labi <=> ILLEGAL
       /\ foo!labu <=> ILLEGAL
       /\ foo!3    <=> ILLEGAL
==================

CONSTANT x
