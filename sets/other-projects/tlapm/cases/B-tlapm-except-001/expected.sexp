
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq)
    definition:
        (except
            expr_to_update: (identifier_ref)
            (except_update
              update_specifier:
                (except_update_specifier
                    (except_update_fn_appl (nat_number)))
              new_val: (string))
        ))
(double_line)))

