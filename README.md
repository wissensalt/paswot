# Paswot - Password Generator & Validator Library

[![Go Version](https://img.shields.io/badge/Go-1.24.5+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

**Paswot** is a comprehensive Go library for secure password generation, validation, hashing, and verification. It provides flexible rule-based password generation with support for salt and pepper for enhanced security.

## Features

- 🔐 **Rule-based Password Generation**: Generate passwords with customizable rules
- ✅ **Password Validation**: Validate passwords against defined rules/policies  
- 🔒 **Secure Hashing**: Hash passwords using bcrypt algorithm
- 🧂 **Salt Support**: Add salt for enhanced security
- 🌶️ **Salt & Pepper Support**: Add both salt and pepper for maximum security
- 🔍 **Password Matching**: Verify passwords against hashes
- 🎛️ **Flexible Rules**: Length, character composition, and whitespace rules
- 🏗️ **Builder Pattern**: Easy-to-use builder pattern for rule configuration

## Installation

```bash
go get github.com/wissensalt/paswot
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/wissensalt/paswot/core"
    "github.com/wissensalt/paswot/rule"
)

func main() {
    // Create password rules
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

    // Generate password
    paswot := core.NewPaswot()
    err := paswot.Generate(paswotRule)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Generated password:", paswot.Plain)

    // Validate password
    isValid, err := paswot.Validate(paswotRule)
    if err != nil {
        fmt.Println("Validation error:", err)
        return
    }
    fmt.Println("Is valid:", isValid)

    // Hash password
    hashed, err := paswot.Hash()
    if err != nil {
        fmt.Println("Hashing error:", err)
        return
    }
    fmt.Println("Hashed password:", string(hashed))

    // Verify password
    isMatch := paswot.Match(string(hashed))
    fmt.Println("Password matches:", isMatch)
}
```

## API Reference

### Core Types

#### Paswot
Basic password object for generation, validation, and hashing.

```go
type Paswot struct {
    Plain string // The plain text password
}
```

#### PaswotWithSalt
Password object with salt support.

```go
type PaswotWithSalt struct {
    *Paswot
    Salt string // The salt value
}
```

#### PaswotWithSaltAndPepper
Password object with both salt and pepper support.

```go
type PaswotWithSaltAndPepper struct {
    *PaswotWithSalt
    Pepper string // The pepper value
}
```

### Constructors

```go
// Create basic password object
paswot := core.NewPaswot()

// Create password object with salt
paswotWithSalt := core.NewPaswotWithSalt("mySalt")

// Create password object with salt and pepper
paswotWithSaltAndPepper := core.NewPaswotWithSaltAndPepper("mySalt", "myPepper")
```

### Password Rules

#### Length Rule
Defines minimum and maximum password length.

```go
lengthRule := rule.NewLengthRuleBuilder().
    WithMin(8).
    WithMax(16).
    Build()
```

#### Character Rule
Defines character composition requirements.

```go
characterRule := rule.NewCharacterRuleBuilder().
    WithMinUppercase(1).    // Minimum uppercase letters
    WithMinLowercase(1).    // Minimum lowercase letters
    WithMinNumber(1).       // Minimum numbers
    WithMinSymbol(1).       // Minimum symbols
    Build()
```

#### No Whitespace Rule
Ensures password contains no whitespace characters.

```go
noWhitespaceRule := rule.NewNoWhitespaceRule()
```

#### Complete Rule Configuration

```go
paswotRule := rule.NewPaswotRuleBuilder().
    WithLength(lengthRule).
    WithCharacter(characterRule).
    WithNoWhitespace(noWhitespaceRule).
    Build()
```

### Methods

#### Generate Password
```go
err := paswot.Generate(paswotRule)
```

#### Validate Password
```go
isValid, err := paswot.Validate(paswotRule)
```

#### Hash Password
```go
// Basic hash
hashed, err := paswot.Hash()

// Hash with salt
hashedWithSalt, err := paswotWithSalt.Hash()

// Hash with salt and pepper
hashedWithSaltAndPepper, err := paswotWithSaltAndPepper.Hash()
```

#### Verify Password
```go
// Basic verification
isMatch := paswot.Match(hashedPassword)

// Verification with salt
isMatch := paswotWithSalt.Match(hashedPassword)

// Verification with salt and pepper
isMatch := paswotWithSaltAndPepper.Match(hashedPassword)
```

## Character Sets

The library uses the following character sets for password generation:

- **Uppercase**: `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
- **Lowercase**: `abcdefghijklmnopqrstuvwxyz`
- **Numbers**: `0123456789`
- **Symbols**: `!@#$%^&*()-_=+[{]};:'\",<.>/?`

## Usage Examples

### Basic Password Generation

```go
// Use default rules (8-16 chars, 1 upper, 1 lower, 1 number, 1 symbol)
paswot := core.NewPaswot()
err := paswot.Generate(rule.DefaultRule())
if err != nil {
    // Handle error
}
fmt.Println("Password:", paswot.Plain)
```

### Custom Rules

```go
// Create custom rule: 12-20 characters, no symbols required
customRule := rule.NewPaswotRuleBuilder().
    WithLength(rule.NewLengthRuleBuilder().
        WithMin(12).
        WithMax(20).
        Build()).
    WithCharacter(rule.NewCharacterRuleBuilder().
        WithMinUppercase(2).
        WithMinLowercase(2).
        WithMinNumber(2).
        Build()).
    WithNoWhitespace(rule.NewNoWhitespaceRule()).
    Build()

paswot := core.NewPaswot()
err := paswot.Generate(customRule)
```

### Password with Salt

```go
saltedPaswot := core.NewPaswotWithSalt("userUniqueSalt")
err := saltedPaswot.Generate(paswotRule)
if err != nil {
    // Handle error
}

// Hash with salt
hashed, err := saltedPaswot.Hash()
if err != nil {
    // Handle error
}

// Verify with salt
isMatch := saltedPaswot.Match(string(hashed))
```

### Password with Salt and Pepper

```go
saltAndPepperPaswot := core.NewPaswotWithSaltAndPepper("userSalt", "appPepper")
err := saltAndPepperPaswot.Generate(paswotRule)
if err != nil {
    // Handle error
}

// Hash with salt and pepper
hashed, err := saltAndPepperPaswot.Hash()
if err != nil {
    // Handle error
}

// Verify with salt and pepper
isMatch := saltAndPepperPaswot.Match(string(hashed))
```

### Password Validation Only

```go
// Validate an existing password against rules
password := "MyP@ssw0rd123"
paswot := &core.Paswot{Plain: password}

isValid, err := paswot.Validate(paswotRule)
if err != nil {
    fmt.Println("Validation failed:", err)
} else if isValid {
    fmt.Println("Password meets all requirements")
}
```

## Error Handling

The library provides detailed error messages for various scenarios:

- **Rule Validation Errors**: When password rules are invalid (e.g., character requirements exceed max length)
- **Generation Errors**: When password generation fails due to cryptographic random generation issues
- **Validation Errors**: When passwords don't meet specified rules
- **Hashing Errors**: When bcrypt hashing fails

```go
paswot := core.NewPaswot()
err := paswot.Generate(paswotRule)
if err != nil {
    switch {
    case strings.Contains(err.Error(), "character rule violates"):
        // Handle rule validation error
    case strings.Contains(err.Error(), "length must be"):
        // Handle length validation error
    default:
        // Handle other errors
    }
}
```

## Security Considerations

1. **Cryptographic Randomness**: The library uses Go's `crypto/rand` for secure random generation
2. **bcrypt Hashing**: Uses bcrypt with default cost (10) for password hashing
3. **Salt Usage**: Always use unique salts per user for password hashing
4. **Pepper Usage**: Use application-wide pepper for additional security layer
5. **Memory Security**: Consider clearing sensitive data from memory after use

## Default Rules

The library provides sensible defaults through `rule.DefaultRule()`:

- **Length**: 8-16 characters
- **Character Requirements**:
  - Minimum 1 uppercase letter
  - Minimum 1 lowercase letter  
  - Minimum 1 number
  - Minimum 1 symbol
- **No whitespace allowed**

## Dependencies

This library uses the following external dependencies:

- [`golang.org/x/crypto`](https://pkg.go.dev/golang.org/x/crypto) (v0.45.0) - For bcrypt password hashing

## Go Version

Requires Go 1.24.5 or later.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Changelog

### v1.0.0
- Initial release
- Basic password generation with rules
- Password validation
- Password hashing with bcrypt
- Salt and pepper support
- Builder pattern for rule configuration

---

For more examples and detailed documentation, please refer to the code examples in the [main.go](main.go) file.
