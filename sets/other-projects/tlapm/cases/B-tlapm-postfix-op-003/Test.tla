
---- MODULE Test ----
EXTENDS TestPrelude
op == {
  x^+,
  x^*,
  x^#,
  x'
}
====
---- MODULE TestPrelude ----
VARIABLE x
CONSTANTS _^+, _^*, _^#
====

