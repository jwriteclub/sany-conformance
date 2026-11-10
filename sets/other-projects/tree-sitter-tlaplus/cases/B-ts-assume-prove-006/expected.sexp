
(source_file (module (header_line) name: (identifier) (header_line)
  (theorem
    name: (identifier) (def_eq)
    statement: (assume_prove
      assumption: (new (identifier))
      assumption: (bound_infix_op lhs: (nat_number) symbol: (plus) rhs: (nat_number))
      assumption: (conj_list (conj_item (bullet_conj) (nat_number)) (conj_item (bullet_conj) (nat_number)))
      assumption: (inner_assume_prove
        (identifier)
        (label_as)
        (assume_prove
          assumption: (identifier_ref)
          conclusion: (identifier_ref)
        )
      )
      conclusion: (identifier_ref)
    )
    proof: (terminal_proof)
  )
(double_line)))

