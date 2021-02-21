package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/harshpreet93/next-gqlgen-ent/backend/ent"
	"github.com/harshpreet93/next-gqlgen-ent/backend/graph/generated"
	"github.com/harshpreet93/next-gqlgen-ent/backend/graph/model"
)

func (r *mutationResolver) AddPet(ctx context.Context, pet model.PetInput) (*ent.Pet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddUser(ctx context.Context, name string) (*ent.User, error) {
	return r.DBClient.User.Create().SetName(name).Save(ctx)
}

func (r *mutationResolver) UpdatePet(ctx context.Context, pet model.PetInput) (*ent.Pet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePet(ctx context.Context, userID string, petID string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *pageInfoResolver) StartCursor(ctx context.Context, obj *ent.PageInfo) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *pageInfoResolver) EndCursor(ctx context.Context, obj *ent.PageInfo) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *petResolver) Owner(ctx context.Context, obj *ent.Pet) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*ent.User, error) {
	idAsInt, _ := strconv.Atoi(id)
	return r.DBClient.User.Get(ctx, idAsInt)
}

func (r *queryResolver) GetPet(ctx context.Context, id string) (*ent.Pet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Pets(ctx context.Context, obj *ent.User) ([]*ent.Pet, error) {
	return obj.QueryPet().All(ctx)
}

func (r *userResolver) PetsConnection(ctx context.Context, obj *ent.User, first *int, after *string) (*model.UserPetConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// PageInfo returns generated.PageInfoResolver implementation.
func (r *Resolver) PageInfo() generated.PageInfoResolver { return &pageInfoResolver{r} }

// Pet returns generated.PetResolver implementation.
func (r *Resolver) Pet() generated.PetResolver { return &petResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type pageInfoResolver struct{ *Resolver }
type petResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
