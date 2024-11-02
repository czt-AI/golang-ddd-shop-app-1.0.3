package domain

import (
	"time"
)

// User defines the structure for a user entity.
type User struct {
	UserId       int
	Username     string
	PasswordHash string
	Email        string
	Phone        string
	Nickname     string
	AvatarUrl    string
	RegistrationTime time.Time
	LastLoginTime time.Time
}

// NewUser creates a new User instance.
func NewUser(userId int, username, passwordHash, email, phone, nickname, avatarUrl string, registrationTime, lastLoginTime time.Time) *User {
	return &User{
		UserId:          userId,
		Username:        username,
		PasswordHash:    passwordHash,
		Email:           email,
		Phone:           phone,
		Nickname:        nickname,
		AvatarUrl:       avatarUrl,
		RegistrationTime: registrationTime,
		LastLoginTime:   lastLoginTime,
	}
}

// UserDTO defines the data transfer object for a user.
type UserDTO struct {
	UserId       int
	Username     string
	Email        string
	Phone        string
	Nickname     string
	AvatarUrl    string
	RegistrationTime time.Time
	LastLoginTime time.Time
}

// ToDTO converts a User to a UserDTO.
func (u *User) ToDTO() *UserDTO {
	return &UserDTO{
		UserId:          u.UserId,
		Username:        u.Username,
		Email:           u.Email,
		Phone:           u.Phone,
		Nickname:        u.Nickname,
		AvatarUrl:       u.AvatarUrl,
		RegistrationTime: u.RegistrationTime,
		LastLoginTime:   u.LastLoginTime,
	}
}
