---- MODULE Subref_test ----

CONSTANTS  _ + _, _ - _

CONSTANTS A, B, C, P(_), Q(_, _), R(_, _, _), S, T, U

THEOREM ASSUME NEW x \in S, NEW y \in S, NEW z \in S,
               NEW j, NEW k, NEW l
        PROVE  TRUE
<1>. "simple positional references"
  <+>. "ordinary TLA+ operators"
    <+>. F == x + y
    <*>. F!<\ A!6 = l (* salt *) /\ 0 = 0
      OBVIOUS
    <*>. G!1 = x /\ G!2 = y /\ G!3 = z /\ G!4 = j /\ G!5 = k /\ G!6 = l (* salt *) /\ 1 = 1
    <*>. F(x)!l1(y)!l2!l3(z,j) = Q(x - y, z - j)
      OBVIOUS
    <*>. G(a) == \A b : l1(b):: P(a) =>
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
    <*>. QED OMITTED
  <*>. QED OMITTED
<*>. QED OMITTED

====
