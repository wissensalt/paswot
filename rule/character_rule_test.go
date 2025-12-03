package rule

import (
	"strings"
	"testing"
)

func TestNewCharacterRule(t *testing.T) {
	rule := NewCharacterRule(1, 2, 3, 4)
	if rule.MinUppercase != 1 || rule.MinLowercase != 2 || rule.MinNumber != 3 || rule.MinSymbol != 4 {
		t.Errorf("NewCharacterRule did not set values correctly. Got: %+v", rule)
	}
}

func TestCharacterRule_Sum(t *testing.T) {
	rule := &CharacterRule{MinUppercase: 1, MinLowercase: 2, MinNumber: 3, MinSymbol: 4}
	expectedSum := 10
	if sum := rule.Sum(); sum != expectedSum {
		t.Errorf("Sum() = %d; want %d", sum, expectedSum)
	}
}

func TestCharacterRuleBuilder(t *testing.T) {
	builder := NewCharacterRuleBuilder()
	if builder.CharacterRule == nil {
		t.Fatal("NewCharacterRuleBuilder returned a builder with a nil CharacterRule")
	}

	rule := builder.
		WithMinUppercase(2).
		WithMinLowercase(3).
		WithMinNumber(4).
		WithMinSymbol(5).
		Build()

	if rule.MinUppercase != 2 {
		t.Errorf("WithMinUppercase not set correctly. Got %d, want 2", rule.MinUppercase)
	}
	if rule.MinLowercase != 3 {
		t.Errorf("WithMinLowercase not set correctly. Got %d, want 3", rule.MinLowercase)
	}
	if rule.MinNumber != 4 {
		t.Errorf("WithMinNumber not set correctly. Got %d, want 4", rule.MinNumber)
	}
	if rule.MinSymbol != 5 {
		t.Errorf("WithMinSymbol not set correctly. Got %d, want 5", rule.MinSymbol)
	}
}

func TestCharacterRule_ToString(t *testing.T) {
	rule := &CharacterRule{MinUppercase: 1, MinLowercase: 1, MinNumber: 1, MinSymbol: 1}
	expected := "MinUppercase: 1, MinLowercase: 1, MinNumber: 1, MinSymbol: 1"
	if str := rule.ToString(); str != expected {
		t.Errorf("ToString() = %q; want %q", str, expected)
	}
}

func TestCharacterRule_Validate(t *testing.T) {
	testCases := []struct {
		name     string
		rule     *CharacterRule
		password string
		wantErr  bool
		errText  string
	}{
		{
			name:     "Valid password",
			rule:     NewCharacterRule(1, 1, 1, 1),
			password: "ValidPass1!",
			wantErr:  false,
		},
		{
			name:     "Missing uppercase",
			rule:     NewCharacterRule(1, 1, 1, 1),
			password: "validpass1!",
			wantErr:  true,
			errText:  "password must contain at least 1 uppercase characters",
		},
		{
			name:     "Missing multiple uppercase",
			rule:     NewCharacterRule(2, 1, 1, 1),
			password: "Validpass1!",
			wantErr:  true,
			errText:  "password must contain at least 2 uppercase characters",
		},
		{
			name:     "Missing lowercase",
			rule:     NewCharacterRule(1, 1, 1, 1),
			password: "VALIDPASS1!",
			wantErr:  true,
			errText:  "password must contain at least 1 lowercase characters",
		},
		{
			name:     "Missing number",
			rule:     NewCharacterRule(1, 1, 1, 1),
			password: "ValidPass!",
			wantErr:  true,
			errText:  "password must contain at least 1 number characters",
		},
		{
			name:     "Missing symbol",
			rule:     NewCharacterRule(1, 1, 1, 1),
			password: "ValidPass1",
			wantErr:  true,
			errText:  "password must contain at least 1 symbol characters",
		},
		{
			name:     "No requirements",
			rule:     NewCharacterRule(0, 0, 0, 0),
			password: "anypassword",
			wantErr:  false,
		},
		{
			name:     "Multiple requirements not met",
			rule:     NewCharacterRule(2, 2, 2, 2),
			password: "Vp1!",
			wantErr:  true,
			errText:  "password must contain at least 2 uppercase characters",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.rule.Validate(tc.password)
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tc.errText) {
				t.Errorf("Validate() error = %q, want error containing %q", err.Error(), tc.errText)
			}
		})
	}
}
