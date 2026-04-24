package repository

import (
	"database/sql"
	"fmt"
	"taskflow/config"
)

func NewDBConnection(cfg *config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	//! we keeps a ready to use connection pool to save the performance via avoiding re-stablish the connection
	db.SetMaxOpenConns(25) //! to do not be baned from the DB 
	db.SetMaxIdleConns(5) //!"I will keep 5 connections ready to go at all times, even if no one is using them... not too much to avoid the memory loosing"

	return db, nil
}