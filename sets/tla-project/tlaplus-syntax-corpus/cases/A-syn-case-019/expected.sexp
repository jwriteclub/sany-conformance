
(source_file
  (module
    (header_line)
    (identifier)
    (header_line)
    (operator_definition
      (identifier)
      (def_eq)
      (case
        (case_arm
          (identifier_ref)
          (case_arrow)
          (disj_list
            (disj_item
              (bullet_disj)
              (case
                (case_arm (identifier_ref) (case_arrow) (identifier_ref))
                (case_box)
                (case_arm (identifier_ref) (case_arrow) (identifier_ref))
              )
            )
          )
        )
        (case_box)
        (other_arm (case_arrow) (identifier_ref))
      )
    )
    (double_line)
  )
)

