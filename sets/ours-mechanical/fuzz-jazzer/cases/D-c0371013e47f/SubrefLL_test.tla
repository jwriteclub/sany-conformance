,----- ---- MODUL%N_NonL(* Test of subexpression references from Le---- MODULE SubrefLL_test ----
EXTENDS Integers, TLAPS
ASSUME c4 == C = 4
USE c4

(sStla) == s = s
Op(Arg(_), p) ==
   \A x \in {1,2}, y, z \in {<<3, 4>>} :
                      lab(x,i,z) ::  LET a ++ b ==
                                          Arg(a) + b + x + 2*y + 3*z
                                     IN  (Arg(1) + Arg(p) + C +
                                            2*x  + 4*y + 6*z) =
              US

THEOREM COp(1, 2)!(3) = <<1, {99, 1, 2, 3} >> /\ Salt("3e57a\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\")

THEOREM COp(1, 2)!(3)!1!>> = {99, 1, 2, 3} /\ SalEFUNCTION(x)=======
