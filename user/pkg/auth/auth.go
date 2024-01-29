package auth

type Authenticator interface {
	// GenerateTokens provides opportunity to encrypt access & refresh token.
	GenerateTokens(options *GenerateTokenClaimsOptions) (string, string, error)

	// GenerateRefreshToken generates refresh token
	GenerateRefreshToken(id string) (string, error)
}

type GenerateTokenClaimsOptions struct {
	UserId string
}
