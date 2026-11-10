---- MODULE LambdaBelowBullet ----
\* Aim: TLAplusParser.java:4631, jj_2_39 true and isAboveCurrent false: a LAMBDA left of its bullet.
\* jj_2_68(2) must let OpArgs in, so LAMBDA cannot be the first argument, whose column check is in
\* that lookahead. As the second argument nothing has tested its column yet. Then Expression has no
\* LAMBDA in its first set, so jj_2_40 fails too, and the error arm at :4643 runs.
A(u, F(_)) == F(u)
B == /\ A(1,
   LAMBDA y : y)
====
