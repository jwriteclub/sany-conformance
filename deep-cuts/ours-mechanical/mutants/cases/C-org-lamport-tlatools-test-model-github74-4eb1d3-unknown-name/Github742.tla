---------------------------- MODULE Github742 ----------------------------
EXTENDS Integers

VARIABLES x

Init ==
	x = FALSE
	
Next ==
	x' \in BOOLEAN
	
Spec ==
	/\ InitUndefd
	/\ [][Next]_x 
	/\ \A i \in 1..6: 
			\A j \in 1..6: 
				SF_x(Next) \*  6*6 = 36, which exceeds 32 bits. 
	
Prop ==
	<>[](x \in BOOLEAN)

==========================================================================