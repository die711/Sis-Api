package repository

import (
	"sis/data"
	"sis/models"
)

type Matter struct {
	Data *data.Data
}

func (m *Matter) GetAll() ([]models.Matter, error) {
	q := `select id,user_id, course_id from matter;`

	rows, err := m.Data.DB.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var matters []models.Matter
	for rows.Next() {
		var matter models.Matter
		rows.Scan(&matter.Id, &matter.UserId, &matter.CourseId)
		matters = append(matters, matter)
	}

	return matters, nil
}

func (m *Matter) GetById(id uint) (models.Matter, error) {
	q := `select id, user_id, course_id from matter where id = $1`

	row := m.Data.DB.QueryRow(q, id)

	var matter models.Matter

	err := row.Scan(&matter.Id, &matter.UserId, &matter.CourseId)
	if err != nil {
		return models.Matter{}, err
	}

	return matter, nil
}

//
func (m *Matter) Create(matter models.Matter) error {
	q := `insert into matter (user_id,course_id) values ($1,$2);`

	stmt, err := m.Data.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(matter.UserId, matter.CourseId)
	if err != nil {
		return err
	}

	return nil
}

//
func (m *Matter) Update(matter models.Matter, id uint) error {
	q := `update matter set user_id =$1,course_id=$2 where id=$3;`

	stmt, err := m.Data.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(matter.UserId, matter.CourseId, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *Matter) Delete(id uint) error {
	q := `delete from matter where id=$1;`

	stmt, err := m.Data.DB.Prepare(q)

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
