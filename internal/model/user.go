package model

import (
    _ "github.com/jinzhu/gorm"
)

//a struct to rep user account
type User struct {
	ID uint64 `gorm:"primary_key";"AUTO_INCREMENT"`
	Email string `gorm:"type:varchar(255)”`
    PasswordHash string `gorm:"type:varchar(511)”`
    FirstName string `gorm:"type:varchar(127)”`
    LastName string `gorm:"type:varchar(127)”`
    Organizations []Organization `gorm:"foreignkey:OwnerID"`
}

// Give custom table name in our database.
func (u User) TableName() string {
    return "cc_users"
}
