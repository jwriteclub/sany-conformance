
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (theorem statement: (boolean) proof:
    (non_terminal_proof
      (proof_step (proof_step_id (level) (name)) (pick_proof_step
        (identifier)
        (boolean)
      ))
      (proof_step (proof_step_id (level) (name)) (pick_proof_step
        (identifier) (identifier) (identifier)
        (boolean)
      ))
      (proof_step (proof_step_id (level) (name)) (pick_proof_step
        (quantifier_bound intro: (identifier) (set_in) set: (identifier_ref))
        (boolean)
      ))
      (proof_step (proof_step_id (level) (name)) (pick_proof_step
        (quantifier_bound
          intro: (identifier)
          intro: (identifier)
          intro: (identifier)
          (set_in)
          set: (identifier_ref)
        )
        (boolean)
      ))
      (proof_step (proof_step_id (level) (name)) (pick_proof_step
        (quantifier_bound intro: (identifier) (set_in) set: (identifier_ref))
        (quantifier_bound intro: (identifier) (set_in) set: (identifier_ref))
        (quantifier_bound intro: (identifier) (set_in) set: (identifier_ref))
        (boolean)
      ))
      (proof_step (proof_step_id (level) (name)) (pick_proof_step
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
        (boolean)
      ))
      (qed_step (proof_step_id (level) (name)))
    )
  )
(double_line)))

