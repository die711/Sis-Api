package repository

import (
	"sis/data"
	"sis/models"
)

type Type struct {
	Data *data.Data
}

func (t *Type) GetAll() ([]models.Type, error) {
	q := `select id, name from type where status = true;`

	rows, err := t.Data.DB.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var types []models.Type
	for rows.Next() {
		var type_ models.Type
		rows.Scan(&type_.Id, &type_.Name)
		types = append(types, type_)
	}

	return types, nil
}

func (t *Type) GetById(id uint) (models.Type, error) {
	q := `select id, name from type where id = $1 and status = true;`

	row := t.Data.DB.QueryRow(q, id)

	var type_ models.Type

	err := row.Scan(&type_.Id, &type_.Name)
	if err != nil {
		return models.Type{}, err
	}

	return type_, nil
}

func (t *Type) Create(type_ models.Type) error {
	q := `insert into type (name,status) values ($1,true);`

	stmt, err := t.Data.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(type_.Name)
	if err != nil {
		return err
	}

	return nil
}

func (t *Type) Update(type_ models.Type, id uint) error {
	q := `update type set name=$1 where id=$2;`

	stmt, err := t.Data.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(type_.Name, id)
	if err != nil {
		return err
	}

	return nil
}

func (t *Type) Delete(id uint) error {
	q := `update type set status=false where id=$1;`

	stmt, err := t.Data.DB.Prepare(q)

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
