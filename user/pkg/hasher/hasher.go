package hasher

// PasswordHasher defines the interface for the password hasher
type PasswordHasher interface {

	// Hash hashes a password
	Hash(password string) (string, error)

	// Compare compares a hashed password with a password
	Compare(hashedPassword, password string) error
}
