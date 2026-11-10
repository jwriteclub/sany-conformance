
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq)
    definition:
        (except
            expr_to_update: (identifier_ref)
            (except_update
              update_specifier:
                (except_update_specifier
                    (except_update_fn_appl
                      (tuple_literal (langle_bracket)
                        (tuple_literal (langle_bracket) (nat_number) (nat_number) (rangle_bracket))
                        (nat_number)
                      (rangle_bracket))))
              new_val: (nat_number))
        ))
(double_line)))

