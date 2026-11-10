---- MODULE P11 ----
VARIABLE x
CONSTANT c
A == x!y
B == c!z
THEOREM T == TRUE
<1>1. TRUE
<1> QED
D == T!<1>1
E == T!<1>1!z
THEOREM T2 == TRUE
<1>2. TRUE
  <2>3. <1>2!x
  <2> QED
<1> QED
====
