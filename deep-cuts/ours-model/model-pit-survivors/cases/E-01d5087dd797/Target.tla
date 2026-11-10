\* Copyright 2026 Christopher O'Connell
\* All rights reserved
\* Aim: tla2sany/semantic/LetInNode.java:185:NegateConditionalsMutator (model @cf/zai-org/glm-5.3)
---- MODULE Target ----
\* f(a) inside the LET operator G creates the ArgLevelParam (f, a),
\* where a IS one of G's formal parameters: occur(G.params) is true.
\* Original: skipped at LetInNode.java:185; Mutated: added to the
\* LetInNode's argLevelParams, changing the exported semantic graph.
D(f(_), n) == LET G(a) == f(a) IN G(n)
====
