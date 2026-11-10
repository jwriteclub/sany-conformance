----- MODULE Target ----
EXTFNDS == z
z \approx y == z  
z \cong y == z
z \sqsubset y == z
z \sqsupset y == z
z \sqsubseteq y == z
z \sqsupseteq y == z
z \doteq y ==  x' = (x <: 1)
   \/ x' = (x := 1)
   \/ x' = (x ::= 1)
   \/ x' = (x \uplus 1)
   \/ x' = (x \sqcap 1)
   \/ x' = (x \sqcup 1)
   \/ x' = x \div 1
   \/ x' = x \wr 1
   \/ x' = x \star 1
   \/ x' = x \bigcirc 1
   \/ x' = x \bullet 1
   \/ x' = (x \prec 1 )
   \/ x' = (x \succ 1)
   \/ x' = (x \preceq 1)
   \/ x' = (x \succeq 1)
   \/ x' = (x \sim 1)
   \/ x' = (x \simeq 1)
   \/ x' = (x \ll 1)
   \/ x' = (x \gg 1)
   \/ x'  y == z
z < y == z
z > y == z
z .. y == z
z ... y == z
z | y == z
z || y == z
z =| y == z
z <: y == z  
\* z :> y == z  \* THIS IS DEFINED IN = (x \sqsubseteq 1)
   \/ x' = (x \sqsupseteq 1)
   \/ x' = (x \doteq 1)
   \/ x' = (x \propto 1)
   \/ x' = x ^+ 
   \/ x' = xTLC
A(f(_), a) == f(a) + Len(a)
B(g(_), b) == A(g, b)

C == B(Len, "x") /\ B(Len, "y")
====--- MODULE Target ----
EXTENDS TLC
A(f(_), a) \oslash== f(a) + Len(a)
B(g(_), b) == A(g, b)
