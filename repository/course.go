package repository

import (
	"sis/data"
	"sis/models"
)

type Course struct {
	Data *data.Data
}

func (c *Course) GetAll() ([]models.Course, error) {
	q := `select id,career_id,name, credits from course where status = true;`

	rows, err := c.Data.DB.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var course models.Course
		rows.Scan(&course.Id, &course.CareerId, &course.Name, &course.Credits)
		courses = append(courses, course)
	}

	return courses, nil
}

func (c *Course) GetById(id uint) (models.Course, error) {
	q := `select id, career_id, name, credits from course where id = $1 and status = true;`

	row := c.Data.DB.QueryRow(q, id)

	var course models.Course

	err := row.Scan(&course.Id, &course.CareerId, &course.Name, &course.Credits)
	if err != nil {
		return models.Course{}, err
	}

	return course, nil
}

//
func (c *Course) Create(course models.Course) error {
	q := `insert into course (career_id,name,credits,status) values ($1,$2,$3,true);`

	stmt, err := c.Data.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(course.CareerId, course.Name, course.Credits)
	if err != nil {
		return err
	}

	return nil
}

//
func (c *Course) Update(course models.Course, id uint) error {
	q := `update course set career_id =$1,name=$2, credits=$3 where id=$4;`

	stmt, err := c.Data.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(course.CareerId, course.Name, course.Credits, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *Course) Delete(id uint) error {
	q := `update course set status=false where id=$1;`

	stmt, err := c.Data.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
