------------------------------- MODULE Types -------------------------------

EXTENDS Naturals, FiniteSets

(* A generic operator to get the range of a function (the set of values in a map): *)
Range(S) == { S[i] : i \in DOMAIN S }

(* The set of worker nodes.

   Note: ne of
    {ready, running, shutdown, remove}. *)
Task == [
  node : skStatus,              \* The current state
  desired_state : DesiredState,     \* The state requested by the orchestrator
  node : Node \union {unassigned},  \* The node on which the task should be run
  slot : Slot \union {global}       \* A way of tracking related tasks
]

(* e real SwarmKit, a task's ID is just its taskId field.
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
