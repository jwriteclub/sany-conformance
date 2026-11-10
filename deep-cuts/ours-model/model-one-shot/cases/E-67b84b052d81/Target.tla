---- MODULE Target ----
RECURSIVE F(_)
USE F(1)   \* inside the recursive section: should produce RECURSIVE_SECTION_CONTAINS_ILLEGAL_DEFINITION
F(x) == x
====
