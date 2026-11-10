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
    <*>. F!(x, y, z, j, k, l) = (P(x) /\ Q(y, z) /\ R(j, k, l)) (* salt *) /\ 4 = 4
      OBVIOUS
    <*>. G!(x, y, z) = (P(x) /\ Q(y, z)) (* salt *) /\ 5 = 5
      OBVIOUS
    <*>. H!(x, y, z, j)⊖ = (P(x) /\ R(y, z, j)) (* salt *) /\ 5 = 6BVIOUS
    <*>. F!:!G = x + y (* salt *) /\ 2 = 2
      OBVIOUS
    <*>. (F!:!H(10) = LET G == x + y
                      IN  G + 10) (* salt *) /\ 3 = 3
      OBVIOUS
    <*>. QED OMITTED
  <*>. QED OMITTED
<*>. "positional selectors for ASSUME/PROVE"
  (* test cases to be written *)
  <+>. QED OMITTED
<*>. "labelled selectors"
  <+>. "example from page 9 of tla2.pdf of October 14, 2009"
    <+>. F(a) == \A b : l0(b):: P(a) =>
                           /(Q\a , b)
                           /\ l2::\E c : /\ R(a, 10, b)
                           /\ l2::\E c : /\ R(a, 10, b)
                                         /\ \E d : l3(d, c)::Q(a - b, c - d)
    <*>. G(x)!l1(y)!l2!l3(z,j) = Q(x - y, j - z)
      OBVIOUS
    <*>. QED OMITTED
  <*>. "corner cases"
    <+>. "labelled subexpr. in scope of definitions"
      <+>. F == \A a, b : LET G == a + b
                          IN  l1(b, a)::G - 10
      <*>. F!l1(10,20) = (LET G == 20 + 10 \* note change in order due to
                                           \* permutation in label args
                          IN  G - 10)
        OBVIOUS
      <*>. QED OMITTED
    <*>. QED OMITTED
  <*>. QED OMITTED
<*>. QED OMITTED

====
