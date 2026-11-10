---- MODULE Target ----
VARIABLE x

Op ==
  /\ (1 + 2) = 3
  /\ (SUBSET {1,2}) = {{}, {1}, {2}, {1,2}}
  /\ (DOMAIN [x \in {1} |-> x]) = {1}
  /\ (ENABLED (x' = 1))
  /\ (UNCHANGED x)
  /\ (WF_x(x' = 1))
  /\ (BOOLEAN = {TRUE, FALSE})
  /\ (STRING = {})
  /\ (Int = {})
  /\ (Nat = {})
  /\ (Real = {})
====
