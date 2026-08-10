package pkg

import (
	"fmt"
	"regexp"
)

var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

func IsValidEmail(email string) bool {

	return emailRegex.MatchString(email)

}

func IsValidPassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	hasUpperCase := false
	hasLowerCase := false
	hasNumber := false


	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpperCase = true
		case 'a' <= char && char <= 'z':
			hasLowerCase = true
		case '0' <= char && char <= '9':
			hasNumber = true
		}
	}

	if !hasUpperCase {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !hasLowerCase {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return fmt.Errorf("password must contain at least one number")
	}

	return nil
}
