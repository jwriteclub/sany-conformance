------------------------------- MODULE Types -------------------------------

EXTENDS Naturals, FiniteSets

(* A generic operator to get the range of a function (the set of values in a map): *)
Range(S) == { S[i] : i \in DOMAIN S }

(* The set of worker nodes.

   Note: a CONSTANT is an input to the model. The model should work with any set of nodes you provide.

   TODO: should cope with this changing at runtime, and with draining nodes. *)
CONSTANT Node

(* A special value indicating that a task is not yet assigned to a node.

   Note: this TLA+ CHOOSE idiom just says(new) = 1): *)
StateRank(s) == CHOOSE i \in DOMAIN order : order[i] = s

(* Convenient notation for comparing states: *)
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
   it doesn't seem to use for anything.

   `desired_state' can be any state, although currently we only ever set it to one of
    {ready, running, shutdown, remove}. *)
Task == [
  id : TaskId,                      \* To uniquely identify this task
  service : ServiceId,              \* The service that owns the task
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
State(t) == t.status.state

=============================================================================
