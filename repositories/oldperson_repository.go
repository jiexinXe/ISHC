package repositories

import (
	"ISHC/config"
	"ISHC/models"
	"database/sql"
	"fmt"
)

func CreateOldPerson(oldPerson *models.OldPersonInfo) error {
	query := `INSERT INTO oldperson_info (
        username, gender, phone, id_card,
        birthday, checkin_date, checkout_date, imgset_dir,
        profile_photo, room_number, firstguardian_name,
        firstguardian_relationship, firstguardian_phone,
        firstguardian_wechat, health_state, description,
        isactive, created, createby, updated, updateby, remove
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// 执行SQL查询
	_, err := config.DB.Exec(query,
		oldPerson.Username,
		oldPerson.Gender,
		oldPerson.Phone,
		oldPerson.IDCard,
		oldPerson.Birthday.Time.Format(models.CtLayoutDate),
		oldPerson.CheckinDate.Time.Format(models.CtLayoutDate),
		oldPerson.CheckoutDate.Time.Format(models.CtLayoutDate),
		oldPerson.ImgsetDir,
		oldPerson.ProfilePhoto,
		oldPerson.RoomNumber,
		oldPerson.FirstGuardianName,
		oldPerson.FirstGuardianRelationship,
		oldPerson.FirstGuardianPhone,
		oldPerson.FirstGuardianWechat,
		oldPerson.HealthState,
		oldPerson.Description,
		oldPerson.IsActive,
		oldPerson.Created.Time.Format(models.CtLayoutDate),
		oldPerson.CreatedBy,
		oldPerson.Updated.Time.Format(models.CtLayoutDate),
		oldPerson.UpdatedBy,
		oldPerson.Remove)

	if err != nil {
		fmt.Println("Error executing query: ", err)
	}
	return err
}

func UpdateOldPerson(oldPerson *models.OldPersonInfo) error {
	query := `UPDATE oldperson_info SET username=?, gender=?, phone=?, id_card=?, birthday=?, checkin_date=?, checkout_date=?, imgset_dir=?, profile_photo=?, room_number=?, firstguardian_name=?, firstguardian_relationship=?, firstguardian_phone=?, firstguardian_wechat=?, health_state=?, description=?, isactive=?, updated=?, updateby=?, remove=? 
              WHERE id=?`
	_, err := config.DB.Exec(query, oldPerson.Username, oldPerson.Gender, oldPerson.Phone, oldPerson.IDCard, oldPerson.Birthday, oldPerson.CheckinDate, oldPerson.CheckoutDate, oldPerson.ImgsetDir, oldPerson.ProfilePhoto, oldPerson.RoomNumber, oldPerson.FirstGuardianName, oldPerson.FirstGuardianRelationship, oldPerson.FirstGuardianPhone, oldPerson.FirstGuardianWechat, oldPerson.HealthState, oldPerson.Description, oldPerson.IsActive, oldPerson.Updated, oldPerson.UpdatedBy, oldPerson.Remove, oldPerson.ID)
	return err
}

func GetAllOldPersons() ([]models.OldPersonInfo, error) {
	query := `SELECT id, username, gender, phone, id_card, birthday, checkin_date, checkout_date, imgset_dir, profile_photo, room_number, firstguardian_name, firstguardian_relationship, firstguardian_phone, firstguardian_wechat, health_state, description, isactive, created, createby, updated, updateby, remove 
              FROM oldperson_info`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying oldperson_info: %v", err)
	}
	defer rows.Close()

	var oldPersons []models.OldPersonInfo
	for rows.Next() {
		var oldPerson models.OldPersonInfo
		err := rows.Scan(&oldPerson.ID, &oldPerson.Username, &oldPerson.Gender, &oldPerson.Phone, &oldPerson.IDCard, &oldPerson.Birthday, &oldPerson.CheckinDate, &oldPerson.CheckoutDate, &oldPerson.ImgsetDir, &oldPerson.ProfilePhoto, &oldPerson.RoomNumber, &oldPerson.FirstGuardianName, &oldPerson.FirstGuardianRelationship, &oldPerson.FirstGuardianPhone, &oldPerson.FirstGuardianWechat, &oldPerson.HealthState, &oldPerson.Description, &oldPerson.IsActive, &oldPerson.Created, &oldPerson.CreatedBy, &oldPerson.Updated, &oldPerson.UpdatedBy, &oldPerson.Remove)
		if err != nil {
			return nil, fmt.Errorf("error scanning oldperson_info row: %v", err)
		}
		oldPersons = append(oldPersons, oldPerson)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over oldperson_info rows: %v", err)
	}

	return oldPersons, nil
}

func GetOldPersonById(id int) (*models.OldPersonInfo, error) {
	query := `SELECT id, username, gender, phone, id_card, birthday, checkin_date, checkout_date, imgset_dir, profile_photo, room_number, firstguardian_name, firstguardian_relationship, firstguardian_phone, firstguardian_wechat, health_state, description, isactive, created, createby, updated, updateby, remove 
              FROM oldperson_info WHERE id=?`
	row := config.DB.QueryRow(query, id)

	var oldPerson models.OldPersonInfo
	err := row.Scan(&oldPerson.ID, &oldPerson.Username, &oldPerson.Gender, &oldPerson.Phone, &oldPerson.IDCard, &oldPerson.Birthday, &oldPerson.CheckinDate, &oldPerson.CheckoutDate, &oldPerson.ImgsetDir, &oldPerson.ProfilePhoto, &oldPerson.RoomNumber, &oldPerson.FirstGuardianName, &oldPerson.FirstGuardianRelationship, &oldPerson.FirstGuardianPhone, &oldPerson.FirstGuardianWechat, &oldPerson.HealthState, &oldPerson.Description, &oldPerson.IsActive, &oldPerson.Created, &oldPerson.CreatedBy, &oldPerson.Updated, &oldPerson.UpdatedBy, &oldPerson.Remove)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &oldPerson, nil
}

func DeleteOldPerson(id int) error {
	query := `DELETE FROM oldperson_info WHERE id=?`
	_, err := config.DB.Exec(query, id)
	return err
}

func SetOldPersonProfilePhoto(id int, profilePhoto string) error {
	query := `UPDATE oldperson_info SET profile_photo=? WHERE id=?`
	_, err := config.DB.Exec(query, profilePhoto, id)
	return err
}
