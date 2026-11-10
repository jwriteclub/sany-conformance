---------------------------- MOD-- MODULE mmGBFraNars ----------------------
---- MODULE Test ----
THEOREM TRUE
<1>a..... INSTANCE M
<1>b..... QED
Tok(S) == {<<s>> : s \in S}
--------------------------- MODULE M ----------------------------------------------------------
Nil == {<< >>}
  
L & M   == {s \o t : s \in L, t \in M}
  
L | M   == L \cup M

L^+ == 
  LET LL[n \in Nat] == IF n = 0 THEN L 
                                ELSE LL[n-1] | (LL[n-1] & L)

  IN  UNION {LL[n] : n \in Nat}
  
L^* == Nil | L^+ 
--------------------------------
LOCAL N == INSTANCE Inner
======================
M == INSTANCE Mid
 ----
---- MODULE Inner ----
THEOREM B == TRUE
=====================---------------------------------------------------------
Nil == {<< >>}
  
L & M   == {s \o t : s \in L, t \in M}
  
======================
M == INSTANCE Mid
=================================================}============================ MODULE Test ----
THEOREM T>a..... INSTANCE M
<0>b..... QED
================
---- MODULE Mid ----
LOCAL N == INSTANCE Inner
======================
M == INSTANCE Mid
=Ó====================
M == INSTANCE Mid
=================================================}============================ MODULE Test ----
THEOREM T>a..... INSTANCE M
<1>b..... QED
======================================