package rule

import (
	"errors"
	"fmt"
)

type LengthRule struct {
	Min int
	Max int
}

func NewLengthRule(min, max int) *LengthRule {
	return &LengthRule{Min: min, Max: max}
}

type LengthRuleBuilder struct {
	LengthRule *LengthRule
}

func NewLengthRuleBuilder() *LengthRuleBuilder {
	return &LengthRuleBuilder{LengthRule: &LengthRule{}}
}

func (builder *LengthRuleBuilder) WithMin(min int) *LengthRuleBuilder {
	builder.LengthRule.Min = min
	return builder
}

func (builder *LengthRuleBuilder) WithMax(max int) *LengthRuleBuilder {
	builder.LengthRule.Max = max
	return builder
}

func (builder *LengthRuleBuilder) Build() *LengthRule {
	return builder.LengthRule
}

func (l *LengthRule) ToString() string {
	return fmt.Sprintf("Min: %d, Max: %d", l.Min, l.Max)
}

func (l *LengthRule) Validate(password string) (bool, error) {
	if len(password) < l.Min || len(password) > l.Max {
		return false, errors.New("password length must be between " + fmt.Sprintf("%d", l.Min) + " and " + fmt.Sprintf("%d", l.Max))
	}

	return true, nil
}
