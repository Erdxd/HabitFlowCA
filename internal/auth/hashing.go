package auth

type Hashing interface {
	Hash(password string) (string, error)
	Compare(HashedPassword, Password string) bool
}
