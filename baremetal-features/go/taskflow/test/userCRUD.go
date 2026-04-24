//TODO: Study That code
package test

import (
	"context"
	"fmt"
	"taskflow/domain"
	"taskflow/helper/logger"
	"taskflow/repository"
	"time"
)

func TestUserRepository(repo repository.UserRepo, logger *logger.Logger) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//& Test 1: Create a new user
		logger.Info("Test 1: Creating a new user...")
		newUser := &domain.UserCreateDTO{
			Email:    "alice@example.com",
			Password: "SecurePass123!",
			FullName: "Alice Johnson",
		}

		createdUser, err := repo.Create(ctx, newUser)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
		logger.Infof("Created user: %s (%s)", createdUser.FullName, createdUser.Email)

	//& Test 2: Find user by email
		logger.Info("Test 2: Finding user by email...")
		foundByEmail, err := repo.FindByEmail(ctx, "alice@example.com")
		if err != nil {
			return fmt.Errorf("failed to find by email: %w", err)
		}
		if foundByEmail == nil {
			return fmt.Errorf("user not found by email")
		}
		logger.Infof("Found user by email: %s", foundByEmail.Email)

	//& Test 3: Find user by ID
		logger.Info("Test 3: Finding user by ID...")
		foundByID, err := repo.FindByID(ctx, createdUser.ID)
		if err != nil {
			return fmt.Errorf("failed to find by ID: %w", err)
		}
		if foundByID == nil {
			return fmt.Errorf("user not found by ID")
		}
		logger.Infof("Found user by ID: %s", foundByID.ID)

	//& Test 4: Check if user exists
		logger.Info("Test 4: Checking user existence...")
		exists, err := repo.Exists(ctx, "alice@example.com")
		if err != nil {
			return fmt.Errorf("failed to check existence: %w", err)
		}
		if !exists {
			return fmt.Errorf("user should exist but doesn't")
		}
		logger.Info("User existence check passed")

	//& Test 5: Update user
		logger.Info("Test 5: Updating user...")
		foundByID.FullName = "Alice B. Johnson"
		if err := repo.Update(ctx, foundByID); err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}
		logger.Infof("Updated user: %s", foundByID.FullName)

	//& Test 6: Verify password (bonus)
		logger.Info("Test 6: Verifying password...")
		userRepo := repo.(*repository.PostgresUserRepository)
		if !userRepo.VerifyPassword(foundByID, "SecurePass123!") {
			return fmt.Errorf("password verification failed")
		}
		logger.Info("Password verification passed")

	//& Test 7: Delete user (cleanup)
		logger.Info("Test 7: Deleting user...")
		if err := repo.Delete(ctx, createdUser.ID); err != nil {
			return fmt.Errorf("failed to delete user: %w", err)
		}
		logger.Info("User deleted successfully")

		return nil
}