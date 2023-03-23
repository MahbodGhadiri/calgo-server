package db

import (
	"database/sql"
)

type User struct {
	Base        Base `gorm:"embedded"`
	Email       string
	Password    string
	ActivatedAt sql.NullTime
}
