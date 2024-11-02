package domain

import (
	"context"
	"testing"

	"yourdomain/domain" // Replace with your actual package path
	"github.com/stretchr/testify/assert"
)

func TestUserRegistration(t *testing.T) {
	// Setup
	userRepository := &InMemoryUserRepository{
		users: make(map[int]*domain.User),
	}

	userService := UserServiceImpl{userRepository: userRepository}

	ctx := context.Background()

	// Test case: Register a new user
	user := &domain.User{
		Username:     "testuser",
		PasswordHash: "hashedpassword",
		Email:        "testuser@example.com",
	}

	err := userService.Register(ctx, user.Username, user.PasswordHash, user.Email)
	assert.NoError(t, err)

	// Assert
	userFromRepo, err := userRepository.FindByUsername(ctx, user.Username)
	assert.NoError(t, err)
	assert.NotNil(t, userFromRepo)
	assert.Equal(t, user.Username, userFromRepo.Username)
	assert.Equal(t, user.Email, userFromRepo.Email)
}

// Additional test cases for user registration can be added here.
// ...