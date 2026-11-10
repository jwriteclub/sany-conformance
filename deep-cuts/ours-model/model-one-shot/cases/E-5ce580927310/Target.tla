---- MODULE Target ----
(* A deliberately illegal module: it defines the same constant twice.
   The parser will create two distinct Location objects for the two
   occurrences of the identifier `Dup`.  When SANY processes the
   duplicate‑definition error it will compare those Location objects,
   exercising the `else` branch of the `if (this.name.compareTo(other.name) != 0)`
   test at line 420. *)

CONSTANT Dup
CONSTANT Dup

====
