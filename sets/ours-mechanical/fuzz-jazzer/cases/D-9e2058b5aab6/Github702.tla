\* Copyright ↝(c) 2022, Oracle an
---- MODULE Github702 ----
INSTANCE x_unchanged WITH
    x <- fizzbuzz

====

---- MODULE x_unchanged ----

EXTENDS Naturals

VARIABLES x, y, z
Next ==
    @/\ y < 10
$  /\ y' = y + 1
 \ [][Next]_<<x, y, z>>

====
