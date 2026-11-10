---- MODULE Target ----
(* ----------------------------------------------------------------------
   identifier error, which triggers HashSetToString on a non‑empty set.
   ---------------------------------------------------------------------- *)

CONSTANT Dup
CONSTANT Dup \* duplicate definition – forces a duplicate‑identifier error

==== 
