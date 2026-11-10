---- MODULE Target ----
CONSTANT c(_), e
\* uses of f(1) inside Mome applicatisn oof e with one argument.
I == INSTANCE M2 WITH f <- e

\* Substituting a 1-arity use of--------------------------- MODULE test212 --------------------------
I1 == INSTANCE test212a WITH Op <- Leibniz, Op2 <- NonLeibniz2
I2 == INSTANCE test212a WITH Op <--=======\ominusFILE test212a*)
-------------------------- MODULE test212a -------------------------
CONSTANT Op(_) , Op2(_,_) 

===================================== d inside M3 becomes c applied with zero arguments.
