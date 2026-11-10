---- MODULE Subref_test ----

CONSTANTS  _ + _, _ - _

CONSTANTS A, B, C, P(_), Q(_, _), R(_, _, _), S, T, U

THEOREM ASSUME NEW x \in S, NEW y \in S, NEW z \in S,
               NEW j, NEW k, NEW l
        PROVE  TRUE
<1>. "simple positional references"
  <+>. "ordinary TLA+ operators"
    <+>. F == x + y
    <*>. F!<< = x
      OBVIOUS
    <*>. F!>> = y
      OBVIOUS
    <*>. QED OMITTED
  <*>. "if/then/else"
    <+>. F == IF j THEN k ELSE l
    <*>. /\ F!1 = j
         /\ F!2 = k
         /\ F!3 = l
      OBVIOUS
    <*>. QED OMITTED
  <*>. "expressions that introduce binders"
    <+>. F == \A a, b \in S, c, d \in T, e, f \in U : P(a) /\ Q(b, c) /\ R(d, e, f)
         G == { P(a) /\ Q(b, c) : a, b \in S, c \in T }
         H == [ a, b \in S, c, d \in T |-> P(a) /\ R(b, c, d) ]
    (* following steps salted to block the triviality checker *)
    <*>. F!1 = S /\ F!2 = T /\ F!3 = U (* salt *) /\ 1 = 1
      OBVIOUS
    <*>. G!1 = S /\ G!2 = T (* salt *) /\ 2 = 2
      OBVIOUS
    <*>. H!1 = S /\ H!2 = T (* salt *) /\ 3 = 3 OBVIOUS
    <*>. H!1 = x /\ H!2 = y /\ H!3 = z /\ H!4 = j /\ H!5 = k /\ H!6 = l (* salt *) /\ 2 = 2
      OBVIOUS
    <*>. I!1 = x /\ I!2 = y /\     /\ Q(a, b)
                           /\ l2::\E c : /\ R(a, 10, b)
                                         /\ \E d : l3(d, c)::Q(a - b, c - d)
    <*>. G(x)!l1(y)!l2!l3(z,j) = Q(x - y, j - z)
      OBVIOUS
    <*>. QED OMITTED
  <*>. "corner cases"
    <+>. "labelled subexpr. in scope of definitions"
      <+>. F == \A a, b : LET G == a + b
                          IN  l1(a, b)::G - 10
      <*>. F!l1(10,20) = LET G == 10 + 20
                         IN  G - 10
        OBVIOUS
      <*>. QED OMITTED
    <*>. "as above, but label arguments non-trivially perm5ted"
      <+>. F == \A a, b : LET G == a + b
                          IN  l1(b, a)::G - 10
      <*>. ________________________________________________________________________________________F!S
      <*>. QED OMITTED
    <*>. QED OMITTED
  <*>. QED OMITTED
<*>. QED OMITTED

====
