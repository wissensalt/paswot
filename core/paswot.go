package core

import (
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/wissensalt/paswot/rule"
)

type Paswot struct {
	Plain string
}

func NewPaswot() *Paswot {
	return &Paswot{}
}

type PaswotWithSalt struct {
	*Paswot
	Salt string
}

func NewPaswotWithSalt(salt string) *PaswotWithSalt {
	return &PaswotWithSalt{Paswot: &Paswot{}, Salt: salt}
}

type PaswotWithSaltAndPepper struct {
	*PaswotWithSalt
	Pepper string
}

func NewPaswotWithSaltAndPepper(salt, pepper string) *PaswotWithSaltAndPepper {
	return &PaswotWithSaltAndPepper{PaswotWithSalt: &PaswotWithSalt{Paswot: &Paswot{}, Salt: salt}, Pepper: pepper}
}

func (p *Paswot) Generate(pasRule *rule.PaswotRule) error {
	if pasRule == nil {
		pasRule = rule.DefaultRule()
	}

	_, err := pasRule.IsValid()
	if err != nil {
		return err
	}

	var passwordChars []rune
	var availableChars string

	// Uppercase
	if pasRule.Character.MinUppercase > 0 {
		availableChars += string(rule.AlphabetUpperCase)
		for i := 0; i < pasRule.Character.MinUppercase; i++ {
			char, err := getRandomChar(string(rule.AlphabetUpperCase))
			if err != nil {
				return err
			}
			passwordChars = append(passwordChars, char)
		}
	}

	// Lowercase
	if pasRule.Character.MinLowercase > 0 {
		availableChars += string(rule.AlphabetLowerCase)
		for i := 0; i < pasRule.Character.MinLowercase; i++ {
			char, err := getRandomChar(string(rule.AlphabetLowerCase))
			if err != nil {
				return err
			}
			passwordChars = append(passwordChars, char)
		}
	}

	// Numbers
	if pasRule.Character.MinNumber > 0 {
		availableChars += string(rule.Number)
		for i := 0; i < pasRule.Character.MinNumber; i++ {
			char, err := getRandomChar(string(rule.Number))
			if err != nil {
				return err
			}
			passwordChars = append(passwordChars, char)
		}
	}

	// Symbols
	if pasRule.Character.MinSymbol > 0 {
		availableChars += string(rule.Symbol)
		for i := 0; i < pasRule.Character.MinSymbol; i++ {
			char, err := getRandomChar(string(rule.Symbol))
			if err != nil {
				return err
			}
			passwordChars = append(passwordChars, char)
		}
	}

	// If no character rules are set, use all characters.
	if availableChars == "" {
		availableChars = string(rule.All)
	}

	// Fill the remaining length
	remainingLen := pasRule.Length.Min - len(passwordChars)
	for i := 0; i < remainingLen; i++ {
		char, err := getRandomChar(availableChars)
		if err != nil {
			return err
		}
		passwordChars = append(passwordChars, char)
	}

	// Shuffle the password
	shuffled, err := shuffle(passwordChars)
	if err != nil {
		return err
	}

	p.Plain = string(shuffled)

	return nil
}

func getRandomChar(charset string) (rune, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		return 0, err
	}
	return rune(charset[n.Int64()]), nil
}

func shuffle(slice []rune) ([]rune, error) {
	for i := range slice {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			// This should not happen in practice
			return nil, err
		}
		slice[i], slice[j.Int64()] = slice[j.Int64()], slice[i]
	}

	return slice, nil
}

func (p *Paswot) Validate(paswotRule *rule.PaswotRule) (bool, error) {
	if paswotRule == nil {
		paswotRule = rule.DefaultRule()
	}

	if p.Plain == "" {
		return false, errors.New("password cannot be empty")
	}

	// No Whitespace Rule
	if paswotRule.NoWhitespace != nil {
		_, err := paswotRule.NoWhitespace.Validate(p.Plain)
		if err != nil {
			return false, err
		}
	}

	// Length Rule
	if paswotRule.Length != nil {
		_, err := paswotRule.Length.Validate(p.Plain)
		if err != nil {
			return false, err
		}
	}

	// Character Rule
	if paswotRule.Character != nil {
		_, err := paswotRule.Character.Validate(p.Plain)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
