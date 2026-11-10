---- MODULE M ----
EXTENDS Naturals, TLC

THEOREM T1 ==
    TRUE

THEOREM T2 ==
  ASSUME
    STATE -._,
    ACTION -._,
    TEMPORAL t(_, _),
    NEW STATE -._
  PROVE
    TRUE
====