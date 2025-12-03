package core

import (
	"golang.org/x/crypto/bcrypt"
)

type PaswotMatcher interface {
	Match(hashed string) bool
}

func (p *Paswot) Match(hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(p.Plain)) == nil
}

func (p *PaswotWithSalt) Match(hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(p.Plain+p.Salt)) == nil
}

func (p *PaswotWithSaltAndPepper) Match(hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(p.Plain+p.Salt+p.Pepper)) == nil
}
