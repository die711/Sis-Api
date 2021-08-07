package repository

import (
	"sis/data"
	"sis/models"
)

type Career struct {
	Data *data.Data
}

func (c *Career) GetAll() ([]models.Career, error) {
	q := `select id, name from career where status = true;`

	rows, err := c.Data.DB.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var careers []models.Career
	for rows.Next() {
		var career models.Career
		rows.Scan(&career.Id, &career.Name)
		careers = append(careers, career)
	}

	return careers, nil
}

func (c *Career) GetById(id uint) (models.Career, error) {
	q := `select id, name from career where id = $1 and status = true;`

	row := c.Data.DB.QueryRow(q, id)

	var career models.Career

	err := row.Scan(&career.Id, &career.Name)
	if err != nil {
		return models.Career{}, err
	}

	return career, nil
}

func (c *Career) Create(career models.Career) error {
	q := `insert into career (name,status) values ($1,true);`

	stmt, err := c.Data.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(career.Name)
	if err != nil {
		return err
	}

	return nil
}

func (c *Career) Update(career models.Career, id uint) error {
	q := `update career set name=$1 where id=$2;`

	stmt, err := c.Data.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(career.Name, id)
	if err != nil {
		return err
	}

	return nil
}

func (c *Career) Delete(id uint) error {
	q := `update career set status=false where id=$1;`

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
