--------------------------------- MODULE XY ---------------------------------
THEOREM Type == Init /\ [][Next]_vars => []TypeOK
<2>2. TypeOK /\ [Next]_vars => TypeOK'
  BY DEF Next, vars, TypeOK
<1>. QED  BY <1>1, <1>2, PTL

