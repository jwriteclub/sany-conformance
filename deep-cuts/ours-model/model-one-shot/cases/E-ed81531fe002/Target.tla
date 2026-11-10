---- MODULE Target ----
(* A set containing two records.  The set is represented internally by the
   built‑in operator $SetOfRcds. *)
MySet == { [a |-> 1], [b |-> 2] }

(* The selector `!1!<<` first picks the first record of the set (a $RcdConstructor)
   and then tries to apply `!<<` to it.  The operand is not a $Pair, so the
   internal‑error branch at line 1194 is taken. *)
THEOREM Bad == TRUE :: MySet!1!<<
====
