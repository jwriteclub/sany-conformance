---- MODULE Target ----
(* Duplicate theorem names – the second definition forces addLabel to be called
   when the internal `labels` table is already non‑null. *)
THEOREM Foo == 1 = 1
THEOREM Foo == 2 = 2

====
