
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq) definition:
    (conj_list
      (conj_item
        (bullet_conj)
        (let_in
          definitions: (module_definition
            name: (identifier) (def_eq)
            definition: (instance (identifier_ref)))
          expression: (identifier_ref)
        )
      )
      (conj_item (bullet_conj) (identifier_ref))
    )
  )
(double_line)))

