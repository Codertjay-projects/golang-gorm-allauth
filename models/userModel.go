package models

import "time"

type User struct {
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

func (u *User) CheckPassword(password string) {

}

func (u *User) SetPassword(password string) {

}

func (u *User) GetHashedPassword() *string {
	return u.Password
}
