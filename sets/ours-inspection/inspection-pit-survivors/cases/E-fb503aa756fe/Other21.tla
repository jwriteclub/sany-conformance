\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: OpApplNode.java:1094 RemoveConditional_EQUAL_IF: D1 error goes
\* Aim: OpApplNode.java:1095 RemoveConditional_EQUAL_IF: D2 error goes
\* Aim: OpApplNode.java:1096 RemoveConditional_EQUAL_IF: D3 error goes
\* Aim: OpApplNode.java:1098 RemoveConditional_EQUAL_ELSE: D4 error goes
\* Aim: OpApplNode.java:1126 RemoveConditional_EQUAL_ELSE: D5 gets an action-level bound error
\* Aim: OpApplNode.java:1127 RemoveConditional_EQUAL_IF: D5 gets an action-level bound error
\* Aim: Location.java:374 RemoveConditional_EQUAL_IF [17]: bLine ignored in equals: D9's two messages differ only there, one is dropped
\* Aim: Location.java:374 RemoveConditional_EQUAL_IF [22]: bColumn ignored in equals: D6's two messages differ only there, one is dropped
\* Aim: Location.java:374 RemoveConditional_EQUAL_IF [27]: eLine ignored in equals: D7's two messages differ only there, one is dropped
\* Aim: TheoremNode.java:286 RemoveConditional_EQUAL_ELSE: proof of the QED step not checked: non-constant HAVE error goes
\* Aim: TheoremNode.java:287 VoidMethodCall: same
---- MODULE Other21 ----
VARIABLE x
D1 == []x \/ x'
D2 == []x => x'
D3 == []x <=> x'
D4 == \/ []x
      \/ x'
D5 == [i \in x' |-> []x]
D6 == x' \/ \E i \in {} : []x \/ x'
D7 == []x \/ x'
      \/ x'
D9 == x' \/ \E i \in {} :
      []x \/ x'
THEOREM []x => []x
<1>1. QED
  <2>1. HAVE x
  <2>2. QED
====
