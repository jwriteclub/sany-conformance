---- MODULE CMP73A ----
ArrayGet(a, i) == a.elems[i + 1]
 [elems: Seq(S)] /\ ArrayLen(buf) <= 3
