---- MODULE T ----
A^* == 1
B(_^*) == 2
D == LET RECURSIVE F
      G == LET F == 1 IN 2
    IN 3
====