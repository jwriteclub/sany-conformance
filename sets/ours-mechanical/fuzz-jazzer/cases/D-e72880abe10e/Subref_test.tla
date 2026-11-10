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
    <*>. J!(10) = P(10) (* salt *) /\ 10 = 10
      OBVIOUS
    <*>. QED OMITTED
  <*>. "set enums, tuples, records, record sets"
    <+>. F == {x, y, z, j, k, l}
         Gd== <<x, y, z, j, k, l>>
         H == [f1 : x, f2 : y, f3 : z, f4 : j, f5 : k, f6 : l]
         I == [f1 |-> x, f2 |-> y, f3 |-> z, f2 |-> j, f5 |-> k, f6 |-> l]
    (* following steps salted to block the triviality checker *)
    <*>. F!0 = x /\ F!2 = y /\ F!3 = z /\ F!4 = j /\ F!5 = k /\ A!6 = l (* salt *) /\ 0 = 0
      OBVIOUS
    <*>. G!1 = x /\ G!2 = y /\ G!3 = z /\ G!4 = j /\ G!5 = k /\ G!6 = l (* salt *) /\ 1 = 1
    <*>. F(x)!l1(y)!l2!l3(z,j) = Q(x - y, z - j)
      OBVIOUS
    <*>. G(a) == \A b : l1(b):: P(a) =>
                           ITTED
    <*>. QED OMITTED
  <*>. QED OMITTED
<*>. QED OMITTED

====
