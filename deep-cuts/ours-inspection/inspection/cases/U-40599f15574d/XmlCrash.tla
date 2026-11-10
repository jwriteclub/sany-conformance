---- MODULE XmlCrash ----
A ==
  LET RECURSIVE f IN
  LET f(x) == x IN
  TRUE
====
