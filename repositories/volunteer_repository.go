package repositories

import (
	"ISHC/config"
	"ISHC/models"
	"database/sql"
	"fmt"
)

func CreateVolunteer(volunteer *models.VolunteerInfo) error {
	query := `INSERT INTO volunteer_info (
		org_id, client_id, name, gender, phone, id_card, birthday, checkin_date, checkout_date, 
		imgset_dir, profile_photo, description, isactive, created, createby, updated, updateby, remove
	) VALUES (1, 1, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := config.DB.Exec(query,
		volunteer.Name,
		volunteer.Gender,
		volunteer.Phone,
		volunteer.IDCard,
		volunteer.Birthday.Time.Format(models.CtLayoutDateTime),
		volunteer.CheckinDate.Time.Format(models.CtLayoutDateTime),
		volunteer.CheckoutDate.Time.Format(models.CtLayoutDateTime),
		volunteer.ImgsetDir,
		volunteer.ProfilePhoto,
		volunteer.Description,
		volunteer.IsActive,
		volunteer.Created.Time.Format(models.CtLayoutDateTime),
		volunteer.CreatedBy,
		volunteer.Updated.Time.Format(models.CtLayoutDateTime),
		volunteer.UpdatedBy,
		volunteer.Remove)
	return err
}

func UpdateVolunteer(volunteer *models.VolunteerInfo) error {
	query := `UPDATE volunteer_info SET 
		name=?, gender=?, phone=?, id_card=?, birthday=?, checkin_date=?, checkout_date=?, 
		imgset_dir=?, profile_photo=?, description=?, isactive=?, updated=?, updateby=?, remove=? 
		WHERE id=?`
	_, err := config.DB.Exec(query,
		volunteer.Name,
		volunteer.Gender,
		volunteer.Phone,
		volunteer.IDCard,
		volunteer.Birthday.Time.Format(models.CtLayoutDateTime),
		volunteer.CheckinDate.Time.Format(models.CtLayoutDateTime),
		volunteer.CheckoutDate.Time.Format(models.CtLayoutDateTime),
		volunteer.ImgsetDir,
		volunteer.ProfilePhoto,
		volunteer.Description,
		volunteer.IsActive,
		volunteer.Updated.Time.Format(models.CtLayoutDateTime),
		volunteer.UpdatedBy,
		volunteer.Remove,
		volunteer.ID)
	return err
}

func GetAllVolunteers() ([]models.VolunteerInfo, error) {
	query := `SELECT id, org_id, client_id, name, gender, phone, id_card, birthday, checkin_date, checkout_date, imgset_dir, profile_photo, description, isactive, created, createby, updated, updateby, remove 
              FROM volunteer_info`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying volunteer_info: %v", err)
	}
	defer rows.Close()

	var volunteers []models.VolunteerInfo
	for rows.Next() {
		var volunteer models.VolunteerInfo
		err := rows.Scan(
			&volunteer.ID,
			&volunteer.OrgID,
			&volunteer.ClientID,
			&volunteer.Name,
			&volunteer.Gender,
			&volunteer.Phone,
			&volunteer.IDCard,
			&volunteer.Birthday,
			&volunteer.CheckinDate,
			&volunteer.CheckoutDate,
			&volunteer.ImgsetDir,
			&volunteer.ProfilePhoto,
			&volunteer.Description,
			&volunteer.IsActive,
			&volunteer.Created,
			&volunteer.CreatedBy,
			&volunteer.Updated,
			&volunteer.UpdatedBy,
			&volunteer.Remove)
		if err != nil {
			return nil, fmt.Errorf("error scanning volunteer_info row: %v", err)
		}
		volunteers = append(volunteers, volunteer)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over volunteer_info rows: %v", err)
	}

	return volunteers, nil
}

func GetVolunteerById(id int) (*models.VolunteerInfo, error) {
	query := `SELECT 
		id, name, gender, phone, id_card, birthday, checkin_date, checkout_date, 
		imgset_dir, profile_photo, description, isactive, created, createby, updated, updateby, remove 
		FROM volunteer_info WHERE id=?`
	row := config.DB.QueryRow(query, id)

	var volunteer models.VolunteerInfo
	err := row.Scan(
		&volunteer.ID, &volunteer.Name, &volunteer.Gender, &volunteer.Phone, &volunteer.IDCard,
		&volunteer.Birthday, &volunteer.CheckinDate, &volunteer.CheckoutDate, &volunteer.ImgsetDir,
		&volunteer.ProfilePhoto, &volunteer.Description, &volunteer.IsActive, &volunteer.Created,
		&volunteer.CreatedBy, &volunteer.Updated, &volunteer.UpdatedBy, &volunteer.Remove)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &volunteer, nil
}

func DeleteVolunteer(id int) error {
	query := `DELETE FROM volunteer_info WHERE id=?`
	_, err := config.DB.Exec(query, id)
	return err
}

func SetVolunteerProfilePhoto(id int, profilePhoto string) error {
	query := `UPDATE volunteer_info SET profile_photo=? WHERE id=?`
	_, err := config.DB.Exec(query, profilePhoto, id)
	return err
}
