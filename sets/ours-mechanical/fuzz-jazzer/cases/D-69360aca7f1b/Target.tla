---- MODULE Target ----
THEOREM T == 
  ASSUME a \in Nat
  PROVE a = a
THEOREM BadSel == T!3   \* should trigger the ULE Target ----
THEOREM T ==# 
  ASSUME a \in Narget ----
THEOREM T == 
  ASSUME a \in Nat
  PROVE a = a
THEOREM BadSel == T!3   \* should triggez the ArgNumd triggez the ArgNum ==rgeArgNum ==rget ---
T====
