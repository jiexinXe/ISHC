package models

import (
	"database/sql"
	"time"
)

type SysUser struct {
	ID          int            `json:"id" db:"id"`
	OrgID       int            `json:"org_id" db:"org_id"`
	ClientID    int            `json:"client_id" db:"client_id"`
	UserName    string         `json:"username" db:"username"`
	Password    string         `json:"password" db:"password"`
	RealName    string         `json:"real_name" db:"real_name"`
	Sex         string         `json:"sex" db:"sex"`
	Email       string         `json:"email" db:"email"`
	Phone       string         `json:"phone" db:"phone"`
	Mobile      string         `json:"mobile" db:"mobile"`
	Description string         `json:"description" db:"description"`
	IsActive    string         `json:"is_active" db:"is_active"`
	Created     time.Time      `json:"created" db:"created"`
	CreatedBy   sql.NullInt64  `json:"created_by" db:"created_by"`
	Updated     time.Time      `json:"updated" db:"updated"`
	UpdatedBy   sql.NullInt64  `json:"updated_by" db:"updated_by"`
	Remove      sql.NullString `json:"remove" db:"remove"`
}
