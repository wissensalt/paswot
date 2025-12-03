package main

import (
	"github.com/wissensalt/paswot/core"
	"github.com/wissensalt/paswot/rule"
)

func main() {
	//paswot := core.Paswot{}
	//pass, _ := paswot.Generate(nil)
	//println(pass)

	//customPaswot := core.NewPaswot()
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
	//res, err := customPaswot.Generate(paswotRule)
	//if err != nil {
	//	println(err.Error())
	//}
	//
	//println(res)
	//isValid, err := customPaswot.Validate(res, paswotRule)
	//println("Is Valid: ", isValid)
	//if err != nil {
	//	println("Error: ", err.Error())
	//}

	pasWithSalt := core.NewPaswotWithSalt("abc123")
	res, err := pasWithSalt.Generate(paswotRule)
	if err != nil {
		println(err.Error())
	}
	println(res)
}
