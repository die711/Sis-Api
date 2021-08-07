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

	err := row.Scan(&course.Id,&course.CareerId, &course.Name,&course.Credits)
	if err != nil {
		return models.Course{}, err
	}

	return course, nil
}
//
//func (c *Career) Create(career models.Career) error {
//	q := `insert into career (name,status) values ($1,true);`
//
//	stmt, err := c.Data.DB.Prepare(q)
//
//	if err != nil {
//		return err
//	}
//
//	defer stmt.Close()
//
//	_, err = stmt.Exec(career.Name)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (c *Career) Update(career models.Career, id uint) error {
//	q := `update career set name=$1 where id=$2;`
//
//	stmt, err := c.Data.DB.Prepare(q)
//
//	if err != nil {
//		return err
//	}
//
//	defer stmt.Close()
//
//	_, err = stmt.Exec(career.Name, id)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (c *Career) Delete(id uint) error {
//	q := `update career set status=false where id=$1;`
//
//	stmt, err := c.Data.DB.Prepare(q)
//
//	if err != nil {
//		return err
//	}
//
//	defer stmt.Close()
//	_, err = stmt.Exec(id)
//	if err != nil {
//		return err
//	}
//	return nil
//}
