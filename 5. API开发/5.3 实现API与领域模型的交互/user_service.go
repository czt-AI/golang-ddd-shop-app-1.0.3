package api

import (
	"context"
	"yourdomain/domain" // Replace with your actual package path
)

type userServiceServer struct {
	userRepository domain.UserRepository
}

func (s *userServiceServer) GetUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	user, err := s.userRepository.FindById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &UserResponse{
		UserId:       user.UserId,
		Username:     user.Username,
		Email:        user.Email,
		Phone:        user.Phone,
		Nickname:     user.Nickname,
		AvatarUrl:    user.AvatarUrl,
		RegistrationTime: user.RegistrationTime,
		LastLoginTime: user.LastLoginTime,
	}, nil
}

func (s *userServiceServer) CreateUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	user := &domain.User{
		Username:     req.Username,
		PasswordHash: req.PasswordHash,
		Email:        req.Email,
		Phone:        req.Phone,
		Nickname:     req.Nickname,
		AvatarUrl:    req.AvatarUrl,
		RegistrationTime: time.Now(),
		LastLoginTime: time.Time{},
	}
	if err := s.userRepository.Save(ctx, user); err != nil {
		return nil, err
	}
	return &UserResponse{
		UserId:       user.UserId,
		Username:     user.Username,
		Email:        user.Email,
		Phone:        user.Phone,
		Nickname:     user.Nickname,
		AvatarUrl:    user.AvatarUrl,
		RegistrationTime: user.RegistrationTime,
		LastLoginTime: user.LastLoginTime,
	}, nil
}

// UserRequest and UserResponse are placeholders for the actual request and response types.
type UserRequest struct {
	UserId int
}

type UserResponse struct {
	UserId       int
	Username     string
	Email        string
	Phone        string
	Nickname     string
	AvatarUrl    string
	RegistrationTime time.Time
	LastLoginTime time.Time
}
