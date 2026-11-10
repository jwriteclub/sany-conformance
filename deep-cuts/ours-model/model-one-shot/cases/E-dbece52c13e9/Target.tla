---- MODULE Target ----
EXTENDS Naturals

r == [a |-> 1, b |-> 2]

\* Valid bang chains (bang followed by identifier selectors): these take
Bad == r!a!:
====
