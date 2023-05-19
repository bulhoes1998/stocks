package session

import (
	"golang.org/x/crypto/bcrypt"
)

type SessionService interface {
	Login(login UserLogin) (User, error)
	Logout() error
	Signup(sinup UserSignup) error
}

type sessionService struct {
	sessionRepository SessionRepository
}

func NewSessionService(sessionRepository SessionRepository) SessionService {
	return &sessionService{
		sessionRepository: sessionRepository,
	}
}

func (s *sessionService) Login(login UserLogin) (User, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(login.Password), bcrypt.DefaultCost)
	login.Password = string(hashed)
	user, err := s.sessionRepository.Get(login)
	return user, err
}

func (s *sessionService) Logout() error {
	err := s.sessionRepository.Logout()
	return err
}

func (s *sessionService) Signup(signup UserSignup) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(signup.Password), bcrypt.DefaultCost)

	signup.Password = string(hashed)

	err := s.sessionRepository.Signup(signup)

	return err
}
