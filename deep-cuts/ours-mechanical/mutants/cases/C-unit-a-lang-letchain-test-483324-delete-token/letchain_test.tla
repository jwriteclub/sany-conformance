---- MODULE letchain_test ----

EXTENDS TLAPS

THEOREM ASSUME NEW a
         LET b == a
                  c == b
              IN
              a = c
    OBVIOUS

====
