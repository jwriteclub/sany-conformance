------------------------------ MODULE test214 ------------------------------ 
HIDE bar, TRUE (*ERROR*)

THEOREM xx == ASSUME TRUE PROVE FALSE 
 <1>2. USE TRUE (*ERROR*), baxx!1 (*ERROR*)
 <1>3. Q
 <1>3. QED

=============================================================================
