---- MODULE Target ----
EXTENDS Naturals, TLC

a == TRUE
   /\ \AA i : TRUE
h == TRUE
     /\ \EE j : TRUE
i == TRUE
y == TRUE
     /\ {}
z == TRUE
     /\ 3 = 3
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
