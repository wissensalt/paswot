package rule

import (
	"errors"
)

type PaswotRule struct {
	Length       *LengthRule
	Character    *CharacterRule
	NoWhitespace *NoWhitespaceRule
}

func (p *PaswotRule) IsValid() (bool, error) {
	if p.Character != nil {
		if p.Length != nil {
			// Character rule violates min length rule
			if p.Character.Sum() == 0 && p.Length.Min > 0 {
				return false, errors.New("character rule violates min length rule")
			}

			// Character rule violates max length rule
			if p.Character.Sum() > p.Length.Max {
				return false, errors.New("character rule violates max length rule")
			}
		}

		// No Whitespace rule violates the character rule
		if p.NoWhitespace != nil && p.Character.Sum() == 0 {
			return false, errors.New("no whitespace rule violates the character rule")
		}
	}

	return true, nil
}

type PaswotRuleBuilder struct {
	PaswotRule *PaswotRule
}

func NewPaswotRuleBuilder() *PaswotRuleBuilder {
	return &PaswotRuleBuilder{PaswotRule: &PaswotRule{}}
}

func (builder *PaswotRuleBuilder) WithLength(length *LengthRule) *PaswotRuleBuilder {
	builder.PaswotRule.Length = length
	return builder
}

func (builder *PaswotRuleBuilder) WithCharacter(character *CharacterRule) *PaswotRuleBuilder {
	builder.PaswotRule.Character = character
	return builder
}

func (builder *PaswotRuleBuilder) WithNoWhitespace(noWhitespace *NoWhitespaceRule) *PaswotRuleBuilder {
	builder.PaswotRule.NoWhitespace = noWhitespace
	return builder
}

func (builder *PaswotRuleBuilder) Build() *PaswotRule {
	return builder.PaswotRule
}

func DefaultRule() *PaswotRule {
	return NewPaswotRuleBuilder().
		WithLength(NewLengthRuleBuilder().
			WithMin(8).
			WithMax(16).
			Build()).
		WithCharacter(NewCharacterRuleBuilder().
			WithMinUppercase(1).
			WithMinLowercase(1).
			WithMinNumber(1).
			WithMinSymbol(1).
			Build()).
		WithNoWhitespace(NewNoWhitespaceRule()).
		Build()
}
