
(source_file (module (header_line) name: (identifier) (header_line)
  (theorem statement: (boolean)
    proof: (non_terminal_proof
      (proof_step (proof_step_id (level) (name))
        (suffices_proof_step (assume_prove
          assumption: (new (statement_level) (identifier) (set_in) (identifier_ref))
          conclusion: (boolean)
        ))
      )
      (qed_step (proof_step_id (level) (name)))
    )
  )
(double_line)))

