package repositories

import (
	"ISHC/config"
	"ISHC/models"
	"database/sql"
	"time"
)

func GetAdminById(id int) (*models.SysUser, error) {
	query := `SELECT id, org_id, client_id, username, password, real_name, sex, email, phone, mobile, description, isactive, created, createby, updated, updateby, remove 
              FROM sys_user WHERE id=?`
	row := config.DB.QueryRow(query, id)

	var admin models.SysUser
	err := row.Scan(&admin.ID, &admin.OrgID, &admin.ClientID, &admin.UserName, &admin.Password, &admin.RealName, &admin.Sex, &admin.Email, &admin.Phone, &admin.Mobile, &admin.Description, &admin.IsActive, &admin.Created, &admin.CreatedBy, &admin.Updated, &admin.UpdatedBy, &admin.Remove)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &admin, nil
}

func GetAdminByUsername(username string) (*models.SysUser, error) {
	query := `SELECT id, org_id, client_id, username, password, real_name, sex, email, phone, mobile, description, isactive, created, createby, updated, updateby, remove 
              FROM sys_user WHERE username=?`
	row := config.DB.QueryRow(query, username)

	var admin models.SysUser
	var created, updated []byte
	err := row.Scan(&admin.ID, &admin.OrgID, &admin.ClientID, &admin.UserName, &admin.Password, &admin.RealName, &admin.Sex, &admin.Email, &admin.Phone, &admin.Mobile, &admin.Description, &admin.IsActive, &created, &admin.CreatedBy, &updated, &admin.UpdatedBy, &admin.Remove)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Convert []byte to time.Time
	admin.Created, err = parseTime(created)
	if err != nil {
		return nil, err
	}
	admin.Updated, err = parseTime(updated)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

func parseTime(data []byte) (time.Time, error) {
	str := string(data)
	return time.Parse("2006-01-02 15:04:05", str)
}

func UpdateAdmin(admin *models.SysUser) error {
	query := `UPDATE sys_user SET org_id=?, client_id=?, username=?, password=?, real_name=?, sex=?, email=?, phone=?, mobile=?, description=?, isactive=?, updated=?, updateby=?, remove=? 
              WHERE id=?`
	_, err := config.DB.Exec(query, admin.OrgID, admin.ClientID, admin.UserName, admin.Password, admin.RealName, admin.Sex, admin.Email, admin.Phone, admin.Mobile, admin.Description, admin.IsActive, admin.Updated, admin.UpdatedBy, admin.Remove, admin.ID)
	return err
}
