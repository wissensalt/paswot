package main

import (
	"github.com/wissensalt/paswot/core"
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
	myPaswot := core.NewPaswot()
	err := myPaswot.Generate(paswotRule)
	if err != nil {
		println(err.Error())
	}

	println(myPaswot.Plain)
	isValid, err := myPaswot.Validate(paswotRule)
	println("Is Valid: ", isValid)
	if err != nil {
		println("Error: ", err.Error())
	}

	hashed, err := myPaswot.Hash()
	println("Hashed: ", string(hashed))

	isMatch := myPaswot.Match(string(hashed))
	println(isMatch)
}

func testSaltedPassword(paswotRule *rule.PaswotRule) {
	pasWithSalt := core.NewPaswotWithSalt("mySalt")
	err := pasWithSalt.Generate(paswotRule)
	if err != nil {
		println(err.Error())
	}
	println("Plain:", pasWithSalt.Plain)

	salted, err := pasWithSalt.Hash()
	if err != nil {
		println(err.Error())
	}
	println("Salted: ", string(salted))

	isMatch := pasWithSalt.Match(string(salted))
	println(isMatch)
}

func testSaltedAndPepperPassword(paswotRule *rule.PaswotRule) {
	pasWithSaltAndPepper := core.NewPaswotWithSaltAndPepper("mySalt", "myPepper")
	err := pasWithSaltAndPepper.Generate(paswotRule)
	if err != nil {
		println(err.Error())
	}
	println("Plain:", pasWithSaltAndPepper.Plain)
	println("Salt:", pasWithSaltAndPepper.Salt)
	println("Pepper:", pasWithSaltAndPepper.Pepper)
	hashed, err := pasWithSaltAndPepper.Hash()
	if err != nil {
		println(err.Error())
	}
	println("Hashed: ", string(hashed))
	isMatch := pasWithSaltAndPepper.Match(string(hashed))
	println(isMatch)
}
