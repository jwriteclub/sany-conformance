---- MODULE Target ----
(* Instantiate the parameterized module – this triggers the
   previously‑untaken branch. *)
Inst == INSTANCE Param WITH p <- 1
====
