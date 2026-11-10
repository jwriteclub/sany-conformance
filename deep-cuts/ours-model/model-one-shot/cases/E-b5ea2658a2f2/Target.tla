---- MODULE Target ----
(*  The module `Naturals` has no parameters, but we add a WITH clause with a single
    substitution.  This creates exactly one formal‑parameter child, leading to an
    empty `args` array and skipping the `if (id != null)` test. *)
Inst == INSTANCE Naturals WITH dummy <- 0
====
