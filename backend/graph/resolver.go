package graph

import "github.com/harshpreet93/next-gqlgen-ent/backend/ent"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver has dependencies that resolvers need
type Resolver struct {
	DBClient *ent.Client
}
