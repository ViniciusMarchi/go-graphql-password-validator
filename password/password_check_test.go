// unit tests to the password validation process

package password

import (
	"graphpass/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests the count of uppercase characters in a string
func TestCountUppercaseChars(t *testing.T) {
	tests := []struct {
		password_input string
		want_output    int
	}{
		{password_input: "aaaAAaaBccD", want_output: 4},
		{password_input: "aaaAAa aBc cD", want_output: 4},
		{password_input: "aaa2612$", want_output: 0},
		{password_input: "", want_output: 0},
		{password_input: "A1234#aB", want_output: 2},
		{password_input: "    ", want_output: 0},
	}

	for _, test := range tests {
		result := countUppercaseChars(test.password_input)
		assert.Equal(t, test.want_output, result,
			"Test of verification of password '%s' failed: it was expected that "+
				"the number of uppercase characters would be %v, but it is %v",
			test.password_input, test.want_output, result,
		)
	}
}

// Tests the count of lowercase characters in a string
func TestCountLowerChars(t *testing.T) {
	tests := []struct {
		password_input string
		want_output    int
	}{
		{password_input: "aaaAAaaBccD", want_output: 7},
		{password_input: "aaaAAa aBc cD", want_output: 7},
		{password_input: "aaa2612$", want_output: 3},
		{password_input: "", want_output: 0},
		{password_input: "A1234#B", want_output: 0},
		{password_input: "    ", want_output: 0},
	}

	for _, test := range tests {
		result := countLowerCaseChars(test.password_input)
		assert.Equal(t, test.want_output, result,
			"Test of verification of password '%s' failed: it was expected that "+
				"the number of lowercase characters would be %v, but it is %v",
			test.password_input, test.want_output, result,
		)
	}
}

// Tests the count of digits (0-9) in a string
func TestCountDigits(t *testing.T) {
	tests := []struct {
		password_input string
		want_output    int
	}{
		{password_input: "aaaAAaaBccD", want_output: 0},
		{password_input: "aaaAAa aBc cD", want_output: 0},
		{password_input: "aaa2612$", want_output: 4},
		{password_input: "aa=a2-6 1 2 $", want_output: 4},
		{password_input: "", want_output: 0},
	}

	for _, test := range tests {
		result := countDigits(test.password_input)

		assert.Equal(t, test.want_output, result,
			"Test of verification of password '%s' failed: it was expected that "+
				"the number of digits would be %v, but it is %v",
			test.password_input, test.want_output, result,
		)
	}
}

// Tests the count of special characters in a string
func TestCountSpecialChars(t *testing.T) {
	tests := []struct {
		password_input string
		want_output    int
	}{
		{password_input: "aaaAAaaBccD", want_output: 0},
		{password_input: "aaaAAa aBc cD", want_output: 0},
		{password_input: "aaa2!612$", want_output: 2},
		{password_input: "senha@ # @$", want_output: 4},
	}

	for _, test := range tests {
		result := countSpecialChars(test.password_input)
		assert.Equal(t, test.want_output, result,
			"Test of verification of password '%s' failed: it was expected that "+
				"the number of special characters would be %v, but it is %v",
			test.password_input, test.want_output, result,
		)
	}
}

// Tests if a string has sequential repeating characters
func TestIsRepeat(t *testing.T) {
	tests := []struct {
		password_input string
		want_output    bool
	}{
		{password_input: "aaaAAaaBccD", want_output: true},
		{password_input: "aaaAAa aBc cD", want_output: true},
		{password_input: "a2!612$", want_output: false},
		{password_input: "senha@ # @$", want_output: false},
	}

	for _, test := range tests {
		result := isRepeat(test.password_input)
		assert.EqualValues(t, test.want_output, result,
			"Test of verification of password '%s' failed: it was expected that "+
				"there were no sequential repeated characters, but there are.",
			test.password_input,
		)
		assert.EqualValues(t, test.want_output, result)
	}
}

// Struct for defining a type for the inputs of tests that verify
// if a string meets a certain minimum requirement
type caseTestMinsFormat struct {
	password_input string
	threshold      int
	want_output    bool
}

// Tests if a string has a minimum length
func TestMinSize(t *testing.T) {
	tests := []caseTestMinsFormat{
		{password_input: "aaaAAaaBccD", threshold: 4, want_output: true},
		{password_input: "461ada616", threshold: 4, want_output: true},
		{password_input: "aaa", threshold: 3, want_output: true},
		{password_input: "", threshold: 4, want_output: false},
	}
	for _, test := range tests {
		result := minSize(test.password_input, test.threshold)
		assert.EqualValues(t, test.want_output, result,
			"Test of verification of password '%s' failed: it was expected that "+
				"it had the minimum length of %v characters.",
			test.password_input, test.threshold,
		)
	}
}

// Tests if a string has a minimum number of uppercase characters
func TestMinUppercase(t *testing.T) {
	tests := []caseTestMinsFormat{
		{password_input: "aaaAAaaBccD", threshold: 4, want_output: true},
		{password_input: "461ada616", threshold: 4, want_output: false},
		{password_input: "aaa", threshold: 0, want_output: true},
	}
	for _, test := range tests {
		result := minUpperCase(test.password_input, test.threshold)
		assert.EqualValues(t, test.want_output, result,
			"Test of verification of password '%s' failed: it was expected that "+
				"it had the minimum number of  %v uppercase letters.",
			test.password_input, test.threshold,
		)
	}
}

// Tests if a string has a minimum number of lowercase characters
func TestMinLowercase(t *testing.T) {
	tests := []caseTestMinsFormat{
		{password_input: "aaaAAaaBccD", threshold: 4, want_output: true},
		{password_input: "461ada616", threshold: 4, want_output: false},
		{password_input: "aaa", threshold: 4, want_output: false},
		{password_input: "TESTE", threshold: 0, want_output: true},
	}
	for _, test := range tests {
		result := minLowerCase(test.password_input, test.threshold)
		assert.EqualValues(t, test.want_output, result,
			"Test of verification of password '%s' failed: it was expected that "+
				"it had the minimum number of %v lowercase letters.",
			test.password_input, test.threshold,
		)
	}
}

// Tests if a string has a minimum number of special characters
func TestSpecialChars(t *testing.T) {
	tests := []caseTestMinsFormat{
		{password_input: "aaaAAaaBccD", threshold: 1, want_output: false},
		{password_input: "461a#@da616", threshold: 2, want_output: true},
		{password_input: "TESTE", threshold: 0, want_output: true},
	}
	for _, test := range tests {
		result := minSpecialChars(test.password_input, test.threshold)
		assert.EqualValues(t, test.want_output, result,
			"Test of verification of password '%s' failed: it was expected that "+
				"it had the minimum number of %v special characters.",
			test.password_input, test.threshold,
		)
	}
}

// Tests if a string has no sequential repeated characters
func TestNoRepeat(t *testing.T) {
	tests := []caseTestMinsFormat{
		{password_input: "aaaAAaaBccD", threshold: 0, want_output: false},
		{password_input: "461a#@da616", threshold: 0, want_output: true},
		{password_input: "TESTE", threshold: 0, want_output: true},
	}
	for _, test := range tests {
		result := noRepeted(test.password_input, test.threshold)
		assert.EqualValues(t, test.want_output, result,
			"Test of verification of password '%s' failed: it was expected that "+
				"it had no sequential repeated characters, but there are.",
			test.password_input,
		)
		assert.EqualValues(t, test.want_output, result)
	}
}

// Tests the password validation process
func TestValidPassword(t *testing.T) {
	type caseTestValidPassword struct {
		password          string
		rules             []utils.Rule
		expectedVerify    bool
		expectedNoMatched []string
	}
	tests := []caseTestValidPassword{
		{
			password: "aa16a6aAAaaBviniD",
			rules: []utils.Rule{
				{Rule: "minDigit", Value: 14},
				{Rule: "noRepeted", Value: 0},
			},
			expectedVerify:    false,
			expectedNoMatched: []string{"minDigit", "noRepeted"},
		}, {
			password: "reeepetindocarActEres",
			rules: []utils.Rule{
				{Rule: "minDigit", Value: 14},
			},
			expectedVerify:    false,
			expectedNoMatched: []string{"minDigit"},
		}, {
			password: "bCD3!",
			rules: []utils.Rule{
				{Rule: "minSize", Value: 8},
				{Rule: "minLowercase", Value: 1},
				{Rule: "minUppercase", Value: 1},
				{Rule: "minSpecialChars", Value: 1},
				{Rule: "noRepeted", Value: 0},
			},
			expectedVerify:    false,
			expectedNoMatched: []string{"minSize"},
		},
		{
			password: "abcdefgh",
			rules: []utils.Rule{
				{Rule: "minSize", Value: 8},
				{Rule: "minLowercase", Value: 0},
				{Rule: "minUppercase", Value: 0},
				{Rule: "minSpecialChars", Value: 0},
				{Rule: "noRepeted", Value: 0},
			},
			expectedVerify:    true,
			expectedNoMatched: []string{},
		},
		{
			password: "aAcd$fg1-h",
			rules: []utils.Rule{
				{Rule: "minSize", Value: 8},
				{Rule: "minLowercase", Value: 3},
				{Rule: "minUppercase", Value: 1},
				{Rule: "minSpecialChars", Value: 1},
				{Rule: "noRepeted", Value: 0},
			},
			expectedVerify:    true,
			expectedNoMatched: []string{},
		},
	}

	for _, test := range tests {
		verify, noMatched := ValidPassword(test.password, test.rules)

		assert.Equal(t, test.expectedVerify, verify,
			"Test of verification of password %s failed: it was expected that 'verify' would be %t, but it is %t", test.password, test.expectedVerify, verify)
		assert.Equal(t, test.expectedNoMatched, noMatched,
			"Test of verification of password %s failed: it was expected that 'matched' would be %v, but it is %v", test.password, test.expectedNoMatched, noMatched)
	}
}
