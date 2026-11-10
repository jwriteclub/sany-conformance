
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (nat_number_set) (def_eq) definition: (finite_set_literal))
  (operator_definition name: (identifier) (def_eq) definition: (prefixed_op
    prefix: (subexpr_prefix (subexpr_component (identifier_ref)))
    op: (nat_number_set)
  ))
  (operator_definition name: (int_number_set) (def_eq) definition: (finite_set_literal))
  (operator_definition name: (identifier) (def_eq) definition: (prefixed_op
    prefix: (subexpr_prefix (subexpr_component (identifier_ref)))
    op: (int_number_set)
  ))
  (operator_definition name: (real_number_set) (def_eq) definition: (finite_set_literal))
  (operator_definition name: (identifier) (def_eq) definition: (prefixed_op
    prefix: (subexpr_prefix (subexpr_component (identifier_ref)))
    op: (real_number_set)
  ))
(double_line)))
