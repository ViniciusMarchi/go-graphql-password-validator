// integration tests
package graph

import (
	"graphpass/graph"
	"graphpass/graph/resolver"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"
)

// API response type
type VerifyResult struct {
	Verify  bool
	NoMatch []string
}

type QueryResponse struct {
	Verify VerifyResult
}

// TEST CASE 01: Query with password and rule valid
func TestQueryWithInvalidPasswordAndRule(t *testing.T) {
	c := client.New(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}})))

	query := `{
		verify(
		  password: "TesteSenhaFortee!123&"
		  rules: [
			{rule: "minSize", value: 8},
			{rule: "minSpecialChars", value: 2},
			{rule: "noRepeted", value: 0},
			{rule: "minDigit", value: 4}
		  ]
		) {
		  verify
		  noMatch
		}
	  }
	`
	// make request
	var resp QueryResponse
	c.MustPost(query, &resp)

	// check if response is as expected
	require.False(t, resp.Verify.Verify)
	require.ElementsMatch(t, []string{
		"noRepeted", "minDigit",
	}, resp.Verify.NoMatch)
}

// TEST CASE 02: Query with a password that does not match the stipulated rules
func TestQueryWithInvalidPassword(t *testing.T) {
	c := client.New(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}})))

	query := `{
		verify(
		  password: "TesteSenhaComum4"
		  rules: [
			{rule: "minSize", value: 4},
			{rule: "minSpecialChars", value: 0},
			{rule: "noRepeted", value: 0},
			{rule: "minDigit", value: 1}
		]
		) {
		  verify
		  noMatch
		}
	  }
	`
	var resp QueryResponse
	c.MustPost(query, &resp)

	require.True(t, resp.Verify.Verify)
	require.Empty(t, resp.Verify.NoMatch)
}

// TEST CASE 03: Query with invalid rule
func TestQueryWithInvalidRule(t *testing.T) {
	c := client.New(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}})))

	query := `{
		verify(
		  password: "TesteSenhaForte123&"
		  rules: [{rule: "ruleInvalida", value: 8}]
		) {
		  verify
		  noMatch
		}
	  }
	`
	var resp interface{}
	// expect errors
	require.Panics(t, func() {
		c.MustPost(query, &resp)
	})
}

// TEST CASE 04: Query with invalid rule value (negative value)
func TestQueryWithInvalidValue(t *testing.T) {
	c := client.New(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}})))

	query := `{
		verify(
		  password: "TesteSenhaForte123&"
		  rules: [{rule: "ruleInvalida", value: -1}]
		) {
		  verify
		  noMatch
		}
	  }
	`
	var resp interface{}
	// expect errors
	require.Panics(t, func() {
		c.MustPost(query, &resp)
	})
}
