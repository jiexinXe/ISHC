package repositories

import (
	"ISHC/config"
	"ISHC/models"
	"database/sql"
	"fmt"
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
		employee.Birthday.Time.Format(models.CtLayoutDateTime),
		employee.HireDate.Time.Format(models.CtLayoutDateTime),
		employee.ResignDate.Time.Format(models.CtLayoutDateTime),
		employee.ImgsetDir,
		employee.ProfilePhoto,
		employee.Description,
		employee.IsActive,
		employee.Created.Time.Format(models.CtLayoutDateTime),
		employee.CreatedBy,
		employee.Updated.Time.Format(models.CtLayoutDateTime),
		employee.UpdatedBy,
		employee.Remove)
	return err
}

func UpdateEmployee(employee *models.EmployeeInfo) error {
	query := `UPDATE employee_info SET 
        username=?, gender=?, phone=?, id_card=?, birthday=?, hire_date=?, resign_date=?, 
        imgset_dir=?, profile_photo=?, description=?, isactive=?, updated=?, updateby=?, remove=? 
        WHERE id=?`

	_, err := config.DB.Exec(query,
		employee.Username,
		employee.Gender,
		employee.Phone,
		employee.IDCard,
		employee.Birthday.Time.Format(models.CtLayoutDateTime),
		employee.HireDate.Time.Format(models.CtLayoutDateTime),
		employee.ResignDate.Time.Format(models.CtLayoutDateTime),
		employee.ImgsetDir,
		employee.ProfilePhoto,
		employee.Description,
		employee.IsActive,
		employee.Updated.Time.Format(models.CtLayoutDateTime),
		employee.UpdatedBy,
		employee.Remove,
		employee.ID)
	return err
}

func GetAllEmployees() ([]models.EmployeeInfo, error) {
	query := `SELECT id, org_id, client_id, username, gender, phone, id_card, birthday, hire_date, resign_date, imgset_dir, profile_photo, description, isactive, created, createby, updated, updateby, remove 
              FROM employee_info`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying employee_info: %v", err)
	}
	defer rows.Close()

	var employees []models.EmployeeInfo
	for rows.Next() {
		var employee models.EmployeeInfo
		err := rows.Scan(
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
			return nil, fmt.Errorf("error scanning employee_info row: %v", err)
		}
		employees = append(employees, employee)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over employee_info rows: %v", err)
	}

	return employees, nil
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
