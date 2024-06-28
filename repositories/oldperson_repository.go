package repositories

import (
	"ISHC/config"
	"ISHC/models"
	"database/sql"
)

func CreateOldPerson(oldPerson *models.OldPersonInfo) error {
	query := `INSERT INTO oldperson_info (username, gender, phone, id_card, birthday, checkin_date, checkout_date, imgset_dir, profile_photo, room_number, firstguardian_name, firstguardian_relationship, firstguardian_phone, firstguardian_wechat, secondguardian_name, secondguardian_relationship, secondguardian_phone, secondguardian_wechat, health_state, description, isactive, created, createby, updated, updateby, remove) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := config.DB.Exec(query, oldPerson.Username, oldPerson.Gender, oldPerson.Phone, oldPerson.IDCard, oldPerson.Birthday, oldPerson.CheckinDate, oldPerson.CheckoutDate, oldPerson.ImgsetDir, oldPerson.ProfilePhoto, oldPerson.RoomNumber, oldPerson.FirstGuardianName, oldPerson.FirstGuardianRelationship, oldPerson.FirstGuardianPhone, oldPerson.FirstGuardianWechat, oldPerson.SecondGuardianName, oldPerson.SecondGuardianRelationship, oldPerson.SecondGuardianPhone, oldPerson.SecondGuardianWechat, oldPerson.HealthState, oldPerson.Description, oldPerson.IsActive, oldPerson.Created, oldPerson.CreatedBy, oldPerson.Updated, oldPerson.UpdatedBy, oldPerson.Remove)
	return err
}

func UpdateOldPerson(oldPerson *models.OldPersonInfo) error {
	query := `UPDATE oldperson_info SET username=?, gender=?, phone=?, id_card=?, birthday=?, checkin_date=?, checkout_date=?, imgset_dir=?, profile_photo=?, room_number=?, firstguardian_name=?, firstguardian_relationship=?, firstguardian_phone=?, firstguardian_wechat=?, secondguardian_name=?, secondguardian_relationship=?, secondguardian_phone=?, secondguardian_wechat=?, health_state=?, description=?, isactive=?, updated=?, updateby=?, remove=? 
              WHERE id=?`
	_, err := config.DB.Exec(query, oldPerson.Username, oldPerson.Gender, oldPerson.Phone, oldPerson.IDCard, oldPerson.Birthday, oldPerson.CheckinDate, oldPerson.CheckoutDate, oldPerson.ImgsetDir, oldPerson.ProfilePhoto, oldPerson.RoomNumber, oldPerson.FirstGuardianName, oldPerson.FirstGuardianRelationship, oldPerson.FirstGuardianPhone, oldPerson.FirstGuardianWechat, oldPerson.SecondGuardianName, oldPerson.SecondGuardianRelationship, oldPerson.SecondGuardianPhone, oldPerson.SecondGuardianWechat, oldPerson.HealthState, oldPerson.Description, oldPerson.IsActive, oldPerson.Updated, oldPerson.UpdatedBy, oldPerson.Remove, oldPerson.ID)
	return err
}

func GetOldPersonById(id int) (*models.OldPersonInfo, error) {
	query := `SELECT id, username, gender, phone, id_card, birthday, checkin_date, checkout_date, imgset_dir, profile_photo, room_number, firstguardian_name, firstguardian_relationship, firstguardian_phone, firstguardian_wechat, secondguardian_name, secondguardian_relationship, secondguardian_phone, secondguardian_wechat, health_state, description, isactive, created, createby, updated, updateby, remove 
              FROM oldperson_info WHERE id=?`
	row := config.DB.QueryRow(query, id)

	var oldPerson models.OldPersonInfo
	err := row.Scan(&oldPerson.ID, &oldPerson.Username, &oldPerson.Gender, &oldPerson.Phone, &oldPerson.IDCard, &oldPerson.Birthday, &oldPerson.CheckinDate, &oldPerson.CheckoutDate, &oldPerson.ImgsetDir, &oldPerson.ProfilePhoto, &oldPerson.RoomNumber, &oldPerson.FirstGuardianName, &oldPerson.FirstGuardianRelationship, &oldPerson.FirstGuardianPhone, &oldPerson.FirstGuardianWechat, &oldPerson.SecondGuardianName, &oldPerson.SecondGuardianRelationship, &oldPerson.SecondGuardianPhone, &oldPerson.SecondGuardianWechat, &oldPerson.HealthState, &oldPerson.Description, &oldPerson.IsActive, &oldPerson.Created, &oldPerson.CreatedBy, &oldPerson.Updated, &oldPerson.UpdatedBy, &oldPerson.Remove)
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
