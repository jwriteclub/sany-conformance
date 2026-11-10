
(source_file (module (header_line) name: (identifier) (header_line)
  (recursive_declaration
    (operator_declaration name: (identifier) parameter: (placeholder))
    name: (identifier)
    (operator_declaration name: (identifier) parameter: (placeholder) parameter: (placeholder) parameter: (placeholder))
    (operator_declaration (placeholder) name: (infix_op_symbol (plus)) (placeholder))
    (operator_declaration name: (prefix_op_symbol (negative)) (placeholder))
    (operator_declaration (placeholder) name: (postfix_op_symbol (sup_plus)))
  )
  (operator_definition name: (identifier) parameter: (identifier) (def_eq)
   definition: (bound_op name: (identifier_ref) parameter: (identifier_ref)))
  (operator_definition name: (identifier) (def_eq) definition:
   (identifier_ref))
  (operator_definition name: (identifier) parameter: (identifier) parameter:
   (identifier) parameter: (identifier) (def_eq) definition:
   (bound_op name: (identifier_ref) parameter: (identifier_ref) parameter:
    (identifier_ref) parameter: (identifier_ref)))
  (operator_definition parameter: (identifier) name: (infix_op_symbol (plus))
   parameter: (identifier) (def_eq) definition:
   (bound_infix_op lhs: (identifier_ref) symbol: (plus) rhs: (identifier_ref)))
  (operator_definition name: (prefix_op_symbol (negative)) parameter:
   (identifier) (def_eq) definition:
   (bound_prefix_op symbol: (negative) rhs: (identifier_ref)))
  (operator_definition parameter: (identifier) name:
   (postfix_op_symbol (sup_plus)) (def_eq) definition:
   (bound_postfix_op lhs: (identifier_ref) symbol: (sup_plus)))
(double_line)))

