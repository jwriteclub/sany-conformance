
(source_file (module (header_line) (identifier) (header_line)
  (operator_definition (identifier) (def_eq)
    (let_in
      (recursive_declaration (operator_declaration (identifier) (placeholder)))
      (operator_definition (identifier) (identifier) (def_eq)
        (if_then_else
          (bound_infix_op (identifier_ref) (leq) (nat_number))
          (nat_number)
          (bound_infix_op (identifier_ref) (mul)
            (bound_op (identifier_ref)
              (bound_infix_op (identifier_ref) (minus) (nat_number))
            )
          )
        )
      )
      (bound_op (identifier_ref) (nat_number))
    )
  )
(double_line)))

