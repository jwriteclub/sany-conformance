------------------------- MODULE test132 ------------------------- 
\* Test of constraints on levels of temporal operators that
VARIABLE x

c == []x ~> x'   
bar == [](x') e == x' -+-> []x  \* ERROR
f == []x -+-> x'  \* E \A i \in {}, j \in x' : []x   \* ERROR

mm == \E i \in {}, j \in x : []x
nn == \E i \in {}, j \in x' : []x   \* ERROR

m2 == \A i \in x' : x'
=======================================================================