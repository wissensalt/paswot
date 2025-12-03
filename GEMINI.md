# Paswot Library Context for Gemini AI

## Overview
Paswot is a Go library for secure password generation, validation, and hashing. This document provides context for AI assistants working with this codebase.

## Library Structure

### Core Package (`core/`)
- **`paswot.go`**: Main password structures and generation logic
- **`hasher.go`**: Password hashing functionality using bcrypt
- **`matcher.go`**: Password verification/matching functionality

### Rule Package (`rule/`)
- **`rule.go`**: Main rule structures and builders
- **`length_rule.go`**: Password length validation rules
- **`character_rule.go`**: Character composition rules
- **`no_whitespace_rule.go`**: Whitespace restriction rules

## Key Types and Their Purpose

### Password Objects
1. **`Paswot`**: Basic password object with plain text
2. **`PaswotWithSalt`**: Extends Paswot with salt for enhanced security
3. **`PaswotWithSaltAndPepper`**: Extends PaswotWithSalt with pepper

### Rule Objects
1. **`PaswotRule`**: Container for all password rules
2. **`LengthRule`**: Defines min/max password length
3. **`CharacterRule`**: Defines character composition requirements
4. **`NoWhitespaceRule`**: Prevents whitespace in passwords

## Key Functionalities

### Password Generation
- Uses cryptographically secure random generation (`crypto/rand`)
- Ensures required character types are included
- Shuffles characters for randomness
- Validates against provided rules

### Password Validation
- Checks length constraints
- Verifies character composition requirements
- Ensures no whitespace (if rule is set)
- Provides detailed error messages

### Password Hashing
- Uses bcrypt algorithm with default cost (10)
- Supports plain password hashing
- Supports salted password hashing
- Supports salt+pepper password hashing

### Password Verification
- Compares plain passwords against bcrypt hashes
- Handles salted password verification
- Handles salt+pepper password verification

## Character Sets Used
- **Uppercase**: `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
- **Lowercase**: `abcdefghijklmnopqrstuvwxyz`
- **Numbers**: `0123456789`
- **Symbols**: `!@#$%^&*()-_=+[{]};:'\",<.>/?`

## Builder Pattern Implementation
The library extensively uses the builder pattern for:
- `PaswotRuleBuilder`: Building complete rule sets
- `LengthRuleBuilder`: Building length rules
- `CharacterRuleBuilder`: Building character rules

## Security Features
1. **Cryptographic Randomness**: Uses `crypto/rand` for secure random generation
2. **bcrypt Hashing**: Industry-standard password hashing
3. **Salt Support**: Prevents rainbow table attacks
4. **Pepper Support**: Additional application-level security

## Dependencies
- `golang.org/x/crypto/bcrypt` for password hashing
- Standard Go `crypto/rand` for secure randomness
- Standard Go `errors` and `fmt` for error handling

## Default Configuration
The `DefaultRule()` function provides:
- Length: 8-16 characters
- Minimum 1 uppercase letter
- Minimum 1 lowercase letter
- Minimum 1 number
- Minimum 1 symbol
- No whitespace allowed

## Common Usage Patterns

### Basic Usage
```go
paswot := core.NewPaswot()
err := paswot.Generate(rule.DefaultRule())
hashed, _ := paswot.Hash()
isMatch := paswot.Match(string(hashed))
```

### Custom Rules
```go
customRule := rule.NewPaswotRuleBuilder().
    WithLength(rule.NewLengthRuleBuilder().WithMin(12).WithMax(20).Build()).
    WithCharacter(rule.NewCharacterRuleBuilder().WithMinUppercase(2).Build()).
    Build()
```

### With Salt and Pepper
```go
paswot := core.NewPaswotWithSaltAndPepper("userSalt", "appPepper")
err := paswot.Generate(rule.DefaultRule())
hashed, _ := paswot.Hash()
```

## Error Handling
The library provides specific error types for:
- Rule validation failures
- Password generation failures
- Password validation failures
- Character requirement violations
- Length constraint violations

## Testing and Validation
The `main.go` file demonstrates all key functionalities:
- Unsalted password generation and validation
- Salted password generation and validation  
- Salt+pepper password generation and validation

This context should help AI assistants understand the codebase structure, purpose, and proper usage patterns when working with the Paswot library.
