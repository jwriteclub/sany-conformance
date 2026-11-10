
(source_file (module (header_line) name: (identifier) (header_line)
  (operator_definition name: (identifier) (def_eq) definition:
    (bound_prefix_op symbol: (always) rhs: (step_expr_or_stutter
      (bound_infix_op
        lhs: (identifier_ref)
        symbol: (eq)
        rhs: (finite_set_literal (nat_number) (nat_number))
      )
      (identifier_ref)
    ))
  )
  (operator_definition name: (identifier) (def_eq) definition:
    (bound_prefix_op symbol: (always) rhs: (step_expr_or_stutter
      (bound_infix_op
        lhs: (nat_number)
        symbol: (in)
        rhs: (finite_set_literal (nat_number) (nat_number))
      )
      (identifier_ref)
    ))
  )
  (operator_definition name: (identifier) (def_eq) definition:
    (bound_prefix_op symbol: (always) rhs: (step_expr_or_stutter
      (let_in
        definitions:
          (operator_definition
            name: (identifier) (def_eq)
            definition: (finite_set_literal (nat_number) (nat_number))
          )
        expression: (bound_infix_op lhs: (identifier_ref) symbol: (in) rhs: (identifier_ref))
      )
      (identifier_ref)
    ))
  )
  (operator_definition name: (identifier) (def_eq) definition:
    (bound_prefix_op symbol: (always) rhs: (step_expr_or_stutter
      (let_in
        definitions:
          (operator_definition
            name: (identifier) (def_eq)
            definition: (finite_set_literal (nat_number) (nat_number))
          )
        expression:
          (bound_infix_op
            lhs: (identifier_ref)
            symbol: (in)
            rhs: (finite_set_literal (nat_number) (nat_number))
          )
      )
      (identifier_ref)
    ))
  )
  (operator_definition name: (identifier) (def_eq) definition:
    (bound_prefix_op symbol: (always) rhs: (step_expr_or_stutter
      (bound_infix_op lhs: (boolean) symbol: (land) rhs:
        (bound_infix_op
          lhs: (identifier_ref)
          symbol: (in)
          rhs: (finite_set_literal (nat_number) (nat_number))
        )
      )
      (identifier_ref)
    ))
  )
(double_line)))

