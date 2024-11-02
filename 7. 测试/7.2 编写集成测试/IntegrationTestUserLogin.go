package domain

import (
	"context"
	"testing"

	"yourdomain/domain" // Replace with your actual package path
	"github.com/stretchr/testify/assert"
)

func TestUserLogin(t *testing.T) {
	// Setup
	userRepository := &InMemoryUserRepository{
		users: make(map[int]*domain.User),
	}

	userService := UserServiceImpl{userRepository: userRepository}

	ctx := context.Background()

	// Test case: User login with correct credentials
	user := &domain.User{
		Username:     "testuser",
		PasswordHash: "hashedpassword",
		Email:        "testuser@example.com",
	}

	userRepository.Save(ctx, user)

	// Test case: Attempt to login with correct credentials
	loginResult, err := userService.Login(ctx, user.Username, user.PasswordHash)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, loginResult)

	// Test case: Attempt to login with incorrect password
	_, err = userService.Login(ctx, user.Username, "wrongpassword")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid password")

	// Test case: Attempt to login with non-existent username
	_, err = userService.Login(ctx, "nonexistentuser", user.PasswordHash)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")
}

// Additional test cases for user login can be added here.
// ...