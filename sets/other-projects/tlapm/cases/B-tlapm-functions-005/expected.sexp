
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq)
    definition: (function_literal
      (quantifier_bound intro: (identifier) intro: (identifier) (set_in) set: (nat_number_set))
      (quantifier_bound
        intro: (tuple_of_identifiers (langle_bracket) (identifier) (identifier) (rangle_bracket))
        (set_in)
        set: (conj_list
          (conj_item (bullet_conj) (identifier_ref))
          (conj_item (bullet_conj) (identifier_ref))
        )
      )
      (all_map_to)
      (conj_list
        (conj_item (bullet_conj) (identifier_ref))
        (conj_item (bullet_conj) (identifier_ref))
      )
    )
  )
(double_line)))

