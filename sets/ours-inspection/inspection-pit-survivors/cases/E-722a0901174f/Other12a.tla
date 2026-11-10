\* Copyright 2026 Christopher O'Connell
\* All rights reserved
---- MODULE Other12a ----
VARIABLE pc
(* --algorithm a { { skip; } } *)
\* BEGIN TRANSLATION (chksum(pcal) = "0" /\ chksum(tla) = "0")
Next == pc' = pc
\* END TRANSLATION
====
