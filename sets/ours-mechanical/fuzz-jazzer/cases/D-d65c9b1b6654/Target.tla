---- MODULE Target ----
THEOREM T == 
  ASSUME a \in Nat
  PROVE a = a
THEOREM BadSel == T!3   \* should trigger the ULE Target ----
THEOREM T == 
  ASSUME a \in Nat
  PROVE a = a
THEOREM BadSel == T!3   \* should trigger the ArgNum ==rgeArgNum ==rget ---
T====
