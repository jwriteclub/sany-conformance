
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq)
    definition: (set_filter
      generator: (quantifier_bound
        intro: (tuple_of_identifiers (langle_bracket) (identifier) (rangle_bracket))
        (set_in)
        set: (nat_number_set)
      )
      filter: (bound_op name: (identifier_ref) parameter: (identifier_ref) parameter: (identifier_ref) parameter: (identifier_ref))
    )
  )
(double_line)))

