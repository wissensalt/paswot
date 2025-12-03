package rule

import (
	"fmt"
	"strings"
)

type NoWhitespaceRule struct{}

func (r NoWhitespaceRule) Validate(password string) (bool, error) {
	if strings.Contains(password, " ") {
		return false, fmt.Errorf("password cannot contain whitespace")
	}

	return true, nil
}

func NewNoWhitespaceRule() *NoWhitespaceRule {
	return &NoWhitespaceRule{}
}
