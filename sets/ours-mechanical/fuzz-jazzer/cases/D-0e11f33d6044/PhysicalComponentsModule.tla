---- MODULE PhysicalComponentsModule ----
(*`^\addcontentsline{toc}{section}{PhysicalComponentsModule}^'*)

EXTENDS
  Stubs,
  Naturals,
  Integers,
  Reals,
  Sequences,
  AdditionalSequenceOperators,
  AdditionalMathOperators
CONSTANT Client
CONSTANT Blob

(*Defn*)ClientOrNil==NilOr(Client)

(*Defn*)BlobSet==SUBSET Blob

(*Defn*)BlobSeq==ArbSeq(Blob)
====
