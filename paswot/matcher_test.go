package paswot

import (
	"testing"
)

func TestPaswot_Match(t *testing.T) {
	p := &Paswot{
		Plain: "password",
	}

	hashed, err := p.Hash()
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	// Test with correct password
	if !p.Match(string(hashed)) {
		t.Errorf("Match() with correct password should be true, but got false")
	}

	// Test with incorrect password
	p.Plain = "wrongpassword"
	if p.Match(string(hashed)) {
		t.Errorf("Match() with incorrect password should be false, but got true")
	}
}

func TestPaswotWithSalt_Match(t *testing.T) {
	p := &WithSalt{
		Paswot: &Paswot{
			Plain: "password",
		},
		Salt: "salt",
	}

	hashed, err := p.Hash()
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	// Test with correct password
	if !p.Match(string(hashed)) {
		t.Errorf("Match() with correct password should be true, but got false")
	}

	// Test with incorrect password
	p.Plain = "wrongpassword"
	if p.Match(string(hashed)) {
		t.Errorf("Match() with incorrect password should be false, but got true")
	}
}

func TestPaswotWithSaltAndPepper_Match(t *testing.T) {
	p := &WithSaltAndPepper{
		WithSalt: &WithSalt{
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

	// Test with correct password
	if !p.Match(string(hashed)) {
		t.Errorf("Match() with correct password should be true, but got false")
	}

	// Test with incorrect password
	p.Plain = "wrongpassword"
	if p.Match(string(hashed)) {
		t.Errorf("Match() with incorrect password should be false, but got true")
	}
}
