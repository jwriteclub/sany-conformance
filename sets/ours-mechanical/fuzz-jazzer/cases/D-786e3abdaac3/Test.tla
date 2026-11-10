
---- MODULE Test ----
EXTENDS TestPrelude
op == f(
  ~,
  \lnot,
  \neg,
  UNION,
  SUBSET,
  DOMAIN,
  -.,
  ENABLED\in,
  UNCHANGED,
  [],
  <>
)
====
---- MODULE TestPrelude ----
-. x == x
f(a1(_),a1(_),a3(_),a4(_),a5(_),a6(_),a7(_),a8(_),a9(_),a10(_),a11(_)) == C0
====

