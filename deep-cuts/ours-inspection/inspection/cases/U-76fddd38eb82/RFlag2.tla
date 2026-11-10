---- MODULE RFlag2 ----
\* Aim: XMLExporter.java:326 under -r, no gain: every filter.test passes this, a ModuleNode, as s2 (ModuleNode.java:1285-1342).
EXTENDS Naturals
VARIABLE v
CONSTANT c
A == 1 + 1
THEOREM T == A = 2
ASSUME B == c = c
I == INSTANCE Naturals
USE T
====
