package user

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type PostgresStorage struct {
	db *sqlx.DB
}

func NewPostgresStorage(db *sqlx.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) Create(ctx context.Context, user User) error {
	query := `
	INSERT INTO users 	(username, firstname, lastname, email, phone)
	VALUES 				(:username, :firstname, :lastname, :email, :phone)
	`
	_, err := s.db.NamedExecContext(ctx, query, &user)
	return err
}

func (s *PostgresStorage) Get(ctx context.Context, username string) (user User, err error) {
	query := `SELECT * FROM users WHERE username=$1`
	err = s.db.GetContext(ctx, &user, query, username)
	if err == sql.ErrNoRows {
		return User{}, NotFound
	}

	return user, err
}

func (s *PostgresStorage) Update(ctx context.Context, user User) error {
	query := `
	UPDATE users SET
  		firstname 	= COALESCE(NULLIF(:firstname, 	''), firstname),
  		lastname 	= COALESCE(NULLIF(:lastname, 	''), lastname),
  		email 		= COALESCE(NULLIF(:email, 		''), email),
  		phone 		= COALESCE(NULLIF(:phone, 		''), phone)
	WHERE username = :username
	`
	resp, err := s.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return err
	}

	rAffected, err := resp.RowsAffected()
	if err != nil {
		return err
	}

	if rAffected == 0 {
		return NotFound
	}

	return nil
}

func (s *PostgresStorage) Delete(ctx context.Context, username string) error {
	query := `DELETE FROM users WHERE username=$1`
	resp, err := s.db.ExecContext(ctx, query, username)
	if err != nil {
		return err
	}

	rAffected, err := resp.RowsAffected()
	if err != nil {
		return err
	}

	if rAffected == 0 {
		return NotFound
	}

	return nil
}
