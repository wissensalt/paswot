package rule

import (
	"fmt"
	"strings"
)

type Charset string

const (
	AlphabetUpperCase Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphabetLowerCase Charset = "abcdefghijklmnopqrstuvwxyz"
	Number            Charset = "0123456789"
	Symbol            Charset = "!@#$%^&*()-_=+[{]};:'\",<.>/?"
	All               Charset = AlphabetUpperCase + AlphabetLowerCase + Number + Symbol
)

type CharacterRule struct {
	MinUppercase int
	MinLowercase int
	MinNumber    int
	MinSymbol    int
}

func (c *CharacterRule) Sum() int {
	return c.MinUppercase + c.MinLowercase + c.MinNumber + c.MinSymbol
}

func NewCharacterRule(minUpperCase, minLowerCase, minNumber, minSymbol int) *CharacterRule {
	return &CharacterRule{MinUppercase: minUpperCase, MinLowercase: minLowerCase, MinNumber: minNumber, MinSymbol: minSymbol}
}

type CharacterRuleBuilder struct {
	CharacterRule *CharacterRule
}

func NewCharacterRuleBuilder() *CharacterRuleBuilder {
	return &CharacterRuleBuilder{CharacterRule: &CharacterRule{}}
}

func (builder *CharacterRuleBuilder) WithMinUppercase(minUppercase int) *CharacterRuleBuilder {
	builder.CharacterRule.MinUppercase = minUppercase
	return builder
}

func (builder *CharacterRuleBuilder) WithMinLowercase(minLowercase int) *CharacterRuleBuilder {
	builder.CharacterRule.MinLowercase = minLowercase
	return builder
}

func (builder *CharacterRuleBuilder) WithMinNumber(minNumber int) *CharacterRuleBuilder {
	builder.CharacterRule.MinNumber = minNumber
	return builder
}

func (builder *CharacterRuleBuilder) WithMinSymbol(minSymbol int) *CharacterRuleBuilder {
	builder.CharacterRule.MinSymbol = minSymbol
	return builder
}

func (builder *CharacterRuleBuilder) Build() *CharacterRule {
	return builder.CharacterRule
}

func (c *CharacterRule) ToString() string {
	var text string
	text += "MinUppercase: " + fmt.Sprintf("%d", c.MinUppercase) + ", "
	text += "MinLowercase: " + fmt.Sprintf("%d", c.MinLowercase) + ", "
	text += "MinNumber: " + fmt.Sprintf("%d", c.MinNumber) + ", "
	text += "MinSymbol: " + fmt.Sprintf("%d", c.MinSymbol)
	return text
}

func (c *CharacterRule) Validate(password string) (bool, error) {
	uppercaseCount := 0
	for _, char := range password {
		if strings.ContainsRune(string(AlphabetUpperCase), char) {
			uppercaseCount++
		}
	}
	if uppercaseCount < c.MinUppercase {
		return false, fmt.Errorf("password must contain at least %d uppercase characters", c.MinUppercase)
	}

	lowercaseCount := 0
	for _, char := range password {
		if strings.ContainsRune(string(AlphabetLowerCase), char) {
			lowercaseCount++
		}
	}
	if lowercaseCount < c.MinLowercase {
		return false, fmt.Errorf("password must contain at least %d lowercase characters", c.MinLowercase)
	}

	numberCount := 0
	for _, char := range password {
		if strings.ContainsRune(string(Number), char) {
			numberCount++
		}
	}
	if numberCount < c.MinNumber {
		return false, fmt.Errorf("password must contain at least %d number characters", c.MinNumber)
	}

	symbolCount := 0
	for _, char := range password {
		if strings.ContainsRune(string(Symbol), char) {
			symbolCount++
		}
	}
	if symbolCount < c.MinSymbol {
		return false, fmt.Errorf("password must contain at least %d symbol characters", c.MinSymbol)
	}

	return true, nil
}
