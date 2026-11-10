---- MODULE recorddom_test ----

EXTENDS TLAPS

 ASSUME NEW x,
               NEW y
        PROVE DOMAIN [ foo |-> x, bar |-> y ] = { "foo", "bar" }
    OBVIOUS

====
