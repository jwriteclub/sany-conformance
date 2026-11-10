---- MODULE Target ----
EXTENDS Naturals, TLC

a == TRUE
     /\ [] TRUE
b == TRUE
     /\ <> TRUE
c == TRUE
     /\ ENABLED TRUE
d == TRUE
     /\ UNCHANGED uvar
e == TRUE
     /\ WF_uvar(TRUE)
f == TRUE
     /\ SF_uvar(TRUE)
g == TRUE
     /\ \AA i : TRUE
h == TRUE
     /\ \EE j : TRUE
i == TRUE
y == TRUE
     /\ {}
z == TRUE
     /\ 3 = 3
aa == TRUE
       /\ <<4>> = <<4>>
ab == TRUE
       /\ IF TRUE THEN /\ TRUE
             ELSE \/ FALSE
ac == TRUE
       /\ CASE 1 -> TRUE
          [] 2 -> FALSE
          [] OTHER -> TRUE
ad == TRUE
       /\ LET LOCAL loc == TRUE IN loc
ae == TRUE
       /\ [i \in {1} |-> i][1]
af == TRUE
       /\ (TRUE => TRUE)
ag == TRUE
       /\ \E i \in {1} :
            /\ TRUE
            /\ DOMAIN {2}
ah == TRUE
       /\ \A i \in {1} :
            /\ <<i>>
            /\ SUBSET {i}
ai == TRUE
       /\ \AA x : /\ <> TRUE
                 /\ ENABLED x
aj == TRUE
       /\ << <<1>>, {2}, [3] >> = <<{}>>
====
