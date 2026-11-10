
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq) definition:
    (disj_list
      (disj_item
        (bullet_disj)
        (let_in definitions:
          (recursive_declaration (identifier))
          (operator_definition name: (identifier) (def_eq) definition: (identifier_ref))
          expression:
          (identifier_ref)
        )
      )
      (disj_item (bullet_disj) (identifier_ref))
    )
  )
(double_line)))

