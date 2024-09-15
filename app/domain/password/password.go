package password

//go:generate go run github.com/golang/mock/mockgen -package passwordtest -destination passwordtest/mock_password.go github.com/shipyardapp/Shipyard/app/internal/lib/password Passworder

import (
	"errors"
	"regexp"

	"github.com/sethvargo/go-password/password"
)

// ErrPasswordNotStrongEnough is returned if the Passworder implementation
// finds the provided password does not meet the minimum strength requirements.
var ErrPasswordNotStrongEnough = errors.New("password: not strong enough")

// Passworder is an interface that provides hashing and comparing plaintext passwords.
// The returned hash from Hash() should be able to input to Equal without any modification.
//
// Implementations should be cryptographically secure.
type Passworder interface {
	// Hash should return a hashed version of password and a non-nil err if it is
	// not able to do so.
	// The return parameter hash should be encoded in whatever means necessary to
	// properly serialize and receive back into Equal() for verification.
	//
	// Note that there is no salt parameter. Implementations should not require
	// an explicit salt for their algorithms.
	Hash(password string) (hash string, err error)

	// Equal should return whether or not hash is a hashed version of password.
	//
	// Note that implementations should do so in a way that avoids timing attacks.
	Equal(password string, hash string) bool

	// EvaluateAndHash confirms that the password and confirmPassword inputs are identical,
	// that password meets the minimum strength requirements, and returns the hashed
	// password value on success.
	//
	// Note that there is no salt parameter. Hash implementations should not require an
	// explicit salt for their algorithms.
	//
	// Note that implementations for equal should do so in a way that avoids timing attacks.
	EvaluateAndHash(password, confirmPassword string) ([]byte, error)
}

// Default is a Passworder that can be set for package use.
var Default Passworder

var (
	containsUpperRegexp        = regexp.MustCompile("[[:upper:]]")
	containsLowerRegexp        = regexp.MustCompile("[[:lower:]]")
	containsSpecialCharsRegexp = regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	containsNumbersRegexp      = regexp.MustCompile(`\d`)
)

// ValidatePasswordStrength determines whether or not pw is strong enough.
// It returns ErrPasswordNotStrongEnough if pw is not stong enough or nil otherwise.
func ValidatePasswordStrength(pw string) error {
	pwOK := 0

	if containsUpperRegexp.MatchString(pw) {
		pwOK++
	}
	if containsLowerRegexp.MatchString(pw) {
		pwOK++
	}
	if containsSpecialCharsRegexp.MatchString(pw) {
		pwOK++
	}
	if containsNumbersRegexp.MatchString(pw) {
		pwOK++
	}

	hasConsecutiveCharactersCheck := hasConsecutiveCharacters(pw)

	if len(pw) < 12 ||
		pwOK < 3 || hasConsecutiveCharactersCheck {
		return ErrPasswordNotStrongEnough
	}

	return nil
}

func hasConsecutiveCharacters(input string) bool {
	count := 1
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
			if count > 2 {
				return true
			}
		} else {
			count = 1
		}
	}
	return false
}

func GenerateStrongRandomPassword() (string, error) {
	return password.Generate(32, 10, 10, false, false)
}
