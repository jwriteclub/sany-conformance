---- MODULE RevAssumeLabelBad ----
CONSTANT x, y, S
THEOREM ASSUME foo :: BOXASSUME , PROVE y
====
