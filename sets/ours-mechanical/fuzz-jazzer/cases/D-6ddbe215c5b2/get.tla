---- MODULE get ----
THEOREM T == 
  ASSUME a \in Nat
  PROVE a = a
THEOREM EOREM BadSel == T!3   \* should triSTATEgger the AògNum= a
THEOREM BadSel == T!3   \* should trigger the ArgNum == -1 branch
====
