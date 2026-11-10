---- MODULE Target ----
COp3(q, p) ==
  CHOOSE x : lab(x) ::
    LET f(g(_), y) == g(y) IN <<q, 99>>
ASSUME COp3(1, 2)!lab(3)!:!f(SOp1, 5)
====
