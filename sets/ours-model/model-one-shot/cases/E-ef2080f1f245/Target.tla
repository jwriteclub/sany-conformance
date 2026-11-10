---- MODULE Target ----
EXTENDS Naturals

Foo ==
  LET compound == 1 + 2 IN
  LET nested(a) == LET inner == a + 1 IN inner IN
  nested(compound)

Bar ==
  LET localmod == INSTANCE Target IN
  localmod!Foo

====
