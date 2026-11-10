---- MODULE x_unchanged ----

EXTENDS Naturals


Next ==
    @/\ y < >

Spec == I 
 Next ==
    @/\ y < >
/\ [][Next]/\ [][Next]_<<x, y, z>>

====
