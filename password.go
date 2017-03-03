package hydraidp

// PasswordHasher allow create password hashes and verify them.
type PasswordHasher interface {
	Hash(password string) (hash string, err error)
	Verify(password, hash string) (newHash string, err error)
}
