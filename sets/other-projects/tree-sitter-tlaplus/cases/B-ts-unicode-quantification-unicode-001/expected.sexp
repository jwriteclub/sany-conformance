
(source_file (module (header_line) name: (identifier) (header_line)
  (operator_definition name: (identifier) (def_eq)
    definition: (bounded_quantification
      quantifier: (forall)
      bound: (quantifier_bound
        intro: (identifier) intro: (identifier)
        (set_in)
        set: (nat_number_set)
      )
      bound: (quantifier_bound
        intro: (identifier)
        (set_in)
        set: (int_number_set)
      )
      bound: (quantifier_bound
        intro: (tuple_of_identifiers (langle_bracket) (identifier) (identifier) (identifier) (rangle_bracket))
        (set_in)
        set: (real_number_set)
      )
      expression: (boolean)
    )
  )
  (operator_definition name: (identifier) (def_eq)
    definition: (bounded_quantification
      quantifier: (exists)
      bound: (quantifier_bound
        intro: (tuple_of_identifiers (langle_bracket) (identifier) (identifier) (rangle_bracket))
        (set_in)
        set: (nat_number_set)
      )
      bound: (quantifier_bound
        intro: (identifier) intro: (identifier)
        (set_in)
        set: (int_number_set)
      )
      expression: (boolean)
    )
  )
(double_line)))

