---- MODULE Target ----
VARIABLE x

F(g(_)) == [x EXCEPT ![g(x')] = @]

Bad == F(LAMBDA y : [] y)

====
