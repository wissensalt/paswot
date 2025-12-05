# Paswot Library Context for AI Agents

## Executive Summary
**Paswot** is a production-ready Go library providing comprehensive password management capabilities including secure generation, policy-based validation, cryptographic hashing, and verification with salt/pepper support.

## Core Value Proposition
- **Security-First Design**: Implements industry best practices for password security
- **Flexibility**: Rule-based configuration allows customization for various security requirements
- **Ease of Use**: Builder pattern provides intuitive API for complex configurations
- **Production Ready**: Comprehensive error handling and validation

## Technical Stack
- **Language**: Go 1.24.5+
- **Cryptography**: golang.org/x/crypto/bcrypt
- **Architecture**: Clean architecture with separated concerns
- **Design Patterns**: Builder, Strategy, Interface Segregation

## Functional Capabilities

### 1. Password Generation
**Purpose**: Create cryptographically secure passwords meeting specific criteria
**Implementation**: 
- Uses `crypto/rand` for secure randomness
- Ensures character type requirements are met
- Shuffles characters to prevent patterns
- Validates rules before generation

### 2. Password Validation  
**Purpose**: Verify passwords against defined policies
**Implementation**:
- Length constraint checking
- Character composition validation
- Whitespace restriction enforcement
- Detailed error reporting

### 3. Password Hashing
**Purpose**: Secure password storage using industry-standard algorithms
**Implementation**:
- bcrypt hashing with configurable cost
- Salt support for unique per-user hashing
- Pepper support for application-level security
- Protection against rainbow table attacks

### 4. Password Verification
**Purpose**: Authenticate users by comparing passwords with stored hashes
**Implementation**:
- Constant-time comparison via bcrypt
- Handles salted hash verification
- Supports salt+pepper configurations

## Domain Model

### Entity Hierarchy
```
Paswot (base)
├── Plain: string
│
├── PaswotWithSalt
│   ├── *Paswot (embedded)
│   └── Salt: string
│
└── PaswotWithSaltAndPepper
    ├── *PaswotWithSalt (embedded)
    └── Pepper: string
```

### Rule System
```
PaswotRule (aggregate root)
├── LengthRule (value object)
│   ├── Min: int
│   └── Max: int
├── CharacterRule (value object)
│   ├── MinUppercase: int
│   ├── MinLowercase: int  
│   ├── MinNumber: int
│   └── MinSymbol: int
└── NoWhitespaceRule (value object)
```

## Business Rules

### Password Generation Rules
1. **Character Distribution**: Required character types must be included
2. **Length Constraints**: Generated passwords must meet min/max length requirements
3. **Randomness**: Each character position must be cryptographically random
4. **Shuffling**: Character order must be randomized after generation

### Validation Rules
1. **Rule Consistency**: Character requirements cannot exceed maximum length
2. **Non-Empty Constraint**: Passwords cannot be empty strings
3. **Character Composition**: Must meet minimum requirements for each character type
4. **Whitespace Policy**: No whitespace allowed when rule is active

### Security Rules
1. **Salt Uniqueness**: Each user must have a unique salt value
2. **Pepper Secrecy**: Pepper values must be kept secret and environment-specific
3. **Hash Strength**: bcrypt cost factor provides adequate computational resistance
4. **Secure Storage**: Only hashed values should be persisted

## Character Classification
- **Uppercase Letters**: A-Z (26 characters)
- **Lowercase Letters**: a-z (26 characters)  
- **Numeric Digits**: 0-9 (10 characters)
- **Symbol Characters**: !@#$%^&*()-_=+[{]};:'\",<.>/? (32 characters)
- **Total Character Space**: 94 possible characters

## API Design Principles

### Builder Pattern Implementation
Provides fluent interface for complex object construction:
- Method chaining for readability
- Immutable final objects
- Validation at build time
- Sensible defaults

### Interface Segregation
Separates concerns through focused interfaces:
- `PaswotHasher`: Hashing responsibility
- `PaswotMatcher`: Verification responsibility
- Clear separation of read/write operations

### Error Handling Strategy
Comprehensive error reporting with:
- Specific error messages for different failure modes
- Validation errors vs runtime errors
- Actionable error information for developers

## Security Architecture

### Cryptographic Foundations
1. **Random Number Generation**: `crypto/rand` provides cryptographically secure randomness
2. **Hash Function**: bcrypt with adaptive cost for future-proofing
3. **Salt Generation**: Should use secure random generation (not provided by library)
4. **Key Derivation**: bcrypt handles key stretching internally

### Attack Resistance
- **Brute Force**: bcrypt computational cost makes attacks expensive
- **Rainbow Tables**: Salt usage prevents precomputed hash attacks  
- **Dictionary Attacks**: Character requirements force complex passwords
- **Timing Attacks**: bcrypt comparison provides constant-time verification

## Performance Characteristics

### Time Complexity
- **Generation**: O(n) where n is password length
- **Validation**: O(n) where n is password length
- **Hashing**: O(2^cost) due to bcrypt work factor
- **Verification**: O(2^cost) due to bcrypt work factor

### Space Complexity
- **Memory Usage**: Minimal allocation during generation
- **Storage Requirements**: Fixed size for bcrypt hashes (60 bytes)

## Integration Patterns

### Recommended Usage Flows

#### User Registration
1. Generate password with appropriate rules
2. Validate password meets policy requirements
3. Hash password with user-specific salt
4. Store hash and salt (never plain password)

#### User Authentication
1. Retrieve stored hash and salt for user
2. Create password object with provided credentials
3. Verify against stored hash using Match method
4. Handle authentication result appropriately

#### Password Policy Updates
1. Define new rules using builder pattern
2. Validate existing passwords against new rules
3. Force password updates for non-compliant passwords
4. Implement gradual rollout for policy changes

## Configuration Management

### Environment-Specific Settings
- **Development**: Lower bcrypt cost for faster tests
- **Production**: Higher bcrypt cost for security
- **Pepper Values**: Different per environment
- **Rule Policies**: May vary by application context

### Default Configurations
The library provides sensible defaults through `DefaultRule()`:
- Balances security with usability
- Meets common compliance requirements
- Can be customized for specific needs

## Testing Strategies

### Unit Testing Approach
- Test each component in isolation
- Mock external dependencies (crypto/rand)
- Verify error conditions and edge cases
- Test rule validation logic thoroughly

### Integration Testing
- Test complete workflows end-to-end
- Verify cryptographic operations produce expected results
- Test with various rule combinations
- Performance testing for acceptable response times

### Security Testing
- Verify randomness quality of generated passwords
- Test resistance to common attack vectors
- Validate proper salt/pepper handling
- Audit cryptographic implementation usage

## Operational Considerations

### Monitoring and Alerting
- Track password generation success rates
- Monitor bcrypt operation performance
- Alert on validation failure spikes
- Log security-relevant events (without sensitive data)

### Scalability Factors
- bcrypt operations are CPU-intensive
- Consider connection pooling for high-load scenarios
- Cache rule validation results where appropriate
- Plan capacity for an authentication load

### Maintenance Requirements
- Regular security audits of dependencies
- Periodic review of password policies
- bcrypt cost factor adjustment as hardware improves
- Monitoring for new security vulnerabilities

This comprehensive context enables AI agents to understand the library's purpose, implementation details, security considerations, and proper usage patterns for any development, analysis, or maintenance tasks.
