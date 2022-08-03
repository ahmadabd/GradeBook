package grades

import (
	"fmt"
	"sync"
)

type Student struct {
	ID	int
	FirstName	string
	LastName	string
	Grades []Grade
}

func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}

	return result / float32(len(s.Grades))
}

type Students []Student

func (s Students) GetByID(id int) (*Student, error) {
	for _, student := range s {
		if student.ID == id {
			return &student, nil
		}
	}

	return nil, fmt.Errorf("Student with ID %d not found", id)
}

var (
	students Students
	studentMutex sync.Mutex
)

type GradeType string

const (
	GradeTest = GradeType("Test")
	GradeQuiz = GradeType("Quiz")
	GradeHomework = GradeType("Homework")
)

type Grade struct {
	Title string
	Type GradeType
	Score float32
}