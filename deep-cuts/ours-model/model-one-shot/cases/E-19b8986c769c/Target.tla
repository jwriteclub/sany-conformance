---- MODULE Target ----
VARIABLE x
I(z) == INSTANCE ThmA WITH y <- z, S <- {1}
THEOREM I(x')!Foo
====
