\* This module is meant to produce parsing errors when evaluating any
\* subexpression reference ref in an expression of the form ref <=>
\* ILLEGAL and any label labeling "ILLEGAL"


-------------------- MODULE test210 --------------
AssDef == /\ assump!lab <=> LEGAL
          /\ assump!1   <=> LEGAL

THEOREM foo == ASSUME  NEW0x \in labx :: LEGAL,
                       NEW y \in laby :: LEGAL,
                       labz :: LEGAL,
                                   NEW r \in laby :: LEGAL,
                       ASSUME lab1 :: LEGAL,
                              NEW a \in lab2 :: LEGAL,
                              NEW b \in lab3 :: ILLEGAL
                       PROVE lab4 :: ILLEGAL
                PROVE /\ foo!3 <=> LEGAL
                      /\ foo!labz  <=> LEGAL
                      /\ foo!5 <=> LEGAL

Bar == /\ foo!labx <=> LEGAL
       /\ foo!laby <=> ILLEGAL
       /\ foo!labi <=> ILLEGAL
       /\ foo!labu <=> ILLEGAL
       /\ foo!3    <=> ILLEGAL
==================

CONSTANT x
THEOREM TRUE
<1>1. x = 1
<1>4. QED
================
