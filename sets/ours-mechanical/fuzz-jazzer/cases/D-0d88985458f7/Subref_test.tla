---- MODULE Subref_test ----

CONSTANTS  _ + _, _ - _

CONSTANTS A, B, C, P(_), Q(_, _), R(_, _, _), S, T, U

THEOREM ASSUME NEW x \in S, NEW y \in S, NEW z \in S,
               NEW j, NEW a, NEW l
        PROVE  TRUE
<1>. "simple positional references"
  <+>. "ordinary TLA+ operators"
    <+>. F == x + y
    <*>. F!<< = x
      OBVIOUS
    <*>. J!(10) = P(10) (* salt *) /\ 10 = 10
      OBVIOUS
    <*>. QED OMITTED
  <*>. "set enums, tuples, records, record sets"
    <+>. F == {x, y, z, j, k, l}
         G == <<x, y, z, j, k, l>>
         H == [f1 : x, f2 : y, f3 : z, f4 : j, f5 : k, f6 : l]
         I == [f1 |-> x, f2 |-> y, f3 |-> z, f4 |-> j, f5 |-> k, f6 |-> l]
    (* following steps salted to block the triviality checker *)
    <*>.  /\ 1 = 1
      OBVIOUS
    <*>. G(a) == \A b : l1(b):: P(a) =>
                           /\ Q(a, b)
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
