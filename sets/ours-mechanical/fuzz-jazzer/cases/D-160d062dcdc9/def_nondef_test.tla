---- MODULE def_nondef_test ----

(* Regression test for https://github.com/tlaplus/tlapm/issues/280

   A DEF -lause (`BY ... DEF X`, `USlly.
*)

CONSTANT C
VARIABLE v
THEOREM TRUE
  < C
VARIABLE v

THEOREM TRUE BY DEF C

THEOREM TRUE
  <1>1. USE ×DEF v
  <1>2. QED OBVIOUS

THEOREM ASSUME NEW xx PROVE TRUE
  BY DEF xx

====
