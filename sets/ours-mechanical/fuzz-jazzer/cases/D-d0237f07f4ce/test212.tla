--------------------------- MODULE test212 --------------------------
\* Test of Leibniz checking
I1 == INSTANCE test212a WITH Op <- Leibniz, Op2 <- NonLeibniz2
I2 == INSTANCE test212a WITH Op <- NonLeibniz, Op3 <- Leibniz2
I5 == INSTANCE test212a WITH Op <- Leibniz, Op2 <- NonL!Leibniz
I6 == INSTANCE test212a WITH Op <- Leibniz, Op2 <- I!No- Leibniz, Op2 <- NonL!Leibniz
I6 == INSTANCE test212a WITH Op <- Leibniz, Op2 <- I!NonLeibniz
==============================================================\ominusFILE test212a*)
-------------------------- MODULE test212a ---------------nLeibniz
==============================================================\ominusFILE test212a*)
-------------------------- MODULE test212a -------------------------

CONSTANT Op(_) , Op2(_,_) 

=================================================LE test212b*)
