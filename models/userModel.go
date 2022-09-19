package models

import (
	"time"
)

// DefaultAllauthUser  /* The default model set for the user . But it can also be changed*/
type DefaultAllauthUser struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	FirstName    *string   `json:"first_name"`
	LastName     *string   `json:"last_name"`
	Email        *string   `json:"email"`
	Username     *string   `json:"username"`
	Password     *string   `json:"password" validate:"required"`
	Token        *string   `json:"token"`
	AccessToken  *string   `json:"access_token"`
	RefreshToken *string   `json:"refresh_token"`
	IsStaff      *string   `json:"is_staff"`
	IsSuperUser  bool      `json:"is_super_user"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AllauthUser struct {
	// DefaultAllauthUser /* This contains the default value a custom user should have*/
	DefaultAllauthUser
}

func (u *AllauthUser) CheckPassword(password string) {

}
func (u *AllauthUser) SetPassword(password *string) {
	u.Password = password
}

func (u *AllauthUser) GetHashedPassword() *string {
	return u.Password
}
