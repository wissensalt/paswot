# Paswot Library Context for Claude AI

## Project Overview
**Paswot** is a comprehensive Go library designed for secure password management, including generation, validation, hashing, and verification. The library follows Go best practices and implements a clean architecture with separation of concerns.

## Architecture Design

### Package Structure
```
paswot/
├── core/           # Core password functionality
├── rule/           # Password rule definitions and validation
├── main.go         # Example usage and demonstration
├── go.mod          # Go module definition
└── go.sum          # Dependency checksums
```

### Core Package Responsibilities
- **Password Generation**: Cryptographically secure random password creation
- **Password Hashing**: bcrypt-based password hashing with salt/pepper support
- **Password Verification**: Secure password matching against hashes

### Rule Package Responsibilities
- **Rule Definition**: Flexible password policy configuration
- **Rule Validation**: Ensure passwords meet defined criteria
- **Builder Patterns**: Fluent API for rule construction

## Technical Implementation Details

### Cryptographic Security
- **Random Generation**: Uses `crypto/rand` for cryptographically secure randomness
- **Hashing Algorithm**: bcrypt with default cost factor (10)
- **Salt Support**: Per-user unique salts to prevent rainbow table attacks
- **Pepper Support**: Application-wide secret for additional security layer

### Design Patterns
1. **Builder Pattern**: Implemented for rule construction with method chaining
2. **Strategy Pattern**: Different password types (plain, salted, salt+pepper)
3. **Interface Segregation**: `PaswotHasher` and `PaswotMatcher` interfaces

### Data Structures

#### Core Types
```go
type Paswot struct {
    Plain string
}

type PaswotWithSalt struct {
    *Paswot
    Salt string
}

type PaswotWithSaltAndPepper struct {
    *PaswotWithSalt
    Pepper string
}
```

#### Rule Types
```go
type PaswotRule struct {
    Length       *LengthRule
    Character    *CharacterRule
    NoWhitespace *NoWhitespaceRule
}
```

### Algorithm Flow

#### Password Generation
1. Validate input rules for consistency
2. Generate required character types based on rules
3. Fill remaining length with random characters from allowed sets
4. Cryptographically shuffle the character array
5. Convert to string

#### Password Validation
1. Check if password is non-empty
2. Validate against no-whitespace rule (if set)
3. Validate against length constraints
4. Validate against character composition requirements

#### Password Hashing
- **Basic**: `bcrypt.GenerateFromPassword(password, cost)`
- **With Salt**: `bcrypt.GenerateFromPassword(password+salt, cost)`
- **With Salt+Pepper**: `bcrypt.GenerateFromPassword(password+salt+pepper, cost)`

## Character Set Configuration
The library defines four character categories:
- **AlphabetUpperCase**: 26 uppercase letters
- **AlphabetLowerCase**: 26 lowercase letters
- **Number**: 10 digits (0-9)
- **Symbol**: 32 special characters for password complexity

## Rule Validation Logic
The library implements sophisticated rule validation:
- Character requirements cannot exceed maximum length
- Minimum character requirements cannot be zero if maximum length is positive
- No-whitespace rule compatibility with character rules

## Error Handling Strategy
- **Validation Errors**: Detailed messages for rule violations
- **Generation Errors**: Cryptographic operation failures
- **Hashing Errors**: bcrypt operation failures
- **Type Safety**: Strong typing prevents runtime errors

## Performance Considerations
- **Memory Efficiency**: Minimal memory allocation during generation
- **CPU Efficiency**: Optimized character selection and shuffling
- **Security vs Performance**: bcrypt cost factor balances security and speed

## Extensibility Points
1. **New Rule Types**: Can be added to `PaswotRule` struct
2. **Custom Character Sets**: Charset type can be extended
3. **Alternative Hashing**: Interface allows different hashing implementations
4. **Additional Password Types**: New structs can embed existing types

## Usage Patterns

### Fluent API Design
```go
rule := rule.NewPaswotRuleBuilder().
    WithLength(/* config */).
    WithCharacter(/* config */).
    WithNoWhitespace(/* config */).
    Build()
```

### Progressive Enhancement
- Start with `Paswot` for basic use cases
- Upgrade to `PaswotWithSalt` for user-specific security
- Use `PaswotWithSaltAndPepper` for maximum security

## Testing Strategy
The `main.go` file provides comprehensive testing scenarios:
- Basic password functionality
- Salt-enhanced password security
- Salt+pepper maximum security configuration

## Dependencies Analysis
- **golang.org/x/crypto**: Provides bcrypt implementation
  - Version: v0.45.0
  - Purpose: Cryptographically secure password hashing
  - Security: Industry-standard, well-audited library

## Security Best Practices Implemented
1. **No Plain Text Storage**: Passwords immediately hashed
2. **Salt Uniqueness**: Each user should have unique salt
3. **Pepper Secrecy**: Application-wide secret never exposed
4. **Secure Defaults**: Default rules enforce strong passwords
5. **Cryptographic Randomness**: No predictable password patterns

## Integration Considerations
- **Database Storage**: Store only hashed passwords with salts
- **Configuration Management**: Pepper values should be environment-specific
- **Logging**: Never log plain text passwords or salts
- **API Design**: Consider rate limiting for password operations

This comprehensive context should enable Claude AI to understand the library's architecture, security model, and proper usage patterns for any development or analysis tasks.
