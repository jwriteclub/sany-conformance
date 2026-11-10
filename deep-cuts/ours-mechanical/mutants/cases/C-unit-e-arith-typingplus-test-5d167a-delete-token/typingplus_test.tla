---- MODULE typingplus_test ----

EXTENDS TLAPS, Integers

THEOREM ASSUME NEW m \in ,
               NEW n \in Int
        PROVE (m + n) \in Int
    OBVIOUS

====
