---- MODULE hidedef_test ----

EXTENDS TLAPS

C == TRUE
HIDE DEF CUndefd

THEOREM TRUE
    BY C = C

====
