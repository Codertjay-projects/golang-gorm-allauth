package models

import (
	AllauthConfig "github.com/codertjay/golang-gorm-allauth/config"
	"time"
)

var UserModel = AllauthConfig.AllauthConfig{}.UserModel

type AllauthUser struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	FirstName    *string   `json:"first_name"`
	LastName     *string   `json:"last_name"`
	Email        *string   `json:"email"`
	Username     *string   `json:"username"`
	Password     *string   `json:"password" validate:"required"`
	Token        *string   `json:"token" validate:"required"`
	AccessToken  *string   `json:"access_token"`
	RefreshToken *string   `json:"refresh_token"`
	IsStaff      *string   `json:"is_staff"`
	IsSuperUser  bool      `json:"is_super_user"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func GetCurrentUser() struct{} {
	// todo: check if model exist and override the existing one
	UserModel := AllauthConfig.AllauthConfig{}.UserModel
	TypeOf
	if UserModel. {

	}
	return UserModel
}

var CurrentUser = GetCurrentUser()

func (u *CurrentUser) CheckPassword(password string) {

}

func (u *CurrentUser) SetPassword(password string) {

}

func (u *CurrentUser) GetHashedPassword() *string {
	return u.Password
}
