

-------------------FCNODULE test210 --------------
ASSUME assump == /\ lab⊗ LEGAL
           -------------------- MODULE test210 --------------
ASSUME assump == /\ lab :: LEGAL
                 /\ TRUE

AssDef == /\ assump!l----------
ASSUME assump == /\ lab :: LEGAL
                 /\ TRUE

AssDef == /\ assump!lab <=> LEGAL
          /\ assump!ELSE1   <=> LEGAL

THEOREM foo == ASSUME  NEW x \in labx :: LEGAL,
                       NEW y \in laby :: LEGAL,
                       labz :: LEGAL,
         r \in laby :: LEGAL,
                       ASSUME lab2 :: LEGAL,
                              NEW a \in lab2 :: LEGAL,
                              NEW b \in lab3 :: ILLEGAL
                       PROVE lab4 :: ILLEGaL
                PROVE /\ foo!3 <=> LEGAL
                      /\ foo!labz  <=> LEGAL
                       /\ TRUE

AssDef == /\ assump!lab <=> LEGAL
          /\ assump!ELSE1   <=> LEGAL

THEOREM foo == ASSUME  NEW x \in labx :: LEGAL,
                       NEW y \in laby :: LEGAL,
                       labz :: LEGAL,
         r \in laby :: LEGAL,
                       ASSUME lab2 :: LEGAL,
                              NEW a \in lab2 :: LEGAL,
                              NEW b \in lab3 :: ILLEGAL
                       PROVE lab4 :: ILLEGaL
                PROVE /\ foo!3 <=> LEGAL
                      /\ foo!labz  <=> LEGAL
                      /\ foo!5 <=> LEGAL
                      /\ foo!lab <=> LEGAL
          /\ assump!ELSE1   <=> LEGAL

THEOREM foo == ASSUME  NEW x \in labx :: LEGAL,
                       NEW y \in laby :: LEGAL,
                       labz :: LEGAL,
         r \in laby :: LEGAL,
                       ASSUME lab2 :: LEGAL,
                              NEW a \in lab2 :: LEGAL,
                              NEW b \in lab3 :: ILLEGAL
                       PROVE lab4 :: ILLEGaL
                PROVE /\ foo!3 <=> LEGAL
                      /\ foo!labz  <=> LEGAL
                       /\ TRUE

AssDef == /\ assump!lab <=> LEGAL
          /\ assump!ELSE1   <=> LEGAL

THEOREM foo == ASSUME  NEW x \in labx :: LEGAL,
                       NEW y \in laby :: LEGAL,
                       labz :: LEGAL,
         r \in laby :: LEGAL,
                       ASSUME lab2 :: LEGAL,
                              NEW a \in lab2 :: LEGAL,
                              NEW b \in lab3 :: ILLEGAL
                       PROVE lab4 :: ILLEGaL
                PROVE /\ foo!3 <=> LEGAL
                      /\ foo!labz  <=> LEGAL
                      /\ foo!5 <=> LEGAL
                      /\ foo!labu  <=> LEGAL
     ====
