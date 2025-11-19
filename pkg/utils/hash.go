package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword mengubah plain password menjadi bcrypt hash.
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// CheckPassword memastikan plain password cocok dengan hash.
func CheckPassword(hash string, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
