---- MODULE ExceptNullSel ----
F == [nosuch EXCEPT ![1] = 2]
G == F!1
====
