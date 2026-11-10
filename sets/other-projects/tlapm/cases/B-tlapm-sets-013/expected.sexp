
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq) definition:
    (set_filter
      generator: (quantifier_bound
        intro: (identifier) (set_in)
        set: (conj_list
          (conj_item (bullet_conj) (identifier_ref))
          (conj_item (bullet_conj) (identifier_ref))
          (conj_item (bullet_conj) (identifier_ref))
        )
      )
      filter: (conj_list
        (conj_item (bullet_conj) (identifier_ref))
        (conj_item (bullet_conj) (identifier_ref))
        (conj_item (bullet_conj) (identifier_ref))
      )
    )
  )
(double_line)))

