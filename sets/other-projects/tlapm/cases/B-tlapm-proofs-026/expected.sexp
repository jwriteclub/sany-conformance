
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (theorem statement: (boolean) proof:
    (non_terminal_proof
      (proof_step (proof_step_id (level) (name)) (take_proof_step
        (identifier)
      ))
      (proof_step (proof_step_id (level) (name)) (take_proof_step
        (identifier) (identifier) (identifier)
      ))
      (proof_step (proof_step_id (level) (name)) (take_proof_step
        (quantifier_bound intro: (identifier) (set_in) set: (identifier_ref))
      ))
      (proof_step (proof_step_id (level) (name)) (take_proof_step
        (quantifier_bound
          intro: (identifier)
          intro: (identifier)
          intro: (identifier)
          (set_in)
          set: (identifier_ref))
      ))
      (proof_step (proof_step_id (level) (name)) (take_proof_step
        (quantifier_bound intro: (identifier) (set_in) set: (identifier_ref))
        (quantifier_bound intro: (identifier) (set_in) set: (identifier_ref))
        (quantifier_bound intro: (identifier) (set_in) set: (identifier_ref))
      ))
      (proof_step (proof_step_id (level) (name)) (take_proof_step
        (quantifier_bound
          intro: (identifier)
          intro: (identifier)
          intro: (identifier)
          (set_in)
          set: (identifier_ref)
        )
        (quantifier_bound
          intro: (identifier)
          intro: (identifier)
          (set_in)
          set: (identifier_ref)
        )
        (quantifier_bound
          intro: (identifier)
          (set_in)
          set: (identifier_ref)
        )
      ))
      (qed_step (proof_step_id (level) (name)))
    )
  )
(double_line)))

