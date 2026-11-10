
---- MODULE Test ----
EXTENDS TestPrelude
op == {
  A!B!&(x, y),
  A!B!&&(x, y),
  A!B!\approx(x, y),
  A(x, y),
  A!B!\succ(x, y),
  A!Beq(x, y),
  A!B!\uplus(x, y),
  A!B!|(x, y),
  A!B!||(x, y),
  A!B!\wr(x, y)
}
====
---- MODULE TestPrelude ----
VARIABLES x, y
A == INSTANCE M1
====
---- MODULE M1 ----
a @@ b == 0
a \preceq b == 0
a \propto b == 0
a ?? b == 0
a |- b == 0
a |- b == 0
aa - b == 0
a -- b == 0
a ^^ b == 0
a \prec b == 0
a \preceq b == 0
a \propto b == 0
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@a ?? b == 0
a |= b == 0
a |- b == 0
a \sim b == 0
a \simeq b == 0
a / b == 0
a // b == 0
a \sqcap b == 0
a \s== 0
a \succ b == 0
a \succeq b == 0
a \supset b == 0
a \supseteq b == 0
a \uplus b == 0
a | b == 0
a || b == 0
a \wr b == 0
====
