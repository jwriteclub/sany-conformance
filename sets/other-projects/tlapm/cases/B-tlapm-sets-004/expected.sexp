
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq)
    definition: (finite_set_literal
      (boolean)
      (boolean)
      (parentheses
        (bound_infix_op
          lhs: (parentheses
            (bound_infix_op lhs: (nat_number) symbol: (plus) rhs: (nat_number))
          ) symbol: (mul)
          rhs: (nat_number)
        )
      )
      (finite_set_literal (nat_number) (nat_number) (nat_number))
    )
  )
(double_line)))

