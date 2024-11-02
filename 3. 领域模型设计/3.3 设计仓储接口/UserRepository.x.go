package domain

// UserRepository defines the interface for user data access.
type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindById(ctx context.Context, userId int) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Delete(ctx context.Context, userId int) error
}

// InMemoryUserRepository is an in-memory implementation of UserRepository.
type InMemoryUserRepository struct {
	users map[int]*User
}

// Save saves a user to the in-memory store.
func (repo *InMemoryUserRepository) Save(ctx context.Context, user *User) error {
	repo.users[user.UserId] = user
	return nil
}

// FindById finds a user by ID.
func (repo *InMemoryUserRepository) FindById(ctx context.Context, userId int) (*User, error) {
	user, exists := repo.users[userId]
	if !exists {
		return nil, fmt.Errorf("user not found with ID: %d", userId)
	}
	return user, nil
}

// FindByUsername finds a user by username.
func (repo *InMemoryUserRepository) FindByUsername(ctx context.Context, username string) (*User, error) {
	for _, user := range repo.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found with username: %s", username)
}

// FindByEmail finds a user by email.
func (repo *InMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	for _, user := range repo.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found with email: %s", email)
}

// Delete deletes a user from the in-memory store.
func (repo *InMemoryUserRepository) Delete(ctx context.Context, userId int) error {
	if _, exists := repo.users[userId]; !exists {
		return fmt.Errorf("user not found with ID: %d", userId)
	}
	delete(repo.users, userId)
	return nil
}
