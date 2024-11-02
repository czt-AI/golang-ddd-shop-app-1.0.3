package domain

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"yourdomain/domain" // Replace with your actual package path
)

func TestRegisterUser(t *testing.T) {
	repo := &InMemoryUserRepository{
		users: make(map[int]*domain.User),
	}

	userService := UserServiceImpl{userRepository: repo}

	ctx := context.Background()

	// Test case 1: User registration with a unique username and email
	user := &domain.User{
		Username:     "testuser1",
		PasswordHash: "hashedpassword",
		Email:        "testuser1@example.com",
	}

	err := userService.Register(ctx, user.Username, user.PasswordHash, user.Email)
	assert.NoError(t, err)
	assert.NotNil(t, repo.users[user.UserId])
	assert.Equal(t, user.Username, repo.users[user.UserId].Username)
	assert.Equal(t, user.Email, repo.users[user.UserId].Email)

	// Test case 2: User registration with an existing username
	user2 := &domain.User{
		Username:     "testuser1",
		PasswordHash: "hashedpassword",
		Email:        "testuser2@example.com",
	}

	err = userService.Register(ctx, user2.Username, user2.PasswordHash, user2.Email)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "username already exists")
}

func TestLoginUser(t *testing.T) {
	repo := &InMemoryUserRepository{
		users: make(map[int]*domain.User),
	}

	userService := UserServiceImpl{userRepository: repo}

	ctx := context.Background()

	// Test case 1: User login with correct credentials
	user := &domain.User{
		Username:     "testuser1",
		PasswordHash: "hashedpassword",
		Email:        "testuser1@example.com",
	}

	repo.users[user.UserId] = user

	err := userService.Register(ctx, user.Username, user.PasswordHash, user.Email)
	assert.NoError(t, err)

	// Test case 2: User login with incorrect password
	err = userService.Login(ctx, user.Username, "wrongpassword")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid password")

	// Test case 3: User login with non-existent username
	err = userService.Login(ctx, "nonexistentuser", "correctpassword")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user not found")
}

// Additional tests for GetUser, UpdateUser, and DeleteUser can be added here.
// ...