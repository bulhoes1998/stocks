package session

import (
	"database/sql"
	"errors"
)

type SessionRepository interface {
	Get(login UserLogin) (User, error)
	Logout() error
	Signup(signup UserSignup) error
}

type sessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) SessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (s *sessionRepository) Get(login UserLogin) (User, error) {
	query := s.db.QueryRow(`SELECT id, password FROM users WHERE email = $1`, login.Email)

	user := User{}
	if err := query.Scan(&user.Id, &user.Password); err != nil {
		return user, errors.New("username or password is incorrect")
	}

	return user, nil
}

func (s *sessionRepository) Logout() error {
	return nil
}

func (s *sessionRepository) Signup(signup UserSignup) error {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)"
	statement, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = statement.Exec(signup.Username, signup.Email, signup.Password)

	return err
}
