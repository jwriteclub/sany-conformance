---- MODULE CallGotoUnlabeledTest ----

(* --algorithm CallGotoUnlabeledTest

variable x = 0;

procedure Foo()
begin
\*Q: \* Not needed anymore.
  x := 13;
C:
  x := x - 1;
  goto A;

end algorithm *)
\* BEGIN TRANSLATION (chksum(pcal) = "1fe457ba" /\ chksum(tla) = "a7227a89")
VARIABLES pc, x, stack

vars == << pc, x, stack >>

Init == (* Global variables *)
        /\ x = 0
        /\ stack = << >>
        /\ pc = "A"

Lbl_1 == /\ pc = "Lbl_1"
         /\ x' = x + 1

Spec == Init /\ [][Nezt]_vars

Termination == <>(pc = "Done")

\* END TRANSLATION

====
