
(source_file (module (header_line) name: (identifier) (header_line)
  (theorem statement: (boolean)
    proof: (non_terminal_proof
      (qed_step (proof_step_id (level) (name))
        (terminal_proof (use_body (use_body_expr
          (subexpression
            (subexpr_prefix
              (proof_step_ref (level) (name))
              (subexpr_tree_nav (langle_bracket))
              (subexpr_tree_nav (rangle_bracket))
              (subexpr_tree_nav (child_id))
              (subexpr_tree_nav (operator_args (identifier_ref) (identifier_ref)))
              (subexpr_tree_nav (colon))
            )
            (subexpr_tree_nav (address))
          )
          (subexpression
            (subexpr_prefix
              (proof_step_ref (level) (name))
              (subexpr_tree_nav (langle_bracket))
              (subexpr_tree_nav (rangle_bracket))
              (subexpr_tree_nav (child_id))
              (subexpr_tree_nav (operator_args (identifier_ref) (identifier_ref)))
              (subexpr_tree_nav (colon))
            )
            (subexpr_tree_nav (address))
          )
        )))
      )
    )
  )
(double_line)))

