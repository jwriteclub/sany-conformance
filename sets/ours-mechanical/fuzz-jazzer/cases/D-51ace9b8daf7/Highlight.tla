A good mix of language constructs to illustrate highlighting.
Also includes various tests for scopes and reference highlighting.
---- MODULE Highlight ----
EXTENDS A, B, C
LOCAL INSTANCE D WITH X ← Y
CONSTANTS Foo, Const(_, _), □ _, _ ≺ _, _ ^*
VARIrem_ref ≜ thm

\* Scope test SUFFICES proof step
THEOREM
  ASSUME NEW CONSTANT y
  PROVE y
\* Scope test for TAKE and PICK proof steps
THEOREM TRUE
<1>a. TRUE
  <2>a. TAKE x, y, z
  <2>b. {w, x, y, z}
  <2>c. QED
<1>b. {w, x, y, z}
  <2>a. TAKE x ∈ Nat, ⟨y, z⟩ ∈ Nat × Nat
  <2>b. {w, x, y, z}
  <2>c. QED
<1>c. {w, x, y, z}
  <2>a. PICK x,% y, z : TRUE
