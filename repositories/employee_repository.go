package repositories

import (
	"ISHC/config"
	"ISHC/models"
	"database/sql"
)

func CreateEmployee(employee *models.EmployeeInfo) error {
	query := `INSERT INTO employee_info (org_id, client_id, username, gender, phone, id_card, birthday, hire_date, resign_date, imgset_dir, profile_photo, description, isactive, created, createby, updated, updateby, remove) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := config.DB.Exec(query,
		employee.OrgID,
		employee.ClientID,
		employee.Username,
		employee.Gender,
		employee.Phone,
		employee.IDCard,
		employee.Birthday.Time.Format("2006-01-02"),
		employee.HireDate.Time.Format("2006-01-02"),
		employee.ResignDate.Time.Format("2006-01-02"),
		employee.ImgsetDir,
		employee.ProfilePhoto,
		employee.Description,
		employee.IsActive,
		employee.Created.Time.Format("2006-01-02"),
		employee.CreatedBy,
		employee.Updated.Time.Format("2006-01-02"),
		employee.UpdatedBy,
		employee.Remove)
	return err
}

func UpdateEmployee(employee *models.EmployeeInfo) error {
	query := `UPDATE employee_info SET username=?, gender=?, phone=?, id_card=?, birthday=?, hire_date=?, resign_date=?, imgset_dir=?, profile_photo=?, description=?, isactive=?, updated=?, updateby=?, remove=? 
              WHERE id=?`
	_, err := config.DB.Exec(query, employee.Username, employee.Gender, employee.Phone, employee.IDCard, employee.Birthday, employee.HireDate, employee.ResignDate, employee.ImgsetDir, employee.ProfilePhoto, employee.Description, employee.IsActive, employee.Updated, employee.UpdatedBy, employee.Remove, employee.ID)
	return err
}

func GetEmployeeById(id int) (*models.EmployeeInfo, error) {
	query := `SELECT id, org_id, client_id, username, gender, phone, id_card, birthday, hire_date, resign_date, imgset_dir, profile_photo, description, isactive, created, createby, updated, updateby, remove 
              FROM employee_info WHERE id=?`
	row := config.DB.QueryRow(query, id)

	var employee models.EmployeeInfo
	err := row.Scan(
		&employee.ID,
		&employee.OrgID,
		&employee.ClientID,
		&employee.Username,
		&employee.Gender,
		&employee.Phone,
		&employee.IDCard,
		&employee.Birthday,
		&employee.HireDate,
		&employee.ResignDate,
		&employee.ImgsetDir,
		&employee.ProfilePhoto,
		&employee.Description,
		&employee.IsActive,
		&employee.Created,
		&employee.CreatedBy,
		&employee.Updated,
		&employee.UpdatedBy,
		&employee.Remove)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &employee, nil
}

func DeleteEmployee(id int) error {
	query := `DELETE FROM employee_info WHERE id=?`
	_, err := config.DB.Exec(query, id)
	return err
}

func SetEmployeeProfilePhoto(id int, profilePhoto string) error {
	query := `UPDATE employee_info SET profile_photo=? WHERE id=?`
	_, err := config.DB.Exec(query, profilePhoto, id)
	return err
}
