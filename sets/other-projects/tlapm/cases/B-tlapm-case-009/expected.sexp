
(source_file (module (header_line) name: (identifier) (header_line)
  (extends (identifier_ref))
  (operator_definition name: (identifier) (def_eq) definition:
    (case
      (case_arm (identifier_ref) (case_arrow) (conj_list (conj_item (bullet_conj)
        (case
          (case_arm (identifier_ref) (case_arrow) (identifier_ref))
          (case_box)
          (case_arm (identifier_ref) (case_arrow) (identifier_ref))
        )
      )))
      (case_box)
      (case_arm (identifier_ref) (case_arrow) (identifier_ref))
    )
  )
  (operator_definition name: (identifier) (def_eq) definition:
    (case
      (case_arm (identifier_ref) (case_arrow) (conj_list (conj_item (bullet_conj)
        (case
          (case_arm (identifier_ref) (case_arrow) (identifier_ref))
          (case_box)
          (other_arm (case_arrow) (identifier_ref))
        )
      )))
      (case_box)
      (case_arm (identifier_ref) (case_arrow) (identifier_ref))
    )
  )
  (operator_definition name: (identifier) (def_eq) definition:
    (case
      (case_arm (identifier_ref) (case_arrow) (conj_list
        (conj_item (bullet_conj)
          (case
            (case_arm (identifier_ref) (case_arrow) (identifier_ref))
            (case_box)
            (case_arm (identifier_ref) (case_arrow) (identifier_ref))
          )
        )
        (conj_item (bullet_conj) (identifier_ref))
      ))
      (case_box)
      (case_arm (identifier_ref) (case_arrow) (identifier_ref))
    )
  )
(double_line)))
