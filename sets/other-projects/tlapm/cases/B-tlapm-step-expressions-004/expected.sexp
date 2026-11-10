
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq) definition:
    (step_expr_or_stutter
      (conj_list (conj_item (bullet_conj) (nat_number)) (conj_item (bullet_conj) (nat_number)))
      (identifier_ref)
    )
  )
  (operator_definition name: (identifier) (def_eq) definition:
    (step_expr_no_stutter
      (langle_bracket)
      (conj_list (conj_item (bullet_conj) (nat_number)) (conj_item (bullet_conj) (nat_number)))
      (rangle_bracket_sub)
      (identifier_ref)
    )
  )
(double_line)))

