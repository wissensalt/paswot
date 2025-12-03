package rule

import (
	"testing"
)

func TestNewPaswotRuleBuilder(t *testing.T) {
	builder := NewPaswotRuleBuilder()
	if builder == nil {
		t.Fatal("NewPaswotRuleBuilder() should not return nil")
	}
	if builder.PaswotRule == nil {
		t.Fatal("NewPaswotRuleBuilder().PaswotRule should not be nil")
	}
}

func TestPaswotRuleBuilder_Build(t *testing.T) {
	lengthRule := NewLengthRule(1, 2)
	charRule := NewCharacterRule(1, 1, 1, 1)
	noWhitespaceRule := NewNoWhitespaceRule()

	builder := NewPaswotRuleBuilder().
		WithLength(lengthRule).
		WithCharacter(charRule).
		WithNoWhitespace(noWhitespaceRule)

	rule := builder.Build()

	if rule.Length != lengthRule {
		t.Error("Builder did not set Length rule correctly")
	}
	if rule.Character != charRule {
		t.Error("Builder did not set Character rule correctly")
	}
	if rule.NoWhitespace != noWhitespaceRule {
		t.Error("Builder did not set NoWhitespace rule correctly")
	}
}

func TestDefaultRule(t *testing.T) {
	rule := DefaultRule()
	if rule == nil {
		t.Fatal("DefaultRule() should not return nil")
	}

	if rule.Length.Min != 8 || rule.Length.Max != 16 {
		t.Errorf("DefaultRule incorrect Length. Expected 8-16, got %d-%d", rule.Length.Min, rule.Length.Max)
	}

	if rule.Character.MinUppercase != 1 || rule.Character.MinLowercase != 1 || rule.Character.MinNumber != 1 || rule.Character.MinSymbol != 1 {
		t.Error("DefaultRule incorrect Character requirements")
	}

	if rule.NoWhitespace == nil {
		t.Error("DefaultRule should include NoWhitespace rule")
	}

	valid, err := rule.IsValid()
	if !valid || err != nil {
		t.Errorf("DefaultRule should be valid, but got valid=%v, err=%v", valid, err)
	}
}

func TestPaswotRule_IsValid(t *testing.T) {
	testCases := []struct {
		name    string
		rule    *PaswotRule
		wantErr bool
		errText string
	}{
		{
			name:    "Valid default rule",
			rule:    DefaultRule(),
			wantErr: false,
		},
		{
			name: "Invalid length min > max",
			rule: NewPaswotRuleBuilder().
				WithLength(NewLengthRule(10, 5)).
				Build(),
			wantErr: true,
			errText: "length rule min violates max rule",
		},
		{
			name: "Invalid character sum > length max",
			rule: NewPaswotRuleBuilder().
				WithLength(NewLengthRule(4, 8)).
				WithCharacter(NewCharacterRule(3, 3, 3, 0)). // sum = 9
				Build(),
			wantErr: true,
			errText: "character rule violates max length rule",
		},
		{
			name: "Valid character sum == length max",
			rule: NewPaswotRuleBuilder().
				WithLength(NewLengthRule(4, 8)).
				WithCharacter(NewCharacterRule(2, 2, 2, 2)). // sum = 8
				Build(),
			wantErr: false,
		},
		{
			name: "Invalid char sum is zero but min length > 0",
			rule: NewPaswotRuleBuilder().
				WithLength(NewLengthRule(1, 8)).
				WithCharacter(NewCharacterRule(0, 0, 0, 0)).
				Build(),
			wantErr: true,
			errText: "character rule violates min length rule",
		},
		{
			name: "Invalid no whitespace with zero char sum",
			rule: NewPaswotRuleBuilder().
				WithCharacter(NewCharacterRule(0, 0, 0, 0)).
				WithNoWhitespace(NewNoWhitespaceRule()).
				Build(),
			wantErr: true,
			errText: "no whitespace rule violates the character rule",
		},
		{
			name: "Valid rule with only length",
			rule: NewPaswotRuleBuilder().
				WithLength(NewLengthRule(8, 16)).
				Build(),
			wantErr: false,
		},
		{
			name: "Valid rule with only character",
			rule: NewPaswotRuleBuilder().
				WithCharacter(NewCharacterRule(1, 1, 0, 0)).
				Build(),
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.rule.IsValid()
			if (err != nil) != tc.wantErr {
				t.Errorf("IsValid() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if err != nil && err.Error() != tc.errText {
				t.Errorf("IsValid() error = %q, want error text %q", err.Error(), tc.errText)
			}
		})
	}
}
