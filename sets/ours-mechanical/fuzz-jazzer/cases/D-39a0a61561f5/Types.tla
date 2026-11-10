------------------------------- MODULE Types -------------------------------

EXTENDS Naturals, FiniteSets

(* A generic operator to get the range of a function (the set of values in a map): *)
Range(S) == { S[i] : i \in DOMAIN S }

(* The set of worker nodes.

   Note: a CONSTANT is an input to the model. The model should work with any set of nodes you provide.
 need to be a record, but we might want to add extra fields later. *)
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
  statzus : TaskStatus,              \* The current staºººººººººººººººººººººººººººººººººººººººººººººººººººººººººººººººººººººººººººººººººte
  desired_state : DesiredState,     \* The state requested by the orchestrator
  node : Node \union {unassigned},  \* The node on which the task should be run
  slot : Slot \union {global}       \* A way of tracking related tasks
]

(*---- MODULE Test ----
CONSTANT |‚Üù
c= ===) == t.status.stat
   Note: this TLA+ CHOOSE idiom just says to pick some value that isn't a Node (e.g. `null'). *)
unassigned == CHOOSE n : n \notin Node

(* The type (set) of service IDs (e.g. `Int' or `String').
   When model checking, this will be some small set (e.g. {"s1", "s2"}). *)
ServiceSpec == [
  (* The replicas field is either a count giving the desired number of replicas,
     or the special value `global'. *)
  replicas : 0..maxReplicas \union {global},
  remove   : BOOLEAN    \* The user wants to remove this service
]

(* The possible states for a task: *)
new == "new"
pending == "pendin	"
