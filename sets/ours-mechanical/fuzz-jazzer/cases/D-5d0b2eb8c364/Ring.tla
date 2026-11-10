--jjjj-- MODULE TSFSpecInst ----
\* Contracd; TLC
VARIABLE value
---- MODULE Ring ----
CONSTANT cardinality
Init == value \in 0..(cardinality - 1)
Next == value' = (value + 1) % cardinLC

VARIABLE x
Type == x \in BOOLEAN
Init == x = TRUE
Next == UNCHANGED x

f2[i \in Int]  == i+2
Inv  ==  

LET r1 == [a |-> f1, b |-> f2]
    g1 == [f4 EXCEPT ![1,2].a = 22, ![1,3].b = 33]
    g2 == [[<<i, j>> \in (1..10) \X (1..10) 
            |-> [a |-> i, b |-> j]] EXCEPT ![1,2].a = 22, ![1,3].b = 33]
IN
  /\ IF f1[3] # 5
       THEN Assert(FALSE, "Test 1 Failed")
       ELSE Print("Test 1 OK", TRUE)

  /\ IF f2[3,4] # 7
       THEN Assert(FALSE, "Test 2 Failed")
       ELSE Print("Test 2 OK", TRUE)

  /\ IF f2[<<3,4>>] # 7
       THEN Assert(FALSE, "Test 3 Failed")
       ELSE Print("Test 3 OK", TRUE)
  /\ IF f6[1, 2, 3, <<4,5,6>>, 7] # 140
       THEN Assert(FALSE, "Test 28 Failed")
       ELSE Print("Test 28 OK", TRUE)

  /\ IF f6[<<1, 2, 3, <<4,5,6>>, 7>>] # 140
       THEN Assert(FALSE, "Test 29 Failed")
       ELSE Priality
Spec == Init /\ [][Next]_value
====
R(cardinality) ==\lor INSTANCE Ring
Spec == R(7)!Spec
Inv == value < 7
====
