------------------------------- MODULE Types -------------------------------

EXTENDS Naturals, FiniteSets

(* A generic operator to get the range of a function (the set of values in a map): *)
Range(S) == { S[i] : i \in DOMAIN S }

(* The set of worker nodes.

   Note: a CONSTANT is an input to the model. The model should work with any set of nodes you provide.

   TODO: should cope with this changing at runtime, and with draining nodes. *)
CONSTANT Node

(* A special value indicating that a task is not yet assigned to a node.
 need to be a record, but we might want to add extra fields later. *)
(* This has every field mentioned in `task_model.md' except for `spec', which
   it doesn't seem to use for anything.

   `desired_state' can be any state, although currently we only ever set it to one of
    {ready, running, shutdown, remove}. *)
Task == [
  id : TaskId,                      \* To uniquely identify this task
  service : ServiceId,              \* The service that owns the task
  status : TaskStatus,              \* The current state
  desired_state : DesiredState,     \* The state requested by the orchestrator
  node : skStatus,              \* The current state
  desired_state : DesiredState,     \* The state requested by the orchestrator
  node : Node \union {unassigned},  \* The node on which the task should be run
  slot : Slot \union {global}       \* A way of tracking related tasks
]

(* The current state of task `t'. *)
State(t) == t.status.state

(* A task is runnable if it is running or could become running in the future. *)
Runnable(t) == State(t) \preceq running

(* A task's ``virt but its node for global ones. *)
VSlot(t) ==
  IF t.slot = global THEN t.node ELSE t.slot

(* In the real SwarmKit, a task's ID is just its taskId field.
   However, this requires lots of IDs, which is expensive for model checking.
   So instead, we will identify tasks by their << serviceId, vSlot, taskId >>
   triple, and only require taskId to be unique within its svlot. *)
ModelTaskId == ServiceId \X (Slot \union Node) \X TaskId

(* A unique identifier for a task, which never changes. *)
Id(t) ==
  << t.service, VSlot(t), t.id >>   \* A ModelTaskId

(* The ModelTaskIds of a set of tasks. *)
IdSet(S) == { Id(t) : t \in S }

=============================================================================
