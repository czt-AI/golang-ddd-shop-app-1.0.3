package domain

import (
	"context"
	"errors"
	"time"
)

// UserService defines the service interface for user management.
type UserService interface {
	Register(ctx context.Context, username, password, email string) error
	Login(ctx context.Context, username, password string) (string, error)
	GetUser(ctx context.Context, userId int) (*User, error)
	UpdateUser(ctx context.Context, userId int, nickname, avatarUrl string) error
	DeleteUser(ctx context.Context, userId int) error
}

// UserServiceImpl implements the UserService interface.
type UserServiceImpl struct {
	userRepository UserRepository
}

// Register registers a new user.
func (s *UserServiceImpl) Register(ctx context.Context, username, password, email string) error {
	if err := s.userRepository.FindByUsername(ctx, username); err == nil {
		return errors.New("username already exists")
	}
	if err := s.userRepository.FindByEmail(ctx, email); err == nil {
		return errors.New("email already exists")
	}
	user := &User{
		Username:     username,
		PasswordHash: hashPassword(password),
		Email:        email,
		RegistrationTime: time.Now(),
	}
	return s.userRepository.Save(ctx, user)
}

// Login logs in a user.
func (s *UserServiceImpl) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepository.FindByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	if !checkPasswordHash(password, user.PasswordHash) {
		return "", errors.New("invalid password")
	}
	return user.Username, nil
}

// GetUser retrieves a user by ID.
func (s *UserServiceImpl) GetUser(ctx context.Context, userId int) (*User, error) {
	return s.userRepository.FindById(ctx, userId)
}

// UpdateUser updates user information.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, userId int, nickname, avatarUrl string) error {
	user, err := s.userRepository.FindById(ctx, userId)
	if err != nil {
		return err
	}
	user.Nickname = nickname
	user.AvatarUrl = avatarUrl
	return s.userRepository.Save(ctx, user)
}

// DeleteUser deletes a user.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, userId int) error {
	return s.userRepository.Delete(ctx, userId)
}

func hashPassword(password string) string {
	// Implement password hashing logic here
	return ""
}

func checkPasswordHash(password, hash string) bool {
	// Implement password hash verification logic here
	return true
}
