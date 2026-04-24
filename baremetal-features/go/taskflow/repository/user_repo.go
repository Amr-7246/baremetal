package repository

import (
	"context"
	"database/sql"
	"fmt"
	"taskflow/domain"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	Create(ctx context.Context, user *domain.UserCreateDTO) (*domain.User, error) //? teach me more about the context package then enumerate its common funcs
	FindByID(ctx context.Context, id string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, email string) (bool, error)
}

//& implements the UserRepository with PostgreSQL
	type PostgresUserRepository struct {
		db *sql.DB
	}

//& creates a new user repository
	func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository { //? why that function role
		return &PostgresUserRepository{db: db}
	}
//& Create a new user at the database
	func (r *PostgresUserRepository) Create(ctx context.Context, userDTO *domain.UserCreateDTO) (*domain.User, error) {
		//~ Hash the password
			hashedPass, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost) //? What is the coast + teach me more about that bcrypt package then enumerate its common funcs
			if err != nil {
				return nil, fmt.Errorf("ailed to hash password: %w", err)
			}
		//~ Generate the new user
			userID := uuid.New().String()  //? teach me more about the uuid package then enumerate its common funcs
			now := time.Now() //? Why I need the time while I made the default = now in the SQL code OR why I do not assign it like that updated_at = NOW()
			query := `
				INSERT INTO users (id, email, hashed_password, full_name, created_at, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6)
				RETURNING id, email, full_name, created_at, updated_at
			`
			var user domain.User 
			err = r.db.QueryRowContext( //? Explain that tow functions - "Query" & "scan" - workflow + Why the query function need the context 
					ctx, query, userID, userDTO.Email, string(hashedPass), userDTO.FullName, now, now,
				).Scan(&user.ID, &user.Email, &user.FullName, &user.CreatedAt, &user.UpdatedAt)
			if err != nil {
				return nil, fmt.Errorf("failed to create user: %w", err)
			}
		//~ Store the hashed password (for internal use)
			user.HashedPassword = string(hashedPass) //? is that secure aproatch
		return &user, nil
	}
//& Find by id 
	func (r *PostgresUserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
		query := `
			SELECT id, email, hashed_password, full_name, created_at, updated_at
			FROM users
			WHERE id = $1
		`

		var user domain.User
		err := r.db.QueryRowContext(ctx, query, id).Scan(
			&user.ID, &user.Email, &user.HashedPassword, &user.FullName,
			&user.CreatedAt, &user.UpdatedAt,
		)

		if err == sql.ErrNoRows { //? what does that return "sql.ErrNoRows "
			return nil, nil
		}
		if err != nil {
			return nil, fmt.Errorf("failed to find user by ID: %w", err)
		}

		return &user, nil
	}
//& retrieves a user by email address 
	func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
		query := `
			SELECT id, email, hashed_password, full_name, created_at, updated_at
			FROM users
			WHERE email = $1
		`

		var user domain.User
		err := r.db.QueryRowContext(ctx, query, email).Scan(
			&user.ID, &user.Email, &user.HashedPassword, &user.FullName,
			&user.CreatedAt, &user.UpdatedAt,
		)

		if err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, fmt.Errorf("failed to find user by email: %w", err)
		}

		return &user, nil
	}
//& updates an existing user
	func (r *PostgresUserRepository) Update(ctx context.Context, user *domain.User) error {
		query := `
			UPDATE users
			SET email = $1, full_name = $2, updated_at = NOW()
			WHERE id = $3
			RETURNING updated_at
		`

		err := r.db.QueryRowContext(ctx, query, user.Email, user.FullName, user.ID).Scan(&user.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}

		return nil
	}
//& removes a user from the database
	func (r *PostgresUserRepository) Delete(ctx context.Context, id string) error {
		query := `DELETE FROM users WHERE id = $1`
		result, err := r.db.ExecContext(ctx, query, id) //? What is the deffrance Execute and ExcuteContext + Why not query eigth here
		if err != nil {
			return fmt.Errorf("failed to delete user: %w", err)
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("failed to get rows affected: %w", err)
		}

		if rowsAffected == 0 {
			return fmt.Errorf("user with id %s not found", id)
		}

		return nil
	}
//& checks if a user with given email exists
	func (r *PostgresUserRepository) Exists(ctx context.Context, email string) (bool, error) {
		query := `SELECT EXISTS (SELECT 1 FROM users WHERE email= $1)` //? explain me that sql command 

		var exists bool
		err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)
		if err != nil {
			return false, fmt.Errorf("failed to check user existence: %w", err)
		}

		return exists, nil
	}
//& checks if the provided password matches the stored hash
	func (r *PostgresUserRepository) VerifyPassword(user *domain.User, password string) bool {
		err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
		return err == nil //? Why that return work well
	}