
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq) definition:
    (bound_op name: (identifier_ref)
      parameter: (infix_op_symbol (land))
      parameter: (identifier_ref)
      parameter: (infix_op_symbol (land))
    )
  )
  (operator_definition name: (identifier) (def_eq) definition:
    (bound_op name: (identifier_ref)
      parameter: (infix_op_symbol (lor))
      parameter: (identifier_ref)
      parameter: (infix_op_symbol (lor))
    )
  )
(double_line)))

