package core

import (
	"strings"
	"testing"

	"github.com/wissensalt/paswot/rule"
)

func TestNewPaswot(t *testing.T) {
	p := NewPaswot()
	if p == nil {
		t.Fatal("NewPaswot() should not return nil")
	}
	if p.Plain != "" {
		t.Errorf("NewPaswot().Plain should be empty, got '%s'", p.Plain)
	}
}

func TestNewPaswotWithSalt(t *testing.T) {
	salt := "my-salt"
	p := NewPaswotWithSalt(salt)
	if p == nil {
		t.Fatal("NewPaswotWithSalt() should not return nil")
	}
	if p.Paswot == nil {
		t.Fatal("NewPaswotWithSalt().Paswot should not be nil")
	}
	if p.Salt != salt {
		t.Errorf("NewPaswotWithSalt().Salt should be '%s', got '%s'", salt, p.Salt)
	}
}

func TestNewPaswotWithSaltAndPepper(t *testing.T) {
	salt := "my-salt"
	pepper := "my-pepper"
	p := NewPaswotWithSaltAndPepper(salt, pepper)
	if p == nil {
		t.Fatal("NewPaswotWithSaltAndPepper() should not return nil")
	}
	if p.PaswotWithSalt == nil {
		t.Fatal("NewPaswotWithSaltAndPepper().PaswotWithSalt should not be nil")
	}
	if p.Salt != salt {
		t.Errorf("NewPaswotWithSaltAndPepper().Salt should be '%s', got '%s'", salt, p.Salt)
	}
	if p.Pepper != pepper {
		t.Errorf("NewPaswotWithSaltAndPepper().Pepper should be '%s', got '%s'", pepper, p.Pepper)
	}
}

func TestPaswot_Generate(t *testing.T) {
	t.Run("DefaultRule", func(t *testing.T) {
		p := NewPaswot()
		paswotRule := rule.DefaultRule()
		err := p.Generate(paswotRule)
		if err != nil {
			t.Fatalf("Generate() with default rule failed: %v", err)
		}

		valid, err := p.Validate(paswotRule)
		if err != nil {
			t.Fatalf("Validate() failed after generation: %v", err)
		}
		if !valid {
			t.Errorf("Generated password '%s' is not valid under default rules", p.Plain)
		}
	})

	t.Run("CustomRule", func(t *testing.T) {
		p := NewPaswot()
		paswotRule := rule.NewPaswotRuleBuilder().
			WithLength(rule.NewLengthRuleBuilder().WithMin(20).WithMax(20).Build()).
			WithCharacter(rule.NewCharacterRuleBuilder().
				WithMinUppercase(5).
				WithMinLowercase(5).
				WithMinNumber(5).
				WithMinSymbol(5).
				Build()).
			Build()

		err := p.Generate(paswotRule)
		if err != nil {
			t.Fatalf("Generate() with custom rule failed: %v", err)
		}

		valid, err := p.Validate(paswotRule)
		if err != nil {
			t.Fatalf("Validate() failed after generation: %v", err)
		}
		if !valid {
			t.Errorf("Generated password '%s' is not valid under custom rules", p.Plain)
		}
		if len(p.Plain) != 20 {
			t.Errorf("Expected password length 20, got %d", len(p.Plain))
		}
	})

	t.Run("InvalidRule", func(t *testing.T) {
		p := NewPaswot()
		// Invalid rule: min > max
		paswotRule := rule.NewPaswotRuleBuilder().
			WithLength(rule.NewLengthRuleBuilder().WithMin(10).WithMax(5).Build()).
			Build()
		err := p.Generate(paswotRule)
		if err == nil {
			t.Error("Generate() with invalid rule should have failed, but it did not")
		}
	})
}

func TestPaswot_Validate(t *testing.T) {
	defaultRule := rule.DefaultRule()

	t.Run("ValidPassword", func(t *testing.T) {
		p := &Paswot{Plain: "Valid123!"}
		valid, err := p.Validate(defaultRule)
		if err != nil {
			t.Fatalf("Validate() returned an unexpected error: %v", err)
		}
		if !valid {
			t.Error("Validate() returned false for a valid password")
		}
	})

	t.Run("EmptyPassword", func(t *testing.T) {
		p := &Paswot{Plain: ""}
		_, err := p.Validate(defaultRule)
		if err == nil {
			t.Fatal("Validate() with empty password should return an error, but it did not")
		}
		if err.Error() != "password cannot be empty" {
			t.Errorf("Unexpected error message: got '%s'", err.Error())
		}
	})

	t.Run("WhitespaceNotAllowed", func(t *testing.T) {
		p := &Paswot{Plain: "Invalid 123!"}
		_, err := p.Validate(defaultRule)
		if err == nil {
			t.Fatal("Validate() with whitespace should return an error, but it did not")
		}
		if !strings.Contains(err.Error(), "password cannot contain whitespace") {
			t.Errorf("Unexpected error message for whitespace: got '%s'", err.Error())
		}
	})

	t.Run("LengthTooShort", func(t *testing.T) {
		p := &Paswot{Plain: "V1!"}
		_, err := p.Validate(defaultRule)
		if err == nil {
			t.Fatal("Validate() with short password should return an error, but it did not")
		}
		if !strings.Contains(err.Error(), "password length must be between 8 and 16") {
			t.Errorf("Unexpected error message for length: got '%s'", err.Error())
		}
	})

	t.Run("MissingCharacterType", func(t *testing.T) {
		p := &Paswot{Plain: "validpassword"} // Missing uppercase, number, symbol
		_, err := p.Validate(defaultRule)
		if err == nil {
			t.Fatal("Validate() with missing char types should return an error, but it did not")
		}
		if !strings.Contains(err.Error(), "password must contain at least 1 uppercase characters") {
			t.Errorf("Unexpected error message for characters: got '%s'", err.Error())
		}
	})
}
