package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type jwtAuthenticator struct {
	signKey string
	logger  *zap.Logger
}

func NewAuth(signKey string, logger *zap.Logger) Authenticator {
	return &jwtAuthenticator{
		logger:  logger,
		signKey: signKey,
	}
}

type TokenClaims struct {
	UserId string
	jwt.RegisteredClaims
}

func (s *jwtAuthenticator) GenerateTokens(options *GenerateTokenClaimsOptions) (string, string, error) {
	const op = "jwtAuthenticator.GenerateTokens"

	mySigningKey := []byte(s.signKey)

	claims := TokenClaims{
		UserId: options.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "NikRo-api",
			Subject:   "client",
			ID:        uuid.NewString(),
			Audience:  []string{"NickRo-api"},
		},
	}

	refreshToken, err := s.GenerateRefreshToken(options.UserId)
	if err != nil {
		s.logger.Error("failed to generate refresh token", zap.Error(err), zap.String("op", op))
		return "", "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString(mySigningKey)
	if err != nil {
		s.logger.Error("failed to generate access token", zap.Error(err), zap.String("op", op))
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *jwtAuthenticator) GenerateRefreshToken(id string) (string, error) {
	mySigningKey := []byte(s.signKey)

	claims := TokenClaims{
		UserId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "NikRo-api",
			Subject:   "client",
			ID:        uuid.NewString(),
			Audience:  []string{"nickro-api"},
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedRefreshToken, err := refreshToken.SignedString(mySigningKey)
	if err != nil {
		s.logger.Error("failed to generate refresh token", zap.Error(err))
		return "", err
	}

	return signedRefreshToken, nil
}
