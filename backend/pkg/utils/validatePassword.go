package utils

import "errors"

func ValidatePassword(password string) error {
	// check if password is empty
	if password == "" {
		return errors.New("password is required")
	}

	// check if password is less than 8 characters
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	// check if password is more than 32 characters
	if len(password) > 32 {
		return errors.New("password must be at most 32 characters")
	}

	// check if password contains at least one uppercase letter
	if !containsUppercaseLetter(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	// check if password contains at least one lowercase letter
	if !containsLowercaseLetter(password) {
		return errors.New("password must contain at least one lowercase letter")
	}

	// check if password contains at least one digit
	if !containsDigit(password) {
		return errors.New("password must contain at least one digit")
	}

	return nil
}

func containsUppercaseLetter(s string) bool {
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			return true
		}
	}
	return false
}

func containsLowercaseLetter(s string) bool {
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			return true
		}
	}
	return false
}

func containsDigit(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}
