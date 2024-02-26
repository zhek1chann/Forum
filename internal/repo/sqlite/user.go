package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/models"

	"golang.org/x/crypto/bcrypt"
)

func (s *Sqlite) GetUserByEmail(email string) (*models.User, error) {
	var u models.User
	stmt := `SELECT id, name, email, created FROM users WHERE id=?`
	err := s.db.QueryRow(stmt, email).Scan(&u.ID, &u.Name, &u.Email, &u.Created)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	return &u, nil

}

func (s *Sqlite) UpdateUserByID(string) (*models.User, error) { return nil, nil }

func (s *Sqlite) CreateUser(u models.User) error {
	stmt := `INSERT INTO users (name, email,hashed_password, created) VALUES(?, ?, ?, CURRENT_TIMESTAMP)`
	_, err := s.db.Exec(stmt, u.Name, u.Email, string(u.HashedPassword))
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.email" {
			return models.ErrDuplicateEmail
		}
		return err
	}
	return nil
}

func (s *Sqlite) GetUserByID(id int) (*models.User, error) {
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

func (s *Sqlite) Authenticate(email, password string) (int, error) {
	var id int
	var hashed_password []byte
	fmt.Println(password)
	stmt := `SELECT id, hashed_password FROM users WHERE email=?`
	err := s.db.QueryRow(stmt, email).Scan(&id, &hashed_password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrNoRecord
		}
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword(hashed_password, []byte(password))
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		}
		return 0, err
	}
	return id, nil
}

func (s *Sqlite) CreateSession(session *models.Session) error {
	stmt := `INSERT INTO sessions(user_id, token, exp_time) VALUES(?, ?, ?)`
	_, err := s.db.Exec(stmt, session.UserID, session.Token, session.ExpTime)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sqlite) DeleteSessionByUserID(userID int) error {
	stmt := `DELETE FROM sessions WHERE user_id = ?`
	if _, err := s.db.Exec(stmt, userID); err != nil {
		return err
	}
	return nil

}

func (s *Sqlite) DeleteSessionByToken(token string) error {
	stmt := `DELETE FROM sessions WHERE token = ?`
	if _, err := s.db.Exec(stmt, token); err != nil {
		return err
	}
	return nil
}
