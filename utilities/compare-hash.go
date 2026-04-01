package utilities

import "golang.org/x/crypto/bcrypt"

func (u *Utility) CompareStringWithHash(plain, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}