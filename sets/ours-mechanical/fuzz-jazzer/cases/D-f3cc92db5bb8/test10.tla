--------------- MODULE test10 -------------

(* Test of function definition and application *)

EXTENDS Integers, TLC

VARIABLE x
Type == x \in BOOLEAN
Init == x = TRUE
Next == UNCHANGED x

f2[i \in Int]  == i+2

f6[i, j, k \in Int, <<l, m, n>> \in Int \X Int \X Int, p \in {5,7,8}] ==
     /\ IF f2[3,4] # 7
       THEN Assert(FALSE, "Test 2 Failed")
       ELSE Print("Test 2 OK", TRUE)


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
