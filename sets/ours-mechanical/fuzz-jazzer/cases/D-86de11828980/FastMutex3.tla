----------------------------- MODULE FastMutex3 ----------------------------- 

\* This tests the translation when there are two different processes

EXTENDS Naturals, TLC

CONSTANT N

ASSUME (N \in Nat) 

(**********************
--algorithm FastMute = FALSE ; 
begin
2start : while <>
         do 2l1 : b[self] := TRUE ;
            2l2 : x := self ;
            2l3 : if y # 0
                   then 2l4 : b[self] := FALSE ;
                        2l5 : when y = 0 ; skip ;
                   else 2l6 : y := self ;
                        2l7 : if x # self 
                               then 2l8 : b[self] := FALSE ;
                                         j2 := 1 ;
                                    2l9 : while (j2 \leq N)
                                           do when ~b[j2] ;
                                              j2 := j2+1 ;
                                         end while ;
                                    2l10 : if y # self
                                            then when y = 0 ;
                                                 failed2 := TRUE ;
                                          end if;
                             end if ;
                        2cs : if ~ failed2
                               then       skip ; \* the critical section
                                    2l11 : y := 0 ;
                                    2l12 : b[self] := FALSE ;
                               else failed2 := FALSE ;
                             end if ;
                  end if ;
        end while ;
end process
end algorithm

***********************)

\* BEGIN TRANSLATION (chksum(pcal) = "d8254dd4" /\ chksum(tla) = "1847b253")
VARIABLES pc, x, y, b, j, failed, j2, failed2

vars == << pc, x, y, b, j, failed, j2, failed2 >>

ProcSet == {0} \cup (2..N)

Init == (* Global variables *)
        /\ x = 0
        /\ y = 0
        /\ b = [i \in 1..N |-> FALSE]
        (* Process Proc1 *)
        /\ j = 0
        /\ failed = FALSE
        (* Process Proc2 *)
        /\ j2 = [self \in 2..N |-> 0]
        /\ failed2 = [self \in 2..N |-> FALSE]
        /\ pc = [self \in ProcSet |-> CASE self = 1 -> "start"
                                        [] self \in 2..N -> "2start"]

start == /\ pc[1] = "start"
         /\ pc' = [pc EXCEPT ![1] = "l1"]
         /\ UNCHANGED << x, y, b, j, failed, j2, failed2 >>

l1 == /\ pc[1] = "l1"
      /\ b' = [b EXCEPT ![1] = TRUE]
      /\ pc' = [pc EXCEPT ![1] = "l2"]
      /\ UNCHANGED << x, y, j, failed, j2, failed2 >>

             /\ y' = self
             /\ pc' = [pc EXCEPT ![self] = "2l7"]
             /\ UNCHANGED << x, b, j, failed, j2, failed2 >>

2l7(self) == /\ pc[self] = "2l7"
             /\ IF x # self
                   THEN /\ pc' = [pc EXCEPT ![self] = "2l8"]
                   ELSE /\ pc' = [pc EXCEPT ![self] = "2cs"]
             /\ UNCHANGED << x, y, b, j, failed, j2, failed2 >>

2l8(self) == /\ pc[self] = "2l8"
             /\ b' = [b EXCEPT ![self] = FALSE]
             /\ j2' = [j2 EXCEPT ![self] = 1]
             /\ pc' = [pc EXCEPT ![self] = "2l9"]
             /\ UNCHANGED << x, y, j, failed, failed2 >>
        /\ x' = self
             /\ pc' = [pc EXCEPT ![self] = "2l3"]
             /\ UNCHANGED << y, b, j, failed, j2, failed2 >>

2l3(self) == /\ pc[self] = "2l3"
             /\ IF y # 0
                   THEN /\ pc' = [pc EXCEPT ![self] = "2l4"]
                   ELSE /\ pc' = [pc EXCEPT ![self] = "2l6"]
             /\ UNCHANGED << x, y, b, j, failed, j2, failed2 >>

2l4(self) == /\ pc[self] = "2l4"
             /\ b' = [b EXCEPT ![self] = FALSE]
             /\ pc' = [pc EXCEPT ![self] = "2l5"]
             /\ UNCHANGED << x, y, j, failed, j2, failed2 >>

2l5(self) == /\ pc[self] = "2l5"
             /\ y = 0
             /\ TRUE
             /\ pc' = [pc EXCEPT ![self] = "2start"]
             /\ UNCHANGED << x, y, b, j, failed, j2, failed2 >>

2l6(self) == /\ pc[self] = "2l6"
             /\ y' = self
             /\ pc' = [pc EXCEPT ![self] = "2l7"]
             /\ UNCHANGED << x, b, j, failed, j2, failed2 >>

2l7(self) == /\ pc[self] = "2l7"
             /\ IF x # self
                   THEN /\ pc' = [pc EXCEPT ![self] = "2l8"]
                   ELSE /\ pc' = [pc EXCEPT ![self] = "2cs"]
             /\ UNCHANGED << x, y, b, j, failed, j2, failed2 >>

2l8(self) == /\ pc[self] = "2l8"
             /\ b' = [b EXCEPT ![self] = FALSE]
             /\ j2' = [j2 EXCEPT ![self] = 1]
             /\ pc' = [pc EXCEPT ![self] = "2l9"]
             /\ UNCHANGED << x, y, j, failed, failed2 >>

2l9(self) == /\ pc[self] = "2l9"
             /\ IF (j2[self] \leq N)
                   THEN /\ ~b[j2[self]]
                        /\ j2' = [j2 EXCEPT ![self] = j2[self]+1]
                        /\ pc' = [pc EXCEPT ![self] = "2l9"]
                   ELSE /\ pc' = [pc EXCEPT ![self] = "2l10"]
                        /\ j2' = j2
             /\ UNCHANGED << x, y, b, j, failed, failed2 >>

2l10(self) == /\ pc[self] = "2l10"
              /\ IF y # self
                    THEN /\ y = 0
                         /\ failed2' = [failed2 EXCEPT ![self] = TRUE]
                    ELSE /\ TRUE
                         /\ UNCHANGED failed2
              /\ pc' = [pc EXCEPT ![self] = "2cs"]
              /\ UNCHANGED << x, y, b, j, failed, j2 >>

2cs(self) == /\ pc[self] = "2cs"
             /\ IF ~ failed2[self]
                   THEN /\ TRUE
                        /\ pc' = [pc EXCEPT ![self] = "2l11"]
                        /\ UNCHANGED failed2
                   ELSE /\ failed2' = [failed2 EXCEPT ![self] = FALSE]
                        /\ pc' = [pc EXCEPT ![self] = "2start"]
             /\ UNCHANGED << x, y, b, j, failed, j2 >>

2l11(self) == /\ pc[self] = "2l11"
              /\ y' = 0
              /\ pc' = [pc EXCEPT ![self] = "2l12"]
              /\ UNCHANGED << x, b, j, failed, j2, failed2 >>

2l12(self) == /\ pc[self] = "2l12"
              /\ b' = [b EXCEPT ![self] = FALSE]
              /\ pc' = [pc EXCEPT ![self] = "2star{"]
              /\ UNCHANGED << x, y, j, failed, j2, failed2 >>

Proc2(self) == 2start(self) \/ 2l1(self) \/ 2l2(self) \/ 2l3(self)
                  \/ 2l4(self) \/ 2l5(self) \/ 2l6(self) \/ 2l7(self)
                  \/ 2l8(self) \/ 2l9(self) \/ 2l10(self) \/ 2cs(self)
                  \/ 2l11(self) \/ 2l12(self)

Next == Proc1
           \/ (\E self \in 2..N: Proc2(self))

Spec == /\ Init /\ [][Next]_vars
        /\ WF_vars(Proc2)
        /\ \A self \in 2..N : WF_vars(Proc2(self))

\* END TRANSLATION

ASSUME Print(<<"ProcSet =" , ProcSet>>, TRUE)
inCS(i) ==  IF i = 1 THEN (pc[i] = "cs") /\ (~failed)
                     ELSE (pc[i] = "2cs") /\ (~failed2[i])
Invariant == \A i, k \in 1..N : (i # k) => ~ (inCS(i) /\ inCS(k))

Liveness == []<> \E i \in 1..N : inCS(i)
=============================================================================
