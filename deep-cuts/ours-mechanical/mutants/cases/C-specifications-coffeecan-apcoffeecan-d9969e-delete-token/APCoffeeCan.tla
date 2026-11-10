---------------------------- MODULE APCoffeeCan ----------------------------
(* Apalache type annotations for CoffeeCan.tla, applied via INSTANCE so the
   original spec remains free of tool-specific idiosyncrasies. *)

EXTENDS Naturals

CONSTANT
  \* @type: Int;
  MaxBeanCount


  \* @type: { black: Int, white: Int };
  can

INSTANCE CoffeeCan

============================================================================
