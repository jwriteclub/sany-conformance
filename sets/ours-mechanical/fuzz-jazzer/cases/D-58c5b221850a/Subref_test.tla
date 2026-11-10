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
    <*>. "as above, but label arguments non-trivially permuted"
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
