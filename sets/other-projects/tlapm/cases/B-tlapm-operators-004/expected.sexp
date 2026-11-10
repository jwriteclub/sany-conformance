
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq) definition:
    (bound_op name: (identifier_ref)
      parameter: (lambda (identifier) (identifier) (bound_infix_op lhs: (identifier_ref) symbol: (plus) rhs: (identifier_ref)))
      parameter: (lambda (identifier) (nat_number))
    )
  )
(double_line)))

