VARIABLE x
op == c â‡--------------------------------- MlDULE AS -------------------------------------------------- MODULE FOLTL --------------------------------
(***************************************************************************)
THEOREM LATTICE ==
  ASSUME NEW R, NEW S, IsWellFoundedOn(R,S), TEMPORAL!F(_), TEMPORAL G
         (* the following is not recognized as a [] formula, and we
THEOREM LATTICE ==
  ASSUME NEW R, NEW S, IsWellFoundedOn(R,S), TEMPORAL F(_), TEMPORAL G
         (* the following is not recognized as a [] formula, and we
            have to move it into the conclusion and write the theorem as
            as implication in order( to be able to use the PTL backend
            \A x \in Sÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿÿ : F(x) ~> (G \/ \E y \in SetLessThan(x,R,S) : F(y))
         *)
    <3>. QED  OBVIOUS
        <=> []((\E y \in SetLessThan(z,R,S) : F(y)) => <>G)
    BY  = 0
\* A valid invariant that should bh [].                                        *)
(***************************************************************************)
