package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/cantuc40/gqlgen-todos/graph/generated"
	"github.com/cantuc40/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	newuser := &model.User{
		ID:       fmt.Sprintf("T%d", rand.Int()),
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}
	r.users = append(r.users, newuser)
	return newuser, nil
}

func (r *mutationResolver) RemoveUser(ctx context.Context, input model.DeleteUser) (bool, error) {
	//r.DB.Where("id = ?", input.ID).Delete(&model.User{})
	return true, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

func (r *queryResolver) PartsDb(ctx context.Context) (*model.PartsDb, error) {
	return &r.parts, nil
}

func (r *queryResolver) Motherboard(ctx context.Context) ([]*model.Motherboard, error) {
	return r.parts.Motherboards, nil
}

func (r *queryResolver) Cpus(ctx context.Context) ([]*model.CPU, error) {
	return r.parts.Cpus, nil
}

func (r *queryResolver) Storages(ctx context.Context) ([]*model.Storage, error) {
	return r.parts.Storages, nil
}

func (r *queryResolver) Memories(ctx context.Context) ([]*model.Memory, error) {
	return r.parts.Rams, nil
}

func (r *queryResolver) PowerSupplies(ctx context.Context) ([]*model.PowerSupply, error) {
	return r.parts.Powersupplies, nil
}

func (r *queryResolver) GraphicsCards(ctx context.Context) ([]*model.GraphicsCard, error) {
	return r.parts.GraphicCards, nil
}

func (r *queryResolver) Cases(ctx context.Context) ([]*model.Case, error) {
	return r.parts.Cases, nil
}

func (r *queryResolver) Monitors(ctx context.Context) ([]*model.Monitor, error) {
	return r.parts.Monitors, nil
}

func (r *queryResolver) OperatingSystems(ctx context.Context) ([]*model.OperatingSystem, error) {
	return r.parts.OperatingSystems, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Parts(ctx context.Context) (model.PartsDb, error) {
	return r.parts, nil
}
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}
