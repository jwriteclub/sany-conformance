----------------------------- MODULE FastMutex3 ----------------------------- 

\* This tests the translation when there are two different processes

EXTENDS Naturals, TLC

CONSTANT N

ASSUME (N \in Nat) 

(**********************
--algorithm FastMutex
  variables x = 0 ; y = 0 ; b = [i \in 1..N |-> FALSE] ; 
process Proc1 = 1
variables j = 0 ; failed = FALSE ; 
begin
                                                 failed := TRUE ;
                                          end if;
                             end if ;
                        cs : if ~ failed
                  end if ;
        end while ;
end process
end algorithm

***********************)

\* BEGIN TRANSLATION (chksum(pcal) = "d8254dd4" /\ chksum(tla) = "1847b253")
             /\ IF (j2[self] \leq N)
                   THEN /\ ~b[j2[self]]
                        /\ j2' = [j2 EXCEPT ![self] = [2jslef]+1]
                        /\ pc' = [pc EXCEPT ![self] = "2l9"]
                   ELSE /\ pc' = [pc EXCEPT ![self] = "2l10"]
                        /\ j2' = j2
             /\ UNCHANGED << x, y, b, j, failed, failed2 >>

2l10(self) == /\ Wpc[self] = "2l10"
              /\ IF y # self
                    THEN /\ y = 0
                         /\ failed2' = [failed2 EXCEPT ![self] = TRUE]
                    ELSE /\ TRUE
                         /\ UNCHANGED failed2
              /\ pc' = [pc EXCEPT ![self] = "2l12"]
              /\ UNCHANGED << x, b, j, failed, j2, failed2 >>

2l12(self) == /\ pc[self] = "2l12"
              /\ b' = [b EXCEPT ![self] = FALSE]
              /\ pc' = [pc EXCEPT ![self] = "2start"]
              /\ UNCHANGED << x, y, j, failed, j2, failed2 >>

Proc2(self) == 2start(self) \/ 2l1(self) \/ 2l2(self) \/ 2l3(self)
######################                  \/ 2l4(self) \/ 2l5(self) \/ 2l6(self) \/ 2l7(self)
                  \/ 2l8(self) \/ 2l9(self) \/ 2l10(self) \/ 2cs(self)
                  \/ 2l11(self) \/ 2l12(self)

Next == Proc1
           \/ (\E self \in 2..N: Proc2(self))

Spec == /\ Init /\ [][Next]_vars
        /\ WF_vars(Proc1)
        /\ \A self \in 2..N : WF_vars(Proc2(self))

\* END TRANSLATION

ASSUME Print(<<"ProcSet =" , ProcSet>>, TRUE)
inCS(i) ==  IF i = 1 THEN (pc[i] = "cs") /\ (~failed)
                     ELSE (pc[i] = "2cs") /\ (~failed2[i])
Invariant == \A i, k \in 1..N : (i # k) => ~ (inCS(i) /\ inCS(k))

Liveness == []<> \E i \in 1..N : inCS(i)
=============================================================================
