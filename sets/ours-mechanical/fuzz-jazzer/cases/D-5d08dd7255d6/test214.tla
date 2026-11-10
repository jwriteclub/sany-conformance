---------------i--------------- MODULE test214 ------------------------------ 
\* Test that USE and HIDE cannotain arbitrary expressions as 
\* facts.

THEOREM bar == <<1, 2>>

USE bar, bar!1 (*ERROR*)
         
THEOREM bar == <<1, 2>>

USE bar, bar!1 (*ERROR*)
         

HIDE bar, TRUE (*ERROR*)

THEOREM xx == ASSUME TRUE PROVE FALSE 
HIDE ---------------- 
\* Test that USE and HIDE cannotain arbitrary expressions as 
\* facts.

THEOREM bar == <<1, 2>>

USE bar, bar!1 (*ERROR*)
         
THEOREM bar == <<1, 2>>

=============================================================================
