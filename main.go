package main

import (
	"github.com/wissensalt/paswot/paswot"
	"github.com/wissensalt/paswot/rule"
)

func main() {
	paswotRule := rule.NewPaswotRuleBuilder().
		WithNoWhitespace(rule.NewNoWhitespaceRule()).
		WithLength(rule.NewLengthRuleBuilder().
			WithMin(8).
			WithMax(16).
			Build()).
		WithCharacter(rule.NewCharacterRuleBuilder().
			WithMinUppercase(1).
			WithMinLowercase(1).
			WithMinNumber(1).
			WithMinSymbol(1).
			Build()).
		Build()

	testUnsaltedPassword(paswotRule)
	testSaltedPassword(paswotRule)
	testSaltedAndPepperPassword(paswotRule)
}

func testUnsaltedPassword(paswotRule *rule.PaswotRule) {
	println("########## Testing Unsalted Password ##########")
	myPaswot := paswot.NewPaswot()
	err := myPaswot.Generate(paswotRule)
	if err != nil {
		println(err.Error())
	}

	println("Generated Password: ", myPaswot.Plain)
	isValid, err := myPaswot.Validate(paswotRule)
	println("Is Valid: ", isValid)
	if err != nil {
		println("Error: ", err.Error())
	}

	hashed, err := myPaswot.Hash()
	println("Hashed Password: ", string(hashed))

	isMatch := myPaswot.Match(string(hashed))
	println("Is Match", isMatch)
}

func testSaltedPassword(paswotRule *rule.PaswotRule) {
	println("########## Testing Salted Password ##########")
	pasWithSalt := paswot.NewPaswotWithSalt("mySalt")
	err := pasWithSalt.Generate(paswotRule)
	if err != nil {
		println(err.Error())
	}

	println("Generated Password:", pasWithSalt.Plain)
	isValid, err := pasWithSalt.Validate(paswotRule)
	println("Is Valid: ", isValid)
	if err != nil {
		println("Error: ", err.Error())
	}

	salted, err := pasWithSalt.Hash()
	if err != nil {
		println(err.Error())
	}
	println("Hashed Password (With Salt): ", string(salted))

	isMatch := pasWithSalt.Match(string(salted))
	println("Is Match", isMatch)
}

func testSaltedAndPepperPassword(paswotRule *rule.PaswotRule) {
	println("########## Testing Salted and Peppered Password ##########")
	pasWithSaltAndPepper := paswot.NewPaswotWithSaltAndPepper("mySalt", "myPepper")
	err := pasWithSaltAndPepper.Generate(paswotRule)
	if err != nil {
		println(err.Error())
	}

	println("Generated Password:", pasWithSaltAndPepper.Plain)
	isValid, err := pasWithSaltAndPepper.Validate(paswotRule)
	println("Is Valid: ", isValid)
	if err != nil {
		println("Error: ", err.Error())
	}

	salted, err := pasWithSaltAndPepper.Hash()
	if err != nil {
		println(err.Error())
	}
	println("Hashed Password (With Salt and Pepper): ", string(salted))

	isMatch := pasWithSaltAndPepper.Match(string(salted))
	println("Is Match", isMatch)
}
