package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Reza",
			LastName:  "Moradi",
			Grades: []Grade{
				{
					Title: "Test 1",
					Type:  GradeTest,
					Score: 10,
				},
				{
					Title: "Test 2",
					Type:  GradeQuiz,
					Score: 14,
				},
				{
					Title: "Test 3",
					Type:  GradeHomework,
					Score: 8,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Sara",
			LastName:  "Naghi",
			Grades: []Grade{
				{
					Title: "Test 1",
					Type:  GradeTest,
					Score: 20,
				},
				{
					Title: "Test 2",
					Type:  GradeQuiz,
					Score: 19,
				},
				{
					Title: "Test 3",
					Type:  GradeHomework,
					Score: 18,
				},
			},
		},
	}
}
