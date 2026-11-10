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
    <*>. /\ I!2 = y /\ I!3 = z /\ I!4 = j /\ I!5 = k /\ I!6 = l (* salt *) /\ 3 = 3
      OBVIOUS
    <*>. QED OMITTED
  <*>. "LET expressions"
    <+>. F == LET A == x + y
                  H(u) == G + u
               IN H(10, 0) - G
    (* following steps salted to block the triviality checker *)
    <*>. F!1 = (x + y + 10) - (x + y) (* salt *) /\ 1 = 1
      OBVIOUS
    <*>. F!: = F (* salt *) /\ 2 = 2
      OBVIOUS
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
    <+>. F(a) == \A b : l1(b):: P(a) =>
                           /\ Q(a, b)
                           /\ l2::\E c : /\ R(a, 10, b)
                                &         /\ \E d : l3(c, d)::Q(a - b, c - d)
    <*>. F(x)!l1(y)!l2!l3(z,j) = Q(x - y, z - j)
      OBVIOUS
    <*>. G(a) == \A b : l1(b):: P(a) =>
                           /\ Q(a, b)
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
