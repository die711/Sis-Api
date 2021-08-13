package repository

import (
	"sis/data"
	"sis/models"
)

type User struct {
	Data *data.Data
}

func (u *User) GetAll() ([]models.User, error) {
	q := `select id, name, last_name, email, password, type_id, career_id from "user" where status = true;`

	rows, err := u.Data.DB.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.Password, &user.TypeId, &user.CareerId)
		users = append(users, user)
	}

	return users, nil
}

func (u *User) GetById(id uint) (models.User, error) {
	q := `select id,name,last_name,email,password,type_id,career_id from "user" where id = $1 and status = true;`

	row := u.Data.DB.QueryRow(q, id)

	var user models.User

	err := row.Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.Password, &user.TypeId, &user.CareerId)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *User) Create(user models.User) error {
	q := `insert into "user" (name,last_name,email,password,type_id,career_id,status) values ($1,$2,$3,$4,$5,$6,true);`

	stmt, err := u.Data.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.LastName, user.Email, user.Password, user.TypeId, user.CareerId)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Update(user models.User, id uint) error {
	q := `update "user" set name =$1,last_name=$2, email=$3, password=$4, type_id=$5,career_id=$6  where id=$7;`

	stmt, err := u.Data.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.LastName, user.Email, user.Password, user.TypeId, user.CareerId, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Delete(id uint) error {
	q := `update "user" set status=false where id=$1;`

	stmt, err := u.Data.DB.Prepare(q)

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
