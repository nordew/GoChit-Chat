package user

import (
	"errors"
	"go.uber.org/zap"
	"regexp"
	"unicode/utf8"
	"user/internal/repository"
	"user/internal/service"
	"user/pkg/hasher"
)

var (
	ErrInvalidName          = errors.New("invalid name")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidPassword      = errors.New("invalid password")
	ErrEmailAlreadyExists   = errors.New("email already exists")
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
)

type userService struct {
	userRepo repository.UserRepository

	hasher hasher.PasswordHasher

	log *zap.Logger
}

func NewUserService(userRepo repository.UserRepository, hasher hasher.PasswordHasher, log *zap.Logger) service.UserService {
	return &userService{
		userRepo: userRepo,
		hasher:   hasher,
		log:      log,
	}
}

func validateUser(name, email, password string) error {
	if name == "" || utf8.RuneCountInString(name) < 3 || utf8.RuneCountInString(name) > 30 {
		return ErrInvalidName
	}

	if !isValidEmail(email) {
		return ErrInvalidEmail
	}

	if password == "" || utf8.RuneCountInString(password) < 8 || utf8.RuneCountInString(password) > 30 {
		return ErrInvalidPassword
	}

	return nil
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email)
}
