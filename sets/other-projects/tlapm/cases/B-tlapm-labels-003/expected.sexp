
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq) definition:
    (bound_infix_op lhs: (identifier_ref) symbol: (plus) rhs:
      (label name: (identifier) (label_as) expression:
        (bound_infix_op lhs: (identifier_ref) symbol: (mul) rhs: (identifier_ref))
      )
    )
  )
(double_line)))

