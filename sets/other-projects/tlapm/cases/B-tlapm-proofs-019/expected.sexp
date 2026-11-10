
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (theorem statement: (boolean) proof:
    (non_terminal_proof
      (proof_step (proof_step_id (level) (name)) (suffices_proof_step (nat_number) (non_terminal_proof
        (proof_step (proof_step_id (level) (name)) (suffices_proof_step (nat_number)))
        (qed_step (proof_step_id (level) (name)) (terminal_proof (use_body (use_body_expr (identifier_ref)))))
      )))
      (proof_step (proof_step_id (level) (name)) (suffices_proof_step (nat_number) (non_terminal_proof
        (proof_step (proof_step_id (level) (name)) (suffices_proof_step (nat_number)))
        (qed_step (proof_step_id (level) (name)) (terminal_proof))
      )))
      (proof_step (proof_step_id (level) (name)) (suffices_proof_step (nat_number) (non_terminal_proof
        (proof_step (proof_step_id (level) (name)) (suffices_proof_step (nat_number)))
        (qed_step (proof_step_id (level) (name)) (terminal_proof))
      )))
      (proof_step (proof_step_id (level) (name)) (suffices_proof_step (nat_number)))
      (qed_step (proof_step_id (level) (name)))
    )
  )
(double_line)))

