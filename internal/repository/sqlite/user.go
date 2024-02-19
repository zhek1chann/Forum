package sqlite

import (
	"database/sql"
	"errors"
	"forum/models"
	"strings"

	_ "github.com/go-sql-driver/sqlite"
	"golang.org/x/crypto/bcrypt"
)

func (s *Storage) CreateUser(u models.User) error {
	hashed_password, err := bcrypt.GenerateFromPassword(u.HashedPassword, 12)
	if err != nil {
		return err
	}
	stmt := `INSERT INTO users (name, email,hashed_password, created) VALUES(?, ?, ?, CURRENT_TIMESTAMP)`
	_, err = s.db.Exec(stmt, u.Name, u.Email, string(hashed_password))
	if err != nil {
		var mySQLError *sql.
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}
	return nil
}

func (s *Storage) DeleteUser(id int) error {
	return nil
}

func (s *Storage) GetUser(id int) (*models.User, error) {
	var u models.User
	stmt := `SELECT id, name, email, created FROM users WHERE id=?`
	err := s.db.QueryRow(stmt, id).Scan(&u.ID, &u.Name, &u.Email, &u.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	return &u, nil
}

func (s *Storage) Authenticate(email, password string) (int, error) {
	var id int
	var hashed_password []byte
	stmt := `SELECT id, hashed_password FROM users WHERE email=?`
	err := s.db.QueryRow(stmt, email).Scan(&id, &hashed_password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrNoRecord
		}
	}
	err = bcrypt.CompareHashAndPassword(hashed_password, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		}
		return 0, err
	}
	return id, nil
}
