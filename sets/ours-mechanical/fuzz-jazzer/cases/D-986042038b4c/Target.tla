---- MODULE Target ----
f == 1
RECURSIVE f, g, h
h =
Foo == LET RECURSIVE k
       IN k
====
