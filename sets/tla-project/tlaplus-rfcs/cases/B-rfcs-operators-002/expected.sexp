
(source_file (module (header_line) (identifier) (header_line)
  (operator_definition (identifier) (def_eq)
    (bound_nonfix_op (prefix_op_symbol (negative)) (nat_number))
  )
  (operator_definition (identifier) (def_eq)
    (bound_prefix_op (negative) (parentheses (nat_number)))
  )
  (operator_definition (identifier) (def_eq)
    (bound_infix_op (bound_prefix_op (negative) (nat_number)) (minus) (nat_number))
  )
(double_line)))

