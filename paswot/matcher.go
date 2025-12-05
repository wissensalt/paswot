package paswot

import (
	"golang.org/x/crypto/bcrypt"
)

type Matcher interface {
	Match(hashed string) bool
}

func (p *Paswot) Match(hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(p.Plain)) == nil
}

func (p *WithSalt) Match(hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(p.Plain+p.Salt)) == nil
}

func (p *WithSaltAndPepper) Match(hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(p.Plain+p.Salt+p.Pepper)) == nil
}
