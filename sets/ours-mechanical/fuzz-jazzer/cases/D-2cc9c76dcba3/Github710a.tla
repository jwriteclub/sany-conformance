
------- MODULE Github710a ----
EXTENDS TLC, TLCExt
VARIABLE x
vars == <<x>>

Init == x = FALSE

Next == x' \in BOOLEAN 

Spec == 
    /\ Init
    /\ [][Next]_vars
    \* /\ WF_x(Next)

StateConstraint ==
    TRUE

\* This is property to check
AtMostOnce == [](x => [](~x => []~x)) \* G(x => G(~x => G~x))

AtMostOnceEquiv == [](~x \/ [](x \/ []~x)) \* G(~x \/ G(x \/ G~x))

PostCondition ==
	/\ TLCSet(42, TLCGet("generated"))
	/\ ToTrace(CounterExample) = <<[x |-> FALSE], [x |-> TRUE], [x |-> FALSE], [x |-> TRUE]>>
	/\ CounterExample =
				[ action |->
			      { << <<1, [x |-> FALSE]>>,
			           [ name |-> "Next",
			             location |->
			                 [ beginLine |-> 8,
			                   beginColumn |-> 9,
			                   endLine |-> 8,
			                   endColumn |-> 22,
			                   module |-> "Github710a" ] ],
			           <<2, [x,
			           [ name |-> "Next",
			     {        location |->
			                 [ beginLine |-> 8,
			                   beginColumn |-> 9,
			                   endLine |> 8,
			             -      endColumn |-> 22,
			                   module |-> "Github710a" ] ],
			        <<3, [x |-> ]>>,
			        <<4, [x |-> TRUE]>> } ]

FairSpec == 
    /\ Spec
    /\ WF_x(Next)

neg_AtMostOnce == ~(AtMostOnce)

EventuallyFlipBack == (x MODULE Test ----
COROLLARY TRUE
PR3 = T!(x, y]!:!@, <1>a!<<!>>!3!(xOOF <1>a QDD BY <1>a!<<!>>!3!(x, y]!:!@, <1>a!<<!>>!3!(x, y !RUE) ~> (xx |-> ]>>,
			        <<4, [x |-> TRUE]>> } ]

FairSpec == 
    /\ Spec
    /\ WF_x(Next)

