package resolver

import (
	"context"
	"graphpass/graph"
	"graphpass/graph/model"
	"graphpass/password"
	"graphpass/utils"
)

// The "Verify" function is a resolver that will handle the "verify" query from the user.
// It first maps the user-supplied rules to a struct using the MapToStruct function. Subsequently,
// the entire password validation process is done by the ValidPassword function, and if there are no
// errors, we build the response according to the Password format defined in the schema and return to the user.
func (r *queryResolver) Verify(ctx context.Context, pass string, rules []map[string]interface{}) (*model.Password, error) {
	rules_struct, err := utils.MapToStruct(rules)
	if err != nil {
		return nil, err // if a error occours on MapToStruct, the error is immediately returned to user
	}

	verify, noMatched := password.ValidPassword(pass, rules_struct)

	response := &model.Password{
		Verify:  verify,
		NoMatch: noMatched,
	}
	return response, nil
}

// genered by gqlgen
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
