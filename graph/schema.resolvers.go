package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/cantuc40/gqlgen-todos/graph/generated"
	"github.com/cantuc40/gqlgen-todos/graph/model"
)

//------------------------User Mutations-------------------------
//Create a New User
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	newuser := &model.User{
		ID:       rand.Int(),
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}

	r.DB.Create(&newuser)
	r.users = append(r.users, newuser) //Used for development enviroment
	log.Println("User added")
	return newuser, nil
}

//Remove a User
func (r *mutationResolver) RemoveUser(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.User{})
	log.Println("User Deleted")
	return true, nil
}

//Update a User
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdatedUser) (*model.User, error) {
	updateduser := &model.User{
		ID:       input.ID,
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}

	r.DB.Model(&updateduser).Updates(&model.User{Username: input.Username, Password: input.Password, Email: input.Email})
	return updateduser, nil
}

//------------------------Motherboard Mutations-------------------------
//Create Motherboard
func (r *mutationResolver) CreateMotherboard(ctx context.Context, input model.NewMotherBoard) (*model.Motherboard, error) {
	newmotherboard := &model.Motherboard{
		ID:            rand.Int(),
		Name:          input.Name,
		Company:       input.Company,
		FormFactor:    input.FormFactor,
		Sockets:       input.Sockets,
		Chipsets:      input.Chipsets,
		PciSlots:      input.PciSlots,
		RAMSlots:      input.RAMSlots,
		InternalPorts: input.InternalPorts,
		ExternalPorts: input.ExternalPorts,
		Price:         input.Price,
	}

	r.DB.Create(&newmotherboard)
	log.Println("Motherboard added")
	return newmotherboard, nil
}

//Remove Motherboard
func (r *mutationResolver) RemoveMotherboard(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.Motherboard{})
	log.Println("Motherboard Deleted")
	return true, nil
}

//Update Motherboard
func (r *mutationResolver) UpdateMotherboard(ctx context.Context, input *model.UpdatedMotherboard) (*model.Motherboard, error) {
	updatedmotherboard := &model.Motherboard{
		ID:            input.ID,
		Name:          input.Name,
		Company:       input.Company,
		FormFactor:    input.FormFactor,
		Sockets:       input.Sockets,
		Chipsets:      input.Chipsets,
		PciSlots:      input.PciSlots,
		RAMSlots:      input.RAMSlots,
		InternalPorts: input.InternalPorts,
		ExternalPorts: input.ExternalPorts,
		Price:         input.Price,
	}

	r.DB.Model(&updatedmotherboard).Updates(&model.Motherboard{
		Name:          input.Name,
		Company:       input.Company,
		FormFactor:    input.FormFactor,
		Sockets:       input.Sockets,
		Chipsets:      input.Chipsets,
		PciSlots:      input.PciSlots,
		RAMSlots:      input.RAMSlots,
		InternalPorts: input.InternalPorts,
		ExternalPorts: input.ExternalPorts,
		Price:         input.Price})

	return updatedmotherboard, nil

}

//------------------------CPU Mutations-------------------------
func (r *mutationResolver) CreateCPU(ctx context.Context, input model.NewCPU) (*model.CPU, error) {
	newcpu := &model.CPU{
		ID:         rand.Int(),
		Name:       input.Name,
		Company:    input.Company,
		Cores:      input.Cores,
		Clockspeed: input.Clockspeed,
		Sockets:    input.Sockets,
		Price:      input.Price,
	}
	r.DB.Create(&newcpu)
	log.Println("CPU added")
	return newcpu, nil
}

func (r *mutationResolver) RemoveCPU(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.CPU{})
	log.Println("CPU Deleted")
	return true, nil
}

func (r *mutationResolver) UpdateCPU(ctx context.Context, input model.UpdatedCPU) (*model.CPU, error) {
	updatedcpu = &model.CPU{
		ID:         rand.Int(),
		Name:       input.Name,
		Company:    input.Company,
		Cores:      input.Cores,
		Clockspeed: input.Clockspeed,
		Sockets:    input.Sockets,
		Price:      input.Price,
	}
}

//------------------------Storage Mutations-------------------------
func (r *mutationResolver) CreateStorage(ctx context.Context, input model.NewStorage) (*model.Storage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveStorage(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.Storage{})
	log.Println("Storage Deleted")
	return true, nil
}

func (r *mutationResolver) UpdateStorage(ctx context.Context, input model.UpdatedStorage) (*model.Storage, error) {
	panic(fmt.Errorf("not implemented"))
}

//------------------------Memory Mutations-------------------------
func (r *mutationResolver) CreateMemory(ctx context.Context, input model.NewMemory) (*model.Memory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveMemory(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.Memory{})
	log.Println("Memory Deleted")
	return true, nil
}

func (r *mutationResolver) UpdateMemory(ctx context.Context, input model.UpdatedMemory) (*model.Memory, error) {
	panic(fmt.Errorf("not implemented"))
}

//------------------------Power Supply Mutations-------------------------
func (r *mutationResolver) CreatePowerSupply(ctx context.Context, input model.NewPowerSupply) (*model.PowerSupply, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemovePowerSupply(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.PowerSupply{})
	log.Println("Power Supply Deleted")
	return true, nil
}

func (r *mutationResolver) UpdatePowerSupply(ctx context.Context, input model.UpdatedPowerSupply) (*model.PowerSupply, error) {
	panic(fmt.Errorf("not implemented"))
}

//------------------------Graphics Card Mutations-------------------------
func (r *mutationResolver) CreateGraphicCard(ctx context.Context, input model.NewGraphicCard) (*model.GraphicsCard, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveGraphicCard(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.GraphicsCard{})
	log.Println("Graphics Card Deleted")
	return true, nil
}

func (r *mutationResolver) UpdateGraphicCard(ctx context.Context, input model.UpdatedGraphicCard) (*model.GraphicsCard, error) {
	panic(fmt.Errorf("not implemented"))
}

//------------------------Case Mutations-------------------------
func (r *mutationResolver) CreateCase(ctx context.Context, input model.NewCase) (*model.Case, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveCase(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.Case{})
	log.Println("Case Deleted")
	return true, nil
}

func (r *mutationResolver) UpdateCase(ctx context.Context, input model.UpdatedCase) (*model.Case, error) {
	panic(fmt.Errorf("not implemented"))
}

//------------------------Monitor Mutations-------------------------
func (r *mutationResolver) CreateMonitor(ctx context.Context, input model.NewMonitor) (*model.Monitor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveMonitor(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.Monitor{})
	log.Println("CPU Deleted")
	return true, nil
}

func (r *mutationResolver) UpdateMonitor(ctx context.Context, input model.UpdatedMonitor) (*model.Monitor, error) {
	panic(fmt.Errorf("not implemented"))
}

//------------------------Operating System Mutations-------------------------
func (r *mutationResolver) CreateOperatingSystem(ctx context.Context, input model.NewOs) (*model.OperatingSystem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveOperatingSystem(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.OperatingSystem{})
	log.Println("CPU Deleted")
	return true, nil
}

func (r *mutationResolver) UpdateOperatingSystem(ctx context.Context, input model.UpdatedOs) (*model.OperatingSystem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Motherboard(ctx context.Context) ([]*model.Motherboard, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cpus(ctx context.Context) ([]*model.CPU, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Storages(ctx context.Context) ([]*model.Storage, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Memories(ctx context.Context) ([]*model.Memory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PowerSupplies(ctx context.Context) ([]*model.PowerSupply, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GraphicsCards(ctx context.Context) ([]*model.GraphicsCard, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cases(ctx context.Context) ([]*model.Case, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Monitors(ctx context.Context) ([]*model.Monitor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OperatingSystems(ctx context.Context) ([]*model.OperatingSystem, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
