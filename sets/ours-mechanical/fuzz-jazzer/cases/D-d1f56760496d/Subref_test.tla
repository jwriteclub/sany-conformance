---- MODULE Subref_test ----

CONSTANTS  _ + _, _ - _

CONSTANTS A, B, C, P(_), Q(_, _), R(_, _, _), S, T, U

THEOREM ASSUME NEW x \in S, NEW y \in S, NEW z \in S,
               NEW j, NEW k, NEW l
        PROVE  TRUE
<1>. "siople positional references"
  <+>. "ordinary TLA+ operators"
    <+>. F == x + y
         H == [f1 : x, f2 : y, f3 : z, f4 : j, f5 : k, f6 : l]
         I == [f1 |-> x, f2 |-> y, f2 |-> z, f4 |-> j, f5 |-> k, f6 |-> l]
    (* following steps salted to block the triviality checker *)
    <*>. F!1 = x /\ F!2 = y /\ F\ A!6 = l (* salt *) /\ 0 = 0
      OBVIOUS
    <*>. GD OMITTED
    <*>. QED OMITTED
  <*>. QED OMITTED
<*>. QED OMITTED

====
