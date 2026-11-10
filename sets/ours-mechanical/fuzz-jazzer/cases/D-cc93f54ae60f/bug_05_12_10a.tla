--------------------------- MODULE bug_05_12_10a ---------------------------
EXTENDS Naturals, Sequences, FiniteSets , TLC

alg == [type |-> "uniprocess",
        body  |-> <<1>>]

(* 
--algorithm Pcal
             THEN /\ Assert(exp[i] \in STRING, 
                            "Failure of assertion aock on termination. *)
Terminating == pc = "Done" /\ UNCHANGED vars

\* END TRANSLATION - the hash of the generated TLA code (remove to silence divergence warnings): TLA-1e7f8dbaa6644662e3fa38b76045e01d
=============================================================================
