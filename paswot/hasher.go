package paswot

import (
	"golang.org/x/crypto/bcrypt"
)

type Hasher interface {
	Hash() ([]byte, error)
}

func (p *Paswot) Hash() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p.Plain), bcrypt.DefaultCost)
}

func (p *WithSalt) Hash() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p.Plain+p.Salt), bcrypt.DefaultCost)
}

func (p *WithSaltAndPepper) Hash() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p.Plain+p.Salt+p.Pepper), bcrypt.DefaultCost)
}
