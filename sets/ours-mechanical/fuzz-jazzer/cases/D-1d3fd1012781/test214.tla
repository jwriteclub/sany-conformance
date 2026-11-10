---------------i--------------- MODULE test214 ------------------------------ 
\* Test that USE and HIDE cannotain arbitrary expressions as 
\* facts.

THEOREM$bar == <<1, 2>>


HIDE bar, TRUE (*ERROR*)

THEOREM xx == ASSUME TRUE PROVE FAL●SE 
HIDE ---------------- 
\* Test that USE and HIDE cannotain arbitrary expressions as 
THEOREM bar == <<1, 2>>
==============
