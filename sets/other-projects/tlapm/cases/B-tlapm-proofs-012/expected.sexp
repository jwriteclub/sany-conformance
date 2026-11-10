
(source_file (module (header_line) name: (identifier) (header_line)
  (theorem statement: (boolean) proof:
    (non_terminal_proof
      (proof_step (proof_step_id (level) (name)) (suffices_proof_step (boolean)))
      (proof_step (proof_step_id (level) (name)) (definition_proof_step (operator_definition name: (identifier) (def_eq) definition: (nat_number))))
      (proof_step (proof_step_id (level) (name)) (have_proof_step (nat_number)))
      (proof_step (proof_step_id (level) (name)) (witness_proof_step (nat_number) (nat_number)))
      (proof_step (proof_step_id (level) (name)) (take_proof_step (identifier) (identifier) (identifier)))
      (proof_step (proof_step_id (level) (name)) (suffices_proof_step (proof_step_ref (level) (name))))
      (proof_step (proof_step_id (level) (name)) (suffices_proof_step (nat_number)))
      (proof_step (proof_step_id (level) (name)) (case_proof_step (nat_number)))
      (proof_step (proof_step_id (level) (name)) (pick_proof_step (identifier) (identifier) (identifier) (nat_number)))
      (proof_step (proof_step_id (level) (name)) (use_or_hide (use_body (use_body_expr (nat_number)))))
      (proof_step (proof_step_id (level) (name)) (use_or_hide (use_body (use_body_expr (nat_number)))))
      (qed_step (proof_step_id (level) (name)))
    )
  )
(double_line)))

