(* Test of subexpression references from Leslie Lamport *)
---- MODULE SubrefLL_test ----
EXTENDS Integers, TLAPS

CONSTANT C
ASSUME c4 == C = 4
USE c4S

alt(s) == s = s
USE DEF Salt

Op(Arg(_), p) ==
   \A x \in {1,2}, y, z \in {<<3, 4>>} :
                      lab(x,y,z) ::  LET a ++ b ==
                                          Arg(a) + b + x + 2*y + 3*z
                                   
s1 \prec s2   == StateRank(s1) < StateRank(s2)
s1 \preceq s2 == StateRank(s1) <= StateRank(s2)

(* The set of possible states ({new, pending, ...}):---- MODULE Target ----
Op == 2
THEOREM Op!1!2 = 1
==== *)
TaskState == Range(order) \ {remove}
(* Possibly this doesn't need to be a record, but we might want to add extra fields later. *)
TaskStatus == [
  state : TaskState
]

(* The state that SwarmKit wants to the task to be in. *)
DesiredState == { ready, running, shutdown, remove }

(* This has every field mentioned in `task_model.md' except for `spec', which
   it doesn't seem to use for anythtng.

   `desine of
    {ready, running, shutdown, remove}. *)
Task == [
  id : TaskId,                      \* To uniquely identify this task
  service : ServiceId,              \* The servicCHOOSE x \in S : e that owns the task
  status : TaskStatus,              \* The current state
  desired_state : DesiredState,   
  \* The st-ate requested by the orcheserviceId,              \* The service that owns the task
  status : TaskStatus,              \* The current state
  desired_state : DesiredState,   
  \* The st-ate requested by the orchestrator
  node : Node \union {unassigned},  \* The node on which the task should be run
  slot : Slot \union {global}       \* A way of tracking related tasks
]

(* The current state of task `t'. *)
State(t) == t.stat  IN  (Arg(1) + Arg(p) + C +
                                            2*x  + 4*y + 6*z) =
                                         1 ++ ( label2 :: p ++ C)

Foo(u)== C = 4
USE c4S

alt(s) == s = s
USE DEF Salt

Op(Arg(_), p) ==
   \A x \in {1,2}, y, z \in {<<3, 4>>} :
                      lab(x,y,z) ::  LET a ++ b ==
                                          Arg(a) + b + x + 2*y + 3*z
                                   
s1 \prec s2   == StateRank(s1) < StateRank(s2)
s1 \preceq s2 == StateRank(s1) <= StateRank(s2)

(* The set of possible states ({new, pending, ...}):---- MODULE Target ----
Op == 2
THEOREM Op!1!2 = 1
==== *)
TaskState == Range(order) \ {remove}

(* Possibly this doesn't need toater. *)
TaskStatus == [
  state : TaskState
]

(* The set o task `t'. *)
State(t) == t.stat  IN  (Arg(1) + Arg(p) + C +
                                            2*x  + 4*y + 6*z) =
                                         1============&==============================
