package password

import (
	"graphpass/utils"
	"regexp"
	"strings"
)

// counts the number of uppercase characters in a string
func countUppercaseChars(password string) int {
	r := regexp.MustCompile("[A-Z]")
	return len(r.FindAllString(password, -1))
}

// counts the number of lowercase characters in a string
func countLowerCaseChars(password string) int {
	r := regexp.MustCompile("[a-z]")
	return len(r.FindAllString(password, -1))
}

// counts the number of digits in a string
func countDigits(password string) int {
	r := regexp.MustCompile("[0-9]")
	return len(r.FindAllString(password, -1))
}

// counts the number of special characters in a string
func countSpecialChars(password string) int {
	r := regexp.MustCompile(`[!@#$%^&*()-+\/{}[]`)
	return len(r.FindAllString(password, -1))
}

// Check if a string has sequential repeating characters
// returns true if there is repetition, false if there is no repetition
func isRepeat(password string) bool {
	var prevChar string // stores the previously visited character

	for _, char := range strings.Split(password, "") {
		// if the previous character is the same as the current character then we have found sequential repetition
		if prevChar == char {
			return true
		}
		prevChar = char
	}
	return false
}

// checks if the password has the minimum length stipulated by the user
func minSize(password string, threshold int) bool {
	return len(password) >= threshold
}

// checks if the password has the minimum amount of uppercase characters defined by the user
func minUpperCase(password string, threshold int) bool {
	return countUppercaseChars(password) >= threshold
}

// checks if the password has the minimum amount of lowercase characters defined by the user
func minLowerCase(password string, threshold int) bool {
	return countLowerCaseChars(password) >= threshold
}

// checks if the password has the minimum amount of digits defined by the user
func minDigit(password string, threshold int) bool {
	return countDigits(password) >= threshold
}

// checks if the password has the minimum amount of special characters defined by the user
func minSpecialChars(password string, threshold int) bool {
	return countSpecialChars(password) >= threshold
}

// This function checks for sequential repetition in the password. It returns true if there
// are no repeated characters in the password, and false otherwise. In other words, true
// indicates a valid password (without sequential repetition) and false indicates an invalid password.
func noRepeted(password string, value int) bool {
	return !isRepeat(password)
}

// The ValidPassword function verifies whether a given password adheres to all rules specified by the user.
// It returns a boolean indicating whether the password is valid or not, and a list of error messages detailing
// any rules that the password failed to meet.
func ValidPassword(password string, rules []utils.Rule) (bool, []string) {
	noMatched := make([]string, 0)
	var validPassword bool = true

	// The user can choose from a set of predefined password rules. By the time this data reaches this
	// function, it has already been validated to ensure that the chosen rules are among the allowed rules.
	// Thus, to cater to different set of rules, this implementation utilizes dynamic function execution
	// technique. To facilitate this, a map structure is used, which indexes the functions by their names.
	mappedFunc := map[string]func(string, int) bool{
		"minSize":         minSize,
		"minUppercase":    minUpperCase,
		"minLowercase":    minLowerCase,
		"minDigit":        minDigit,
		"minSpecialChars": minSpecialChars,
		"noRepeted":       noRepeted,
	}

	for _, m := range rules {
		rule := m.Rule
		value := m.Value

		result := mappedFunc[rule](password, value) // run function dinamically
		// if the result is false, we know that the rule has not been matched
		if !result {
			// put the no matched rule in a slice, to return to user
			noMatched = append(noMatched, rule)
		}
	}
	// if the noMatched slice are empty the password is valid
	if len(noMatched) > 0 {
		validPassword = false
	}
	return validPassword, noMatched
}
