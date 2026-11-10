---- MODULE CommentEdges ----
\* Aim: SyntaxTreeNode.java:686, 698, 708, 721, 725, 766, 801, under XMLExporter -u only. The gate does not run -u.
\* Why each comment group: pre-comments are joined and then unboxed. See reference/hand-targets.md.
(** no closing marker *)
A == 1
(** two
lines, no closing marker *)
B == 2
(***************
(* boxed, content before the end *)
(*)
*)  plain middle line
(* short *)
no markers at all here
***************)
C == 3
(***************)
\* a line comment inside
(*** (** second opener *) ***)
(***************)
D == 4
\*
\*no space
\* one space
E == 5
(** single line **)
F == 6
(* opener not on the first line
(** box
**)
last line *)
G == 7
\* a line comment comes first
(***************)
(* boxed *)
(***************)
H == 8
(* ends in the closing marker **)
(** starts with the opening marker *)
I == 9
(*****)
(**)
(*****)
J == 10
====
