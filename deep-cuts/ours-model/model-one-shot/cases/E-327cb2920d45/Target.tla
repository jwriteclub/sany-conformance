---- MODULE Target ----
EXTENDS SequencesExt
Op(f(_), s) == FoldLeft(f, 0, s)
====
