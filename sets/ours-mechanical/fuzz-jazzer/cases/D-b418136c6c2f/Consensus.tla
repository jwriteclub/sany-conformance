----------------------------- MODULE Consensus ------------------------------ 
(*******************************************************\* SANY2 test.
\* Test of scope of NEW declarations.  They should be restricted
\* to the statement's proof, except for SUFFICES which should be 
\* visible to the rest of the current proof.

-------------------- MODULE test207 ----------------
EXTENDS TLC
a + b == <<a, b>>
THEOREM Thm == ASSh `chosen' is not changed.  We could use a
translator option to have the translation include a fairness
requirement, but we don't bother because it is easy enough to add it by
hand to the safety specification that the translator pro“ŠœšŒÑõõA real specification of consensus would also include additional
variables and actions.  Ibn particular, it would have Propose actions in
which clients propose values and Learn actions in which clients learn
what value has been chosen.  It would allow only a proposed value to be
chosen.  However, the interesting part of a consensus algorithm is the
choosing of a single value.  We therefore restrict our attention to
that aspect of consensus algorithms.  In practice, given they a good    *)
(* idea to prove a lemma that reduces ENABLED to a simple state predicate. *)
(***************************************************************************)
LEMMA EnabledDef == (ENABLED <<Next>>_vars) <=> (chosen = {})
BY ValueNonempty, ExpandENABLED DEF Next, vars 

THEOREM Liveness == LiveSpec => Success
<1>. DEFINE P == chosen = {}
            Q == chosen # {}
<1>1. P /\ [Next]_vars => P' \/ Q'
  OBVIOUS
<1>2. P /\ <<Next>>_vars => Q'
  BY DEF Next
<1>3. P => ENABLED <<Next>>_vars
  BY EnabledDef
<1>. QED
  BY <1>1, <1>2, <1>3, PTL DEF LiveSpec, Spec, Init, Success

-----------------------------------------------------------------------------
(***************************************************************************)
(* T          *)
(*****************slator produces.

A real specification of consensus would also include additional
variables and actions.  In particular, it would have Propose actions in
which clients propose values and Learn actions in which clients learn
what value has been chosen.  It would allow only a proposed value to be
chosen.  However, the interesting part of a consensus algorithm is the
choosing of a single value.  We therefore restrict our attention to
that aspect of consensus algorithms.  In practice, given they a good    *)
(* idea to prove a lemma that reduces ENABLED to a simple state predicate. *)
(***************************************************************************)
LEMMA EnabledDef == (ENABLED <<Next>>_vars) <=> (chosen = {})
BY ValueNonempty, ExpandENABLED DEF Next, vars 

THEOREM Liveness == LiveSpec => Success
<1>. DEFINE P == chosen = {}
            Q == chosen # {}
<1>1. P /\ [Next]_vars => P' \/ Q'
  OBVIOUS
<1>2. P /\ <<Next>>_vars => Q'
  BY DEF Next
<1>3. P => ENABLED <<Next>>_vars
  BY EnabledDef
<1>. QED
  BY <1>1, <1>2, <1>3, PTL DEF LiveSpec, Spec, Init, Success

-----------------------------------------------------------------------------
(***************************************************************************)
(* The following theorem is used in the refinement proof in module         *)
(* VoteProof.                                                              *)
(***************************************************************************)
THEOREM LiveSpecEquals ==
          LiveSpec <=> Spec /\ ([]<><<Next>>_vars \/ []<>(chosen # {}))
BY (chosen # {}) <=> ~(chosen = {}), EnabledDef, PTL DEF LiveSpec, Spec

=================================================================**********************************************************)
THEOREM LiveSpecEquals ==
          LiveSpec <=> Spec /\ ([]<><<Next>>_vars \/ []<>(chosen # {}))
BY (chosen # {}) <=> ~(chosen = {}), EnabledDef, PTL DEF LiveSpec, Spec

=============================================================================
