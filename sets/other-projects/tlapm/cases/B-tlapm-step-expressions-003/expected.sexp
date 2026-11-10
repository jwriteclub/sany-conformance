
(source_file (module (header_line) name: (identifier) (header_line)
   (extends (identifier_ref))
   (operator_definition name: (identifier) (def_eq) definition:
    (step_expr_or_stutter
      (identifier_ref)
      (prefixed_op
        prefix: (subexpr_prefix (subexpr_component (identifier_ref)))
        op: (bound_op name: (identifier_ref) parameter: (identifier_ref))
      )
    )
  )
(double_line)))

