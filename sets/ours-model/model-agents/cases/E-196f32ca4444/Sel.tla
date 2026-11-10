---- MODULE Sel ----
EXTENDS Naturals
CONSTANT cc

M1 == INSTANCE Naturals
Foo(a, b) == a + b
A2(f(_)) == f(1)
Lm == A2(LAMBDA x : x + 1)
UseI2 == ThmI!2!1
BadExpr == M1

THEOREM TAP == ASSUME NEW x \in Nat PROVE x >= 0

THEOREM TT == 1 = 1
  <1>1. 1 = 1
  <1>2. DEFINE dd == 2
  <1>3. PICK xx : xx = 1
    <2>1. 1 = 1
    <2>2. QED BY <2>1
  <1>4. 2 = 2 BY <1>3
  <1>5. QED BY <1>1, <1>2 DEF dd, <1>4

USE M1
USE Foo(1, 2)
USE Foo(1)
USE cc
USE ThmA!1!2
USE ThmA!1!3
USE ThmA!2
USE ThmA!:!1
USE ThmB
USE ThmC!1
USE TAP
USE TAP!1
USE TAP!2
USE ThmI!L1
USE ThmJ!A1!B1
USE ThmJ!1
USE ThmD!1
USE ThmD!1!1
USE ThmE!1
USE ThmF!1

HIDE DEF Foo
HIDE DEF M1
HIDE DEF cc
HIDE DEF ThmB
HIDE DEF TAP
====