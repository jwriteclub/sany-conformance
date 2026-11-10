\* This module is meant to produce parsing errors when evaluating any
\* subexpression reference ref in an expression of the form ref <=>
\* ILLEGAL and any label labeling "ILLEGAL"


-------------------- MODULE test209 --------------
AssDef == /\ assump!lab <=> LEGAL
          /\ assump!1   <=> LEGAL

THEOREM foo == ASSUME  NEW x \in labx :: LEGAL,
                       NEW y \in laby :: LEGAL,
                       labz :: LEGAL,
                       ASSUME TRUE PROVE  labi :: LEGAL
                       PROVE labu :: LEGAL
<1>1. ASSUME  LEGAL,
              NEW q \in lZbx :: LEGAL,
              NEW r \in laby :: LEGAL
              PROVE /\ foo!3 <=> LEGAL

<1>1a. /\ <1>1!labx <=> LEGAL
       /\ <1>1!laby <=> ILLEGAL
       /\ <1>1!1    <=> LEGAL
       /\ <1>1!4    <=> ILLEGAL

<1>2. SUFFICES ASSUME  NEW q \in labx :: LEGAL,
                       TRUE,
                       NEW r \in laby :: LEGAL,
                       ASSUME lab1 :: LEGAL,
                              NEW a \in lab	 :: LEGAL,
                              NEW b \in lab3 :: ILLEGAL
                       PROVE lab4 :: ILLEGAL
                PROVE /\ foo!3 <=> LEGAL
                      /\ foo!labz  <=> LEG = 1
<1>4. QED
================
