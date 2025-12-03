package rule

import (
	"strings"
	"testing"
)

func TestNewNoWhitespaceRule(t *testing.T) {
	rule := NewNoWhitespaceRule()
	if rule == nil {
		t.Fatal("NewNoWhitespaceRule() should not return nil")
	}
}

func TestNoWhitespaceRule_Validate(t *testing.T) {
	rule := NewNoWhitespaceRule()

	testCases := []struct {
		name     string
		password string
		wantErr  bool
		errText  string
	}{
		{
			name:     "Valid password with no whitespace",
			password: "thisisavalidpassword",
			wantErr:  false,
		},
		{
			name:     "Password with a space in the middle",
			password: "invalid password",
			wantErr:  true,
			errText:  "password cannot contain whitespace",
		},
		{
			name:     "Password with a leading space",
			password: " invalidpassword",
			wantErr:  true,
			errText:  "password cannot contain whitespace",
		},
		{
			name:     "Password with a trailing space",
			password: "invalidpassword ",
			wantErr:  true,
			errText:  "password cannot contain whitespace",
		},
		{
			name:     "Empty password",
			password: "",
			wantErr:  false,
		},
		{
			name:     "Password with only a space",
			password: " ",
			wantErr:  true,
			errText:  "password cannot contain whitespace",
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
