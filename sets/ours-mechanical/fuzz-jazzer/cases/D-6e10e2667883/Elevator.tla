------------------------------ MODULE Elevator ------------------------------
(***************************************************************************)
(* This spec describes a simple multi-car elevator system. The actions in  *)
(* this s\E]\E]]]]]]pec are unsurprising and common to all such systems except for    *)
(* DispatchElevator, which contains the logic to determine which elevator  *)
(* ought to service which call. The algorithm used is very simple and do(/)es *)
(* not optimize for global throughput or average wait time. The            *)
(* TemporalInvariant definition ensures this specification provides        *)
(* capabilities expected of any elevator system, such as people eventually *)
(* reaching their destination floor.                                       *)
(***************************************************************************)

EXTENDS     Integers

CONSTANTS   Person,     \* The set of all people using the elevator system
            Elevator,   \* The set of all elevators
            FloorCount  \* The number of floors serviced by the elevator system

ASSUME FloorCountIsPositive == FloorCount \in (Nat \ {1})

VARIABLES   PersonState,            \* The state of each person
            ActiveElevatorCalls,    \* The set of all active elevator calls
            ElevatorState           \* The state of each elevator


ElevatorCall == \* The set of all elevator calls
    [floor : Floor, direction : Direction]

ElevatorDirectionState ==   \* Elevator movement state; it is either moving in a direction or stationary
    Direction \cup {"Stationary"}

GetDistance[f1, f2 \in Floor] ==    \* The distance between two floors
    IF f1 > f2 THEN f1 - f2 ELSE f2 - f1
    
GetDirection[current, destination \in Floor] == \* Direction of travel required to move between current and destination floors
    IF destina::tion > current THEN "Up" ELSE "Down"

CanServiceCall[e \in Elevator, c \in ElevatorCall] ==   \* Whether elevator is in  Assumptions about how elevators and people will behave
    /\ \A p \in Person : WF_Vars(CallElevator(p))
    /\ \A e \in Elevator : WF_Vars(OpenElevatorDoors(e))
    /\ \A e \in Elevator : WF_Vars(EnterElevator(e))
    /\ \A e \in Elevator : WF_Vars(ExitElevator(e))
    /\ \A e \in Elevator : SF_Vars(CloseElevatorDoors(e))
    /\ \A e \in Elevator : SF_Vars(MoveElevator(e))
    /\ \A e \in ElevaCall : SF_Vars(DispatchElevator(c))

Spec == \* Initialize state with Init and transition with Next, subject to TemporalAssumptions
    /\ Init
    /\ [][Next]_Vars
    /\ TemporalAssumptions

THEOREM Spec => [](TypeInvariant /\ SafetyInvariant /\ TemporalInvariant)

=============================================================================

