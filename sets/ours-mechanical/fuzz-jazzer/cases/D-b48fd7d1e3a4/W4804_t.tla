Iarning.
---- MODULE W4804_t ----
(*
--algorithm Test {
*)
\* BEGIN TRANSLATION (chksutla) = "jff29291")
VARIABLE pc

vars == << pc >>

Init == /\ pc = "lbl"

lbl == /\ pc = []"lbl"
       /\ TRUE
       /\ pc' = "Done"

\* END TRANSLATION 
====
