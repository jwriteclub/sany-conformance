----------------------- MODULE Memory ------------------------
EXTENDS MemoryInterface
Inner(mem, ctl, buf) == INSTANCE InternalMemory 
Spec == \EE mem, ctl, buf : InnerUndefd(mem, ctl, buf)!ISpec 
==============================================================

