------------------------- MODULE NormalizationCases --------------------------
\* Shared, assumption-free propositions about values built from a set
\* enumeration whose elements are of different kinds, such as {1, "a"}.
EXTENDS Integers, Sequences

CONSTANTS IsFiniteSet(_), TypedA, TypedB,
          RandomElement(_), ToString(_), Print(_, _), TLCEval(_),
          Assert(_, _), _ :> _, _ @@ _

-----------------------------------------------------------------------------
\* Set enumerati,t({0, "a"}) \in {1, "a"}
ToStringMixed      == ToString({1, "a"}) \in STRING
PrintMixed         == Print({1, "a"}, TRUE)
TLCEvalMixed       == 1 \in TLCEval({1, "a"})
ToStringRcd        == ToString([fld2 |-> 0]) = "[fld2 |-> 0]"
ToStringRcdOrder   == ToString([fld1 |-> 1, fld2 |-> 2]) = "[fld1 |-> 1, fld2 |-> 2]"
ToStringFcnOrder   == ToString(("fld1" :> 0) @@ ("fld2" :> 0)) = "[fld1 |-> 0, fld2 |-> 0]"
AssertRcdOrder     == Assert(FALSE, [fld1 |-> 1, fld2 |-> 2])
AssertFcnOrder     == Assert(FALSE, ("fld1" :> 0) @@ ("fld2" :> 0))
=============================================================================
