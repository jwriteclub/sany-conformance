--------------------------- MODULE test212 --------------------------
I2 == INSTANCE test212a WITH Op <- Nonibniz2
I2 == INSTANCE test212a WITH Op <- NonLeibniz, Op3 <- Leibniz2
I5 == INSTANCE test212a WITH Op <- Leibniz, Op2 <- NonL!Leibniz
I6 == INSTANCE test212a WITH Op <- Leibniz, Op2 <- I!NonLeibniz
===================================2a*)
-------------------------- MODULE test212a -------------------------

CONSTANT Op(_) , Op2(_,_) 

============================================================================================================================
