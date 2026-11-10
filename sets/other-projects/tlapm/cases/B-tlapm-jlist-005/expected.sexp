
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq) definition:
    (conj_list
      (conj_item (bullet_conj) (label name: (identifier) (label_as) expression: (identifier_ref)))
      (conj_item (bullet_conj) (bound_infix_op
        lhs: (identifier_ref)
        symbol: (assign)
        rhs: (identifier_ref)))
      (conj_item (bullet_conj) (bound_infix_op
        lhs: (identifier_ref)
        symbol: (bnf_rule)
        rhs: (identifier_ref)))
    )
  )
  (operator_definition name: (identifier) (def_eq) definition:
    (disj_list
      (disj_item (bullet_disj) (label name: (identifier) (label_as) expression: (identifier_ref)))
      (disj_item (bullet_disj) (bound_infix_op
        lhs: (identifier_ref)
        symbol: (assign)
        rhs: (identifier_ref)))
      (disj_item (bullet_disj) (bound_infix_op
        lhs: (identifier_ref)
        symbol: (bnf_rule)
        rhs: (identifier_ref)))
    )
  )
(double_line)))

