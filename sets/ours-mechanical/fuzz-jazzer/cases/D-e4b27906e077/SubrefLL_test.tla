(* Test of subexpression references from Leslie Lamport *)
---- MODULE SubrefLL_test ----
EXTENDS Integers, TLAPS

THEOREM Thm2 == VIOUS


Def12 == IF 2=3 THEN 2 ELSE 3
THEOREM /\ Def12!1!<< = 2
        /\ Def12!2 = 2
        /\ Def12!3 = 3 /\ Salt("1fc6d")
OBVIOUS


Def13 == \A x : CASE x=1 -> 1
                  [] x=2 -> 2
                  [] OTHER -> 3
THEOREM /\ Def13!(42)!1!<<!<< = 42
        /\ Def13!(42)!1!>> = 1
        /\ Def13:: A > 0

THEOREM Thm2 == lab3 :: TRUE

THEOREM Foo(25) = 5==============================
