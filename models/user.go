package models

import (
	"tugas10/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel `json:"-"`
	Role 	 string `gorm:"not null" json:"role" form:"role" valid:"required~Role is required "`
	FullName string `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Full name is required "`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required ,email~Email is not valid "`
	Password string `gorm:"not null" json:"-" form:"password" valid:"required~Password is required ,minstringlength(6)~Password must be at least 6 characters" `
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"  json:"-"`
}


func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPassword(u.Password)
	err = nil
	return
}