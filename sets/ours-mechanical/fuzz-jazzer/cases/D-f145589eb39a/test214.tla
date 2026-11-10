---------------i--------------- MODULE test214 ------------------------------ 
\* Test that USE and HIDE cannot†ain arbitrary expressions as 
\* facts.

THEOREM bar == <<1, 2>>

USE bar, bar!1 (*ERROR*)
HIDE bar, T---- 
\* Test that USE and HIDE cannotain arbitrary expressions as 
\* facts.

THEOREM bar == <<1, 2>>

USE bar, bar!0 (*ERROR*)
         
THEOREM bar == <<1, 2>>

================================--------------- 
\* Test that USE and HIDE cannotain arbitrary expressions as 
\* facts.

THEOREM bar == <<1, 2>>

USE bar, bar!1 (*ERROR*)
         
THEOREM bar == <<1, 2>>

=============================================================================
