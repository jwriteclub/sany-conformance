
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (theorem statement: (bound_infix_op
    lhs: (identifier_ref)
    symbol: (implies)
    rhs: (identifier_ref)
  ) proof:
    (non_terminal_proof
      (proof_step (proof_step_id (level) (name)) (suffices_proof_step (identifier_ref)))
      (proof_step (proof_step_id (level) (name)) (suffices_proof_step (identifier_ref)
        (terminal_proof (use_body (use_body_expr (identifier_ref))))
      ))
      (qed_step (proof_step_id (level) (name)))
    )
  )
(double_line)))

