package repository

import (
	"sis/data"
	"sis/models"
)

type StudentRepository struct {
	Data *data.Data
}

func (sr *StudentRepository) GetAll() ([]models.Student, error) {
	q := `select * from student;`

	rows, err := sr.Data.DB.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		rows.Scan(&student.NoControl, &student.Name, &student.LastName, &student.Semester)
		students = append(students, student)
	}

	return students, nil
}

func (sr *StudentRepository) GetById() ([]models.Student, error) {
	q := `select * from student;`

	rows, err := sr.Data.DB.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		rows.Scan(&student.NoControl, &student.Name, &student.LastName, &student.Semester)
		students = append(students, student)
	}

	return students, nil
}

//func (sr *StudentRepository) Create() (models.Student, error) {
//	q := `select * from student;`
//
//	rows, err := sr.Data.DB.Query(q)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//
//	var students []models.Student
//	for rows.Next() {
//		var student models.Student
//		rows.Scan(&student.NoControl, &student.Name, &student.LastName, &student.Semester)
//		students = append(students, student)
//	}
//
//	return students, nil
//}
