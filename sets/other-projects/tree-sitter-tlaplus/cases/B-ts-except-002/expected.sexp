
(source_file (module (header_line) name: (identifier) (header_line)
  (operator_definition name: (identifier) (def_eq)
    definition: (except 
      expr_to_update: (identifier_ref) 
      (except_update
        update_specifier: (except_update_specifier
          (except_update_fn_appl (nat_number))
          (except_update_fn_appl (string))
          (except_update_fn_appl (nat_number))
        )
        new_val: (bound_infix_op lhs: (prev_func_val) symbol: (plus) rhs: (nat_number))
      )
    )
  )
(double_line)))

