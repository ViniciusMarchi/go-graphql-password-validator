package utils

import "fmt"

type Rule struct {
	Rule  string
	Value int
}

var acceptedRules = []string{
	"minSize",
	"minUppercase",
	"minLowercase",
	"minDigit",
	"minSpecialChars",
	"noRepeted",
}

// generic helper function that checks if an element exists within a slice of strings
func contains(slice_string []string, elem string) bool {
	for _, item := range slice_string {
		if item == elem {
			return true
		}
	}
	return false
}

// MapToStruct is a helper function that converts the rules received from the user from the format of
// []map[string]interface{} to a struct. The use of this function is to enforce the typing of the received
// rules, which are originally in the format of JSON array with pairs of <key>:<value>. The gqlgen library
// converts the Map scalar type into []map[string]interface{}, which has generic typing for the <value>.
// This function receives this format, converts it into a struct and verifies the validity of the received rules.
// The rules are considered valid if they are within the accepted rules and if the configuration value of the rule
// is positive. This function is also one of the first points of data validation in the API which ensures that the
// next functions that retrieve the data do so in a correct and valid format
func MapToStruct(rules_map []map[string]interface{}) ([]Rule, error) {
	rules_struct := []Rule{}

	for _, rule_item := range rules_map {
		rule := rule_item["rule"].(string)
		value := int(rule_item["value"].(int64)) // by default gqlgen converts the received data to int64

		if value < 0 {
			return nil, fmt.Errorf("the value %d of the rule '%s' is invalid. Negative values are not accepted", value, rule)
		}

		if !contains(acceptedRules, rule) {
			return nil, fmt.Errorf("the rule '%s' is invalid. List of accepted rules: %v", rule, acceptedRules)
		}

		rules_struct = append(rules_struct, Rule{
			Rule:  rule,
			Value: value,
		})
	}
	return rules_struct, nil
}
