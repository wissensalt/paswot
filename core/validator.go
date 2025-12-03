package core

import (
	"errors"

	"github.com/wissensalt/paswot/rule"
)

type PaswotValidator interface {
	Validate(password string, paswotRule *rule.PaswotRule) (bool, error)
}

func (p *Paswot) Validate(password string, paswotRule *rule.PaswotRule) (bool, error) {
	if paswotRule == nil {
		paswotRule = rule.DefaultRule()
	}

	if password == "" {
		return false, errors.New("password cannot be empty")
	}

	// No Whitespace Rule
	if paswotRule.NoWhitespace != nil {
		_, err := paswotRule.NoWhitespace.Validate(password)
		if err != nil {
			return false, err
		}
	}

	// Length Rule
	if paswotRule.Length != nil {
		_, err := paswotRule.Length.Validate(password)
		if err != nil {
			return false, err
		}
	}

	// Character Rule
	if paswotRule.Character != nil {
		_, err := paswotRule.Character.Validate(password)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (p *PaswotWithSalt) Validate(password string, paswotRule *rule.PaswotRule) (bool, error) {
	isValid, err := p.Paswot.Validate(password, paswotRule)
	if err != nil {
		return false, err
	}

	return isValid, nil
}
