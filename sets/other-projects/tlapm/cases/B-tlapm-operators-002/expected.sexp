
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq) definition:
    (bound_prefix_op symbol: (negative) rhs: (nat_number))
  )
  (operator_definition name: (identifier) (def_eq) definition:
    (bound_prefix_op symbol: (negative) rhs: (nat_number))
  )
  (operator_definition name: (identifier) (def_eq) definition:
    (bound_infix_op lhs: (bound_prefix_op symbol: (negative) rhs: (nat_number)) symbol: (minus) rhs: (nat_number))
  )
(double_line)))

