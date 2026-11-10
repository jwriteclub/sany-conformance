
(source_file (module (header_line) name: (identifier) (header_line)
  (theorem statement: (boolean) proof:
    (non_terminal_proof
      (proof_step (proof_step_id (level) (name)) (suffices_proof_step (boolean)))
      (qed_step (proof_step_id (level) (name))
        (terminal_proof (use_body (use_body_expr
          (proof_step_ref (level) (name))
        )))
      )
    )
  )
(double_line)))

