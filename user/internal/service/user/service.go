package user

import (
	"regexp"
	"unicode/utf8"
	"user/internal/repository"
	"user/internal/service"
	"user/pkg/auth"
	"user/pkg/hasher"
	userErrors "user/pkg/user_errors"

	"go.uber.org/zap"
)

type userService struct {
	userRepo repository.UserRepository

	hasher hasher.PasswordHasher
	auth   auth.Authenticator

	log *zap.Logger
}

func NewUserService(userRepo repository.UserRepository, hasher hasher.PasswordHasher, auth auth.Authenticator, log *zap.Logger) service.UserService {
	return &userService{
		userRepo: userRepo,
		hasher:   hasher,
		auth:     auth,
		log:      log,
	}
}

func validateUser(name, email, password string) error {
	if name == "" || utf8.RuneCountInString(name) < 3 || utf8.RuneCountInString(name) > 30 {
		return userErrors.ErrInvalidName
	}

	if !IsValidEmail(email) {
		return userErrors.ErrInvalidEmail
	}

	if password == "" || utf8.RuneCountInString(password) < 8 || utf8.RuneCountInString(password) > 30 {
		return userErrors.ErrInvalidPassword
	}

	return nil
}

func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email)
}
