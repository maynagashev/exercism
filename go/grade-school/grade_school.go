package school

import (
	"sort"
)

type School struct {
	grades map[int]*Grade
}

type Grade struct {
	level    int
	students []string
}

func New() *School {
	return &School{grades: make(map[int]*Grade)}
}

func (s *School) Add(student string, grade int) {
	gradePtr, ok := s.grades[grade]

	// if no grade in school, create new grade and store pointer
	if !ok {
		s.grades[grade] = &Grade{level: grade}
		gradePtr = s.grades[grade]
	}

	// check if not exists
	for _, v := range gradePtr.students {
		if student == v {
			// exists, skipping
			return
		}
	}

	gradePtr.students = append(gradePtr.students, student)

	sort.Strings(gradePtr.students)

	//fmt.Printf("students: %v \ns.grades: %v\n", gradePtr.students, s.Grade(grade))
}

func (s *School) Grade(level int) []string {
	gradePtr, ok := s.grades[level]
	if !ok {
		return []string{}
	}
	return gradePtr.students
}

func (s *School) Enrollment() []Grade {
	var res []Grade
	keys := make([]int, 0, len(s.grades))
	for k, _ := range s.grades {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	//fmt.Printf("keys: %v\n len grades: %v\n", keys, len(s.grades))
	for _, k := range keys {
		res = append(res, *s.grades[k])
	}
	return res
}
