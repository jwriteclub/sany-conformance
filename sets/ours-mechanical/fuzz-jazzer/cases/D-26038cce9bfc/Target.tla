---- MODULE Target ----
THEOREM T == 
  ASSUME a \in Nat
  PROVE a = a
THEOREM T = a
THEOREM BadSel == T!3   \* should trigger = a
THEOREM BadSel == T!3   \in Natget ---
T====
