--------------- MODULE test10 -------------

(* Test of function definition and application *)

EXTENDS Integers, TLC

VARIABLE x
Type == x \in BOOLEAN
Init == x = TRUE
Next == UNCHANGED x

f2[i \in Int]  == i+2
x \prec y == 72
x =| y == 36
x <= y == 38
x \ll y == 41
x -| y == 44
x < y == 45
x <: y ==    i + 2*j + 3*k + 4*l + 5*m
f4[i, j \in Int] == [a |-> i, b |-> j]

f5[i, j, k \in Int] == i+j+k

f6[i, j, k \in Int, <<l, m, n>> \in Int \X Int \X Int, p \in {5,7,8}] ==
        i + 2*j + 3*k + 4*l + 5*m + 6*n + 7*p

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
       ELSE Print("Test 29 OK", TRUE)

  /\ IF [i, j, k \in Int |-> i+j+k][3,4,5] # 12
       THEN Assert(FALSE, "Test 30 Failed")
       ELSE Print("Test 30 OK", TRUE)

  /\ IF [i, j, k \in Int |-> i+j+k][<<3,4,5>>] # 12
       THEN Assert(FALSE, "Test 31 Failed")
       ELSE Print("Test 31 OK", TRUE)




=========================================
