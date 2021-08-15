package repository

import (
	"sis/data"
	"sis/models"
)

type MatterUser struct {
	Data *data.Data
}

func (mu *MatterUser) GetAll() ([]models.MatterUser, error) {
	q := `select user_id,matter_id, cal1, cal2, cal3, cal4 from matter_user;`

	rows, err := mu.Data.DB.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var mattersUsers []models.MatterUser
	for rows.Next() {
		var matterUser models.MatterUser
		rows.Scan(&matterUser.UserId, &matterUser.MatterId, &matterUser.Cal1, &matterUser.Cal2, &matterUser.Cal3, &matterUser.Cal4)
		mattersUsers = append(mattersUsers, matterUser)

	}
	return mattersUsers, nil
}

func (mu *MatterUser) GetById(user_id uint, matter_id uint) (models.MatterUser, error) {
	q := `select user_id,matter_id, cal1, cal2, cal3, cal4 from matter_user where user_id=$1 and matter_id=$2;`

	row := mu.Data.DB.QueryRow(q, user_id, matter_id)

	var matterUser models.MatterUser

	err := row.Scan(&matterUser.UserId, &matterUser.MatterId, &matterUser.Cal1, &matterUser.Cal2, &matterUser.Cal3, &matterUser.Cal4)
	if err != nil {
		return models.MatterUser{}, err
	}

	return matterUser, nil
}

func (mu *MatterUser) Create(matterUser models.MatterUser) error {
	q := `insert into matterUser (user_id,matter_id,cal1, cal2, cal3, cal4) values ($1,$2,$3,$4,$5,$6);`

	stmt, err := mu.Data.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(matterUser.UserId, matterUser.MatterId, matterUser.Cal1, matterUser.Cal2, matterUser.Cal3, matterUser.Cal4)
	if err != nil {
		return err
	}

	return nil
}

//
////
//func (mu *MatterUser) Update(matter models.Matter, id uint) error {
//	q := `update matter set user_id =$1,course_id=$2 where id=$3;`
//
//	stmt, err := m.Data.DB.Prepare(q)
//
//	if err != nil {
//		return err
//	}
//
//	defer stmt.Close()
//
//	_, err = stmt.Exec(matter.UserId, matter.CourseId, id)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (mu *MatterUser) Delete(id uint) error {
//	q := `delete from matter where id=$1;`
//
//	stmt, err := m.Data.DB.Prepare(q)
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
