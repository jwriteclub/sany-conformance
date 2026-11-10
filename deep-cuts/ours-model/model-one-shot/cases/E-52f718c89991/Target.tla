---- MODULE Target ----
HO(op(_)) == op(0)

\* Use the higher‑order operator, selecting its operator argument.
THEOREM HO!op
PROOF BY DEF HO
====
