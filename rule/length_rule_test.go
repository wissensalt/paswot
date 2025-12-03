package rule

import (
	"strings"
	"testing"
)

func TestNewLengthRule(t *testing.T) {
	rule := NewLengthRule(8, 16)
	if rule.Min != 8 {
		t.Errorf("NewLengthRule Min = %d; want 8", rule.Min)
	}
	if rule.Max != 16 {
		t.Errorf("NewLengthRule Max = %d; want 16", rule.Max)
	}
}

func TestLengthRuleBuilder(t *testing.T) {
	builder := NewLengthRuleBuilder()
	if builder.LengthRule == nil {
		t.Fatal("NewLengthRuleBuilder returned a builder with a nil LengthRule")
	}

	rule := builder.
		WithMin(10).
		WithMax(20).
		Build()

	if rule.Min != 10 {
		t.Errorf("WithMin not set correctly. Got %d, want 10", rule.Min)
	}
	if rule.Max != 20 {
		t.Errorf("WithMax not set correctly. Got %d, want 20", rule.Max)
	}
}

func TestLengthRule_ToString(t *testing.T) {
	rule := NewLengthRule(8, 16)
	expected := "Min: 8, Max: 16"
	if str := rule.ToString(); str != expected {
		t.Errorf("ToString() = %q; want %q", str, expected)
	}
}

func TestLengthRule_Validate(t *testing.T) {
	rule := NewLengthRule(8, 12)

	testCases := []struct {
		name     string
		password string
		wantErr  bool
		errText  string
	}{
		{
			name:     "Valid length",
			password: "password10", // length 10
			wantErr:  false,
		},
		{
			name:     "Length too short",
			password: "pass", // length 4
			wantErr:  true,
			errText:  "password length must be between 8 and 12",
		},
		{
			name:     "Length too long",
			password: "longpassword123", // length 15
			wantErr:  true,
			errText:  "password length must be between 8 and 12",
		},
		{
			name:     "Length at min boundary",
			password: "eightlen", // length 8
			wantErr:  false,
		},
		{
			name:     "Length at max boundary",
			password: "twelvelength", // length 12
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := rule.Validate(tc.password)
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
