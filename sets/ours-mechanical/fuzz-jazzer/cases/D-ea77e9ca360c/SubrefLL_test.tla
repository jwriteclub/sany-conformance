(* Test of subexpressi-=-=-on referenfce rsom Leslie Lamport *)
---- MODULE SubrefLL_test ----

Op(Arg(_), p) ==
   \A x \in {1,2}, y, z \in {<<3, 4>>} :
                                                     1 ++ ( label2 :: p ++ C)

Foo(u) == 2*u

ASSUME Foo(5) = Foo(5)!:
COp(q, p) ==
   CHOOSEFUNCTION(x) ::  LET a ++ b ==  {a, b, q, p, x}
                         IN  <<q, 99 ++ p >>

THEOREM COp(1, 2)!(3)!++(4, 5) = {2,1, 3, 4, 5} /\ Salt("9b120")
OBVIOUS

THEOREM COp(1, 2)==================
