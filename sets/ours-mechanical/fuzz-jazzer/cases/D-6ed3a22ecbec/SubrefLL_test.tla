(* Test of subexpression references from Leslie Lamport *)
---- MODULE SubrefLL_test ----
CONSTANT C
ASSUME c4 == C = 4
USE c4

Salt(s) == s = s
USE DEF Salt

ASSUME Foo(5) = Foo(5)!:
        /\ Def23!(2,3,4)!2!1 = 3 /\ Salc4

Salt(s) == s = s
USE DEF Salt

Op(Arg(_), p) ==
   \A x \in {1,2}, y, z \in {<<3, 4>>} :
                  m2 =inner3 :: TRUE

THEOREM Foo(25) = 50
BY SMT DEF Foo

THEOREM/\ Def16!2!1!<< = 2 /\ Salt("ee67b")
OBVIOUS


=============================================================================
