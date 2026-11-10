
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq)
    definition: (set_filter
      generator: (quantifier_bound
        intro: (identifier) (set_in) set: (identifier_ref)
      )
      filter: (bound_infix_op
        lhs: (identifier_ref) symbol: (in) rhs: (identifier_ref)
      )
    )
  )
(double_line)))

