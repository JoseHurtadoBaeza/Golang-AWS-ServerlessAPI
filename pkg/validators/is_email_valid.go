package validators

import "regexp" // Importing regular expression package

// IsEmailValid checks if the provided email string is a valid email address
// email: The email address string to be validated
// Returns true if the email is valid, false otherwise
func IsEmailValid(email string) bool {

	// Define a regular expression pattern to match valid email addresses
	// This regex covers typical email formats, allowing letters, digits, and certain special characters in the local part
	// and enforcing standard rules for the domain part
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]{1,64}@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	// Check the length of the email
	// The length must be between 3 and 254 characters to be considered valid
	// Also, validate the email against the regular expression pattern
	if len(email) < 3 || len(email) > 254 || !rxEmail.MatchString(email) {

		// If the length is invalid or it doesn't match the regex pattern, return false
		return false
	}

	// If the email passes both length and regex checks, return true
	return true

}
