package utils

import "regexp"

func ValidateUsername(username string) bool {
	// Regular expression to match a string with no spaces and only alphanumeric characters
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)

	return re.MatchString(username)
}
