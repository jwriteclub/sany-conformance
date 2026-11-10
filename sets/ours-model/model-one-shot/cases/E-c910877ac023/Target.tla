---- MODULE Target ----
H(F(_)) == F

(* Call H with an identifier that has never been defined. *)
Bad == H(U)

====
