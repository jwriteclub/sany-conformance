
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq) definition:
    (let_in
      (recursive_declaration (operator_declaration (identifier) (placeholder)))
      definitions:
        (operator_definition name: (identifier) parameter: (identifier) (def_eq) definition:
          (if_then_else
            if: (bound_infix_op lhs: (identifier_ref) symbol: (leq) rhs: (nat_number))
            then: (nat_number)
            else: (bound_infix_op lhs: (identifier_ref) symbol: (mul) rhs:
              (bound_op name: (identifier_ref) parameter:
                (bound_infix_op lhs: (identifier_ref) symbol: (minus) rhs: (nat_number))
              )
            )
          )
        )
      expression:
        (bound_op name: (identifier_ref) parameter: (nat_number))
    )
  )
(double_line)))

