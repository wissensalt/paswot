package core

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestPaswot_Hash(t *testing.T) {
	p := &Paswot{
		Plain: "password",
	}

	hashed, err := p.Hash()
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	if err := bcrypt.CompareHashAndPassword(hashed, []byte("password")); err != nil {
		t.Errorf("Hashed password does not match original: %v", err)
	}
}

func TestPaswotWithSalt_Hash(t *testing.T) {
	p := &PaswotWithSalt{
		Paswot: &Paswot{
			Plain: "password",
		},
		Salt: "salt",
	}

	hashed, err := p.Hash()
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	if err := bcrypt.CompareHashAndPassword(hashed, []byte("password"+"salt")); err != nil {
		t.Errorf("Hashed password does not match original: %v", err)
	}
}

func TestPaswotWithSaltAndPepper_Hash(t *testing.T) {
	p := &PaswotWithSaltAndPepper{
		PaswotWithSalt: &PaswotWithSalt{
			Paswot: &Paswot{
				Plain: "password",
			},
			Salt: "salt",
		},
		Pepper: "pepper",
	}

	hashed, err := p.Hash()
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	if err := bcrypt.CompareHashAndPassword(hashed, []byte("password"+"salt"+"pepper")); err != nil {
		t.Errorf("Hashed password does not match original: %v", err)
	}
}
