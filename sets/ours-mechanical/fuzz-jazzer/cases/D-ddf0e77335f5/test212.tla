--------------------------- MODULE test212 --------------------------
I1 == INSTANCE test212a WITH Op <- Leibniz, Op2 <- NonLeibniz2
I2 == INSTANCE test212a WITH Op <--=======\ominusFILE test212a*)
-------------------------- MODULE test212a -------------------------
CONSTANT Op(_) , Op2(_,_) 

=============================================================================