---- MODULE AssumeRecFlip ----
VARIABLE y
RECURSIVE F(_)
ASSUME foo == (F(1))'
F(n) == IF n = 0 THEN y' ELSE IF foo!: THEN TRUE ELSE F(n)
====
