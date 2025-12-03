package core

import (
	"golang.org/x/crypto/bcrypt"
)

type PaswotHasher interface {
	Hash() ([]byte, error)
}

func (p *Paswot) Hash() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p.Plain), bcrypt.DefaultCost)
}

func (p *PaswotWithSalt) Hash() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p.Plain+p.Salt), bcrypt.DefaultCost)
}

func (p *PaswotWithSaltAndPepper) Hash() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p.Plain+p.Salt+p.Pepper), bcrypt.DefaultCost)
}
