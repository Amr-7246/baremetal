//TODO: Study That code
package helper

import (
	"fmt"
	"database/sql"
	"taskflow/helper/logger"
)

func RunMigrations(db *sql.DB, logger *logger.Logger) error {
	logger.Info("Running migrations...")

	//& Check if users table exists
		var exists bool
		query := `
			SELECT EXISTS (
				SELECT FROM information_schema.tables 
				WHERE table_name = 'users'
			)
		`
		if err := db.QueryRow(query).Scan(&exists); err != nil {
			return fmt.Errorf("failed to check table existence: %w", err)
		}

	if !exists {
		logger.Info("Creating users table...")

		//! Read migration SQL (simplified - in production use migration tool)
		migrationSQL := `
			CREATE TABLE IF NOT EXISTS users (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				email VARCHAR(255) UNIQUE NOT NULL,
				hashed_password VARCHAR(255) NOT NULL,
				full_name VARCHAR(255) NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT NOW(),
				updated_at TIMESTAMP NOT NULL DEFAULT NOW()
			);
			
			CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
			
			CREATE OR REPLACE FUNCTION update_updated_at_column()
			RETURNS TRIGGER AS $$
			BEGIN
				NEW.updated_at = NOW();
				RETURN NEW;
			END;
			$$ language 'plpgsql';
			
			CREATE TRIGGER update_users_updated_at 
				BEFORE UPDATE ON users 
				FOR EACH ROW 
				EXECUTE FUNCTION update_updated_at_column();
		`

		if _, err := db.Exec(migrationSQL); err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}

		logger.Info("Users table created successfully")
	} else {
		logger.Info("Users table already exists")
	}

	return nil
}