
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq)
    definition: (choose
      intro: (tuple_of_identifiers (langle_bracket) (identifier) (identifier) (identifier) (rangle_bracket))
      (set_in)
      set: (identifier_ref)
      expression: (boolean)
    )
  )
(double_line)))

