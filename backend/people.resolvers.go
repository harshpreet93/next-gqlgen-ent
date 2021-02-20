package backend

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/harshpreet93/next-gqlgen-ent/backend/ent"
)

func (r *mutationResolver) AddPet(ctx context.Context, pet PetInput) (*ent.Pet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePet(ctx context.Context, pet PetInput) (*ent.Pet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePet(ctx context.Context, userID int, petID int) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *pageInfoResolver) StartCursor(ctx context.Context, obj *ent.PageInfo) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *pageInfoResolver) EndCursor(ctx context.Context, obj *ent.PageInfo) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *petResolver) Owner(ctx context.Context, obj *ent.Pet) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context, id int) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPet(ctx context.Context, id int) (*ent.Pet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Pets(ctx context.Context, obj *ent.User) ([]*ent.Pet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) PetsConnection(ctx context.Context, obj *ent.User, first *int, after *int) (*UserPetConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// PageInfo returns PageInfoResolver implementation.
func (r *Resolver) PageInfo() PageInfoResolver { return &pageInfoResolver{r} }

// Pet returns PetResolver implementation.
func (r *Resolver) Pet() PetResolver { return &petResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type pageInfoResolver struct{ *Resolver }
type petResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
