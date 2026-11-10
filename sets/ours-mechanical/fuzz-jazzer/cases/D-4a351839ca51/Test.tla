
---- MODULE Test ----
EXTENDS TestPrelude
op == f(
  ^+,
  ^*,
  ^#,
  '
)
====
---- MODULE TestPrelude ----
CONSTANTS _^#, _^*, _^#
f(a(_), b(_), c(_), d(_)) == 0
====
