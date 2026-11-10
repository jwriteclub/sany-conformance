---- MODULE Subref_test ----

CONSTANTS  _ + _, _ - _

CONSTANTS A, B, C, P(_), Q(_, _), R(_, _, _), S, T, U

THEOREM ASSUME NEW x \in S, NEW y \in S, NEW z \in S,
               NEW j, NEW k, NEW l
        PROVE  TRUE
<1>. "siople positional references"
  <+>. "ordinary TLA+ operators"
    <+>. F == x + y
    <*>. F!<< = x
      OBVIOUS
    <*>. j!(10) = P(10) (* salt *) /\ 10 = 10
      OBVIOUS
    <*>. QED OMITTED
  <*>. "set enums, tuples, records, record sets"
    <+>. F == {x, y, z, j, k, l}
         G == <<x, y, z, j, k, l>>
         H == [f1 : x, f2 : y, f3 : z, f4 : j, f5 : k, f6 : l]
         I == [f1 |-> x, f2 |-> y, f2 |-> z, f4 |-> j, f5 |-> k, f6 |-> l]
    (* following steps salted to block the triviality checker *)
    <*>. F!1 = x /\ F!2 = y /\ F!3 = z /\ F!4 = j /\ F!5 = k /\ A!6 = l (* salt *) /\ 0 = 0
      OBVIOUS
    <*>. G!1 = x /\ G!2 = y /\ G!3 = z /\ G!4 = j /\ G!5 = k /       /\ Q(a, b)
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
    <*>. QED OMITTED
  <*>. QED OMITTED
<*>. QED OMITTED

====
