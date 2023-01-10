// unit tests to process of transforming the type []map[string]interface{} to struct
package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// CASE 01: valid input
func TestMapToStructWithValidInput(t *testing.T) {
	rulesMap := []map[string]interface{}{
		{"rule": "minSize", "value": int64(10)},
		{"rule": "minUppercase", "value": int64(3)},
		{"rule": "minLowercase", "value": int64(2)},
		{"rule": "minDigit", "value": int64(1)},
	}
	expectedRulesStruct := []Rule{
		{Rule: "minSize", Value: 10},
		{Rule: "minUppercase", Value: 3},
		{Rule: "minLowercase", Value: 2},
		{Rule: "minDigit", Value: 1},
	}

	rulesStruct, err := MapToStruct(rulesMap)

	// in this case no error should be thrown
	assert.Nil(t, err, "MapToStruct returned an unexpected error, even with valid input.")
	assert.Equal(t, expectedRulesStruct, rulesStruct,
		"MapToStruct returned an incorrect output, even with valid input.",
	)
}

// CASE 02: invalid rule
func TestMapToStructWithInvalidRule(t *testing.T) {
	rulesMap := []map[string]interface{}{
		{"rule": "invalidRule", "value": int64(10)},
	}

	_, err := MapToStruct(rulesMap)

	// a error must have been thrown
	assert.NotNil(t, err, "MapToStruct did not return an error, even with an invalid rule.")
	expectedErrorMsg := fmt.Sprintf("the rule '%s' is invalid. List of accepted rules: %v", rulesMap[0]["rule"], acceptedRules)
	assert.Equal(t, expectedErrorMsg, err.Error())
}

// CASE 03: invalid rule value
func TestMapToStructInvalidValue(t *testing.T) {
	rulesMap := []map[string]interface{}{
		{"rule": "minSize", "value": int64(-10)},
	}

	_, err := MapToStruct(rulesMap)

	assert.NotNil(t, err, "MapToStruct did not return an error, even with an negative rule value.")
	expectedErrorValueMsg := fmt.Sprintf("the value %d of the rule '%s' is invalid. Negative values are not accepted", rulesMap[0]["value"], rulesMap[0]["rule"])
	assert.Equal(t, expectedErrorValueMsg, err.Error())
}
