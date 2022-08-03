package grades

func init() {
	students = []Student{
		Student{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Grades: []Grade{
				Grade{
					Title: "Test 1",
					Type:  GradeTest,
					Score: 10,
				},
				Grade{
					Title: "Test 2",
					Type:  GradeQuiz,
					Score: 14,
				},
				Grade{
					Title: "Test 3",
					Type:  GradeHomework,
					Score: 8,
				},
			},
		},
		Student{
			ID:        2,
			FirstName: "Jane",
			LastName:  "Doe",
			Grades: []Grade{
				Grade{
					Title: "Test 1",
					Type:  GradeTest,
					Score: 20,
				},
				Grade{
					Title: "Test 2",
					Type:  GradeQuiz,
					Score: 19,
				},
				Grade{
					Title: "Test 3",
					Type:  GradeHomework,
					Score: 18,
				},
			},
		},
	}
}
