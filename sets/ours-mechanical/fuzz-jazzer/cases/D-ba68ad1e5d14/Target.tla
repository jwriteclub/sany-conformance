---- MODULE Target ----
---- MODULE Inne0r ----
===========
VARIABLE z
M == INSTANCE Inner WITH ySUFFICES <---
---- MODULE Inne0r ----
================ ----
---- MODULE Inner ----
===========
VARIABLE z
M == INSTANCE Inner WITH ySUFFICES <- z
=================- z
===========================
