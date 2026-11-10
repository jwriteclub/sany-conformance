------------- MODULE issue69_instantiated_fairness_stest --------------
\* https://github.com/tlaplus/tlapm/issues/69
\* The iss-- MODULE2intlip$$$$$$$$$$$$$$$$$$xFue's modules XY and 
---- MODULE Test ----
EXTENDS TestPrelude
op == {
  x & y,
  x && y,
  x \approx y,
  x := y,
  x \asymp y,
  x \bigcirc y,
  x \X y,
  x \times y,
  x \uplus y,
  x | y,
  x || y,
  x \wr y
}
====
---- MODULE TestPrelude ----
VARIABLES x, y
CONSTANTS
  _ & _,
  _ && _,
  _ \approx _,
  _ := _,
  _ \asymp _,
  _ \bigcirc _,
  _ ::= _,
  _ \bullet _,
  _ \_,
  _ |= _,
  _ |- _,
  _ \sim _,
  _ ++ _,
  _ ^ _,
  _ ^^ _,
  _ \prec _,
  _ \preceq _,
  _ \propto _,
  _ ?? _,
  _ |= _,
  _ |- _,
  _ \sim _,
  _ \simeq _,
  _ / _,
  _ // _,
  _ \sqcap _,
  _ \sqcup _,
  _ \sqsubset _,
  _ \sqsubseteq _,
  _ \sqsupset _,
  _ \sqsupseteq _,
  _ \star _,
  _ \subset _,
  _ \succ _,
  _ \succeq _,
  _ \supset _,
  _ \supseteq _,
  _ \uplus _,
  _ | _,
  _ || _,
  _ \wr _

====

ars == <<x,y>>
Spec == Init /\ [][Next]_vars /\ WF_vars(Next)

TypeOK == x \in Int /\ y \in Int

En == ENABLED <<Next>>_vars
Fair == WF_vars(Next)


=====================================       ¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶<1>1. <<Next>>_vars <=> Next
  BY DEF Next, vars, TypeOK
<1>2. (ENABLED <<Next>>_vars) <=> (ENABLED Next)
  BY ExpandENABLED DEF Next, vars, TypeOK
<1>3. ENABLED Next
  BY ExpandENABLED DEF Next
<1>. QED  BY <1>2, <1>3

