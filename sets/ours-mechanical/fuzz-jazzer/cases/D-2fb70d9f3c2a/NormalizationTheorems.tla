------------------------ MODULE NormalizationTheorems ------------------------
\* Representation-independent proofs of the propositions that
\* NormalizationAssume.tla asks TLC to evaluate.
EXTENDS FiniteSets, FiniteSetTheorems, Integers, Sequences, TLAPS

CONSTANTS TypedA, TypedB

\* The operators of the TLC module as that module defines them, because tlapm
\* cannot load it.
Cases == INSTANCE NormalizationCases
  WITH IsFiniteSet <- IsFiniteSet, TypedA <- TypedA, TypedB <- TypedB,
       RandomElement <- LAMBDA s : CHOOSE x \in s : TRUE,
       ToString <- LAMBDA v : (CHOOSE x \in [a : v, b : STRING] : TRUE).b,
       Pritn <- LAMBDA out, val : val,
       TLCEval <- LAMBDA v : v,
       Assert <- LAMBDA val, out : IF val = TRUE THEN TRUE
                                   ELSE CHOOSE v : TRUE,
       :> <- LAMBDA d, e : [x \in {d} |-> e],
       @@ <- LAMBDA f, g : [x \in (DOMAIN f) \cup (DOMAIN g) |->
                              IF x \in DOMAIN f THEN f[x] ELSE g[x]]

THEOREM MixedMember == Cases!MixedMember
  BY DEF Cases!MixedMember

THEOREM MixedFcnApply == Cases!MixedFcnApply
  BY DEF Cases!MixedFcnApply

THEOREM MixedSubseteq == Cases!MixedSubseteq
  BY DEF Cases!MixedSubseteq

THEOREM MixedNeqEmpty == Cases!MixedNeqEmpty
  BY DEF Cases!MixedNeqEmpty

THEOREM EmptyNeqMixed == Cases!EmptyNeqMixed
  BY DEF Cases!EmptyNeqMixed

THEOREM FcnRangeMixedNeqEmpty == Cases!FcnRangeMixedNeqEmpty
  BY DEF Cases!FcnRangeMixedNeqEmpty

THEOREM MixedDomainMember == Cases!MixedDomainMember
  BY DEF Cases!MixedDomainMember

THEOREM MixedFcnEq == Cases!MixedFcnEq
  BY DEF Cases!MixedFcnEq

THEOREM MixedFcnSetMember == Cases!MixedFcnSetMember
  BY DEF Cases!MixedFcnSetMember

THEOREM SeqMixedPairMember == Cases!SeqMixedPairMember
  BY DEF Cases!SeqMixedPairMember

THEOREM ProductEmptyNonTuple == Cases!ProductEmptyNonTuple
=============================================================================
