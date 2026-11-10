
(source_file (module (header_line) name: (identifier) (header_line)
  (theorem statement: (boolean) proof: (non_terminal_proof
    (proof_step (proof_step_id (level) (name)) (suffices_proof_step
      (conj_list
        (conj_item (bullet_conj) (nat_number))
        (conj_item (bullet_conj) (nat_number))
      )
      (non_terminal_proof
        (qed_step (proof_step_id (level) (name)) (terminal_proof (use_body (use_body_expr
          (disj_list
            (disj_item (bullet_disj) (nat_number))
            (disj_item (bullet_disj) (nat_number))
          )
        ))))
      )
    ))
    (qed_step (proof_step_id (level) (name)))
  ))
(double_line)))

