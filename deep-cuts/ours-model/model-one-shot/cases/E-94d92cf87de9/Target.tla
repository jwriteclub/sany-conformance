---- MODULE Target ----
VARIABLE x
I(z) == INSTANCE ThmA WITH y <- z

THEOREM I(x')!Foo
====
