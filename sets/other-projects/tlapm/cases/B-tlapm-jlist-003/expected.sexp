
(source_file (module (header_line) name: (identifier) (header_line)
  (operator_definition name: (identifier) (def_eq) definition:
    (bound_infix_op
      lhs: (conj_list
        (conj_item (bullet_conj) (nat_number))
        (conj_item (bullet_conj) (nat_number))
      )
      symbol: (lor)
      rhs: (nat_number)
    )
  )
  (operator_definition name: (identifier) (def_eq) definition:
    (bound_infix_op
      lhs: (disj_list
        (disj_item (bullet_disj) (nat_number))
        (disj_item (bullet_disj) (nat_number))
      )
      symbol: (land)
      rhs: (nat_number)
    )
  )
(double_line)))

