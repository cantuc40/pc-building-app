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

func (r *mutationResolver) RemoveUser(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.User{})
	log.Println("User Deleted")
	return true, nil
}

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

func (r *mutationResolver) RemoveMotherboard(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.Motherboard{})
	log.Println("Motherboard Deleted")
	return true, nil
}

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
	updatedcpu := &model.CPU{
		ID:         rand.Int(),
		Name:       input.Name,
		Company:    input.Company,
		Cores:      input.Cores,
		Clockspeed: input.Clockspeed,
		Sockets:    input.Sockets,
		Price:      input.Price,
	}

	r.DB.Model(&updatedcpu).Updates(&model.CPU{
		Name:       input.Name,
		Company:    input.Company,
		Cores:      input.Cores,
		Clockspeed: input.Clockspeed,
		Sockets:    input.Sockets,
		Price:      input.Price,
	})

	return updatedcpu, nil
}

func (r *mutationResolver) CreateStorage(ctx context.Context, input model.NewStorage) (*model.Storage, error) {
	newstorage := &model.Storage{
		ID:       rand.Int(),
		Name:     input.Name,
		Company:  input.Company,
		Format:   input.Format,
		Capacity: input.Capacity,
		Speed:    input.Speed,
		Price:    input.Price,
	}

	r.DB.Create(&newstorage)
	log.Println("Storage Created")
	return newstorage, nil
}

func (r *mutationResolver) RemoveStorage(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.Storage{})
	log.Println("Storage Deleted")
	return true, nil
}

func (r *mutationResolver) UpdateStorage(ctx context.Context, input model.UpdatedStorage) (*model.Storage, error) {
	newstorage := &model.Storage{
		ID:       input.ID,
		Name:     input.Name,
		Company:  input.Company,
		Format:   input.Format,
		Capacity: input.Capacity,
		Speed:    input.Speed,
		Price:    input.Price,
	}

	r.DB.Model(&newstorage).Updates(&model.Storage{
		Name:     input.Name,
		Company:  input.Company,
		Format:   input.Format,
		Capacity: input.Capacity,
		Speed:    input.Speed,
		Price:    input.Price,
	})

	return newstorage, nil
}

func (r *mutationResolver) CreateMemory(ctx context.Context, input model.NewMemory) (*model.Memory, error) {
	newmemory := &model.Memory{
		ID:             rand.Int(),
		Name:           input.Name,
		Company:        input.Company,
		Frequency:      input.Frequency,
		CasLatency:     input.CasLatency,
		MemoryChannels: input.MemoryChannels,
		Price:          input.Price,
	}

	r.DB.Create(&newmemory)
	return newmemory, nil
}

func (r *mutationResolver) RemoveMemory(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.Memory{})
	log.Println("Memory Deleted")
	return true, nil
}

func (r *mutationResolver) UpdateMemory(ctx context.Context, input model.UpdatedMemory) (*model.Memory, error) {
	updatememory := &model.Memory{
		ID:             input.ID,
		Name:           input.Name,
		Company:        input.Company,
		Frequency:      input.Frequency,
		CasLatency:     input.CasLatency,
		MemoryChannels: input.MemoryChannels,
		Price:          input.Price,
	}

	r.DB.Model(&updatememory).Updates(&model.Memory{
		Name:           input.Name,
		Company:        input.Company,
		Frequency:      input.Frequency,
		CasLatency:     input.CasLatency,
		MemoryChannels: input.MemoryChannels,
		Price:          input.Price,
	})

	return updatememory, nil
}

func (r *mutationResolver) CreatePowerSupply(ctx context.Context, input model.NewPowerSupply) (*model.PowerSupply, error) {
	newpowersupply := &model.PowerSupply{
		ID:         rand.Int(),
		Name:       input.Name,
		Company:    input.Company,
		FormFactor: input.FormFactor,
		Modularity: input.Modularity,
		Price:      input.Price,
	}

	r.DB.Create(&newpowersupply)
	return newpowersupply, nil
}

func (r *mutationResolver) RemovePowerSupply(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.PowerSupply{})
	log.Println("Power Supply Deleted")
	return true, nil
}

func (r *mutationResolver) UpdatePowerSupply(ctx context.Context, input model.UpdatedPowerSupply) (*model.PowerSupply, error) {
	updatepowersupply := &model.PowerSupply{
		ID:         input.ID,
		Name:       input.Name,
		Company:    input.Company,
		FormFactor: input.FormFactor,
		Modularity: input.Modularity,
		Price:      input.Price,
	}

	r.DB.Model(&updatepowersupply).Updates(&model.PowerSupply{
		Name:       input.Name,
		Company:    input.Company,
		FormFactor: input.FormFactor,
		Modularity: input.Modularity,
		Price:      input.Price,
	})

	return updatepowersupply, nil
}

func (r *mutationResolver) CreateGraphicCard(ctx context.Context, input model.NewGraphicCard) (*model.GraphicsCard, error) {
	newgraphicscard := &model.GraphicsCard{
		ID:        rand.Int(),
		Name:      input.Name,
		Company:   input.Company,
		Gpu:       input.Gpu,
		Memory:    input.Memory,
		CoreClock: input.CoreClock,
		Price:     input.Price,
	}

	r.DB.Create(&newgraphicscard)
	return newgraphicscard, nil
}

func (r *mutationResolver) RemoveGraphicCard(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.GraphicsCard{})
	log.Println("Graphics Card Deleted")
	return true, nil
}

func (r *mutationResolver) UpdateGraphicCard(ctx context.Context, input model.UpdatedGraphicCard) (*model.GraphicsCard, error) {
	updategraphicscard := &model.GraphicsCard{
		ID:        input.ID,
		Name:      input.Name,
		Company:   input.Company,
		Gpu:       input.Gpu,
		Memory:    input.Memory,
		CoreClock: input.CoreClock,
		Price:     input.Price,
	}

	r.DB.Model(&updategraphicscard).Updates(&model.GraphicsCard{
		Name:      input.Name,
		Company:   input.Company,
		Gpu:       input.Gpu,
		Memory:    input.Memory,
		CoreClock: input.CoreClock,
		Price:     input.Price,
	})

	return updategraphicscard, nil
}

func (r *mutationResolver) CreateCase(ctx context.Context, input model.NewCase) (*model.Case, error) {
	newcase := &model.Case{
		ID:       rand.Int(),
		Name:     input.Name,
		Company:  input.Company,
		Windowed: input.Windowed,
		Material: input.Material,
		Price:    input.Price,
	}

	r.DB.Create(&newcase)
	return newcase, nil
}

func (r *mutationResolver) RemoveCase(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.Case{})
	log.Println("Case Deleted")
	return true, nil
}

func (r *mutationResolver) UpdateCase(ctx context.Context, input model.UpdatedCase) (*model.Case, error) {
	updatecase := &model.Case{
		ID:       input.ID,
		Name:     input.Name,
		Company:  input.Company,
		Windowed: input.Windowed,
		Material: input.Material,
		Price:    input.Price,
	}

	r.DB.Model(&updatecase).Updates(&model.Case{
		Name:     input.Name,
		Company:  input.Company,
		Windowed: input.Windowed,
		Material: input.Material,
		Price:    input.Price,
	})

	return updatecase, nil
}

func (r *mutationResolver) CreateMonitor(ctx context.Context, input model.NewMonitor) (*model.Monitor, error) {
	newmonitor := &model.Monitor{
		ID:         rand.Int(),
		Name:       input.Name,
		Company:    input.Company,
		Resolution: input.Resolution,
		Hz:         input.Hz,
		Price:      input.Price,
	}

	r.DB.Create(&newmonitor)
	return newmonitor, nil
}

func (r *mutationResolver) RemoveMonitor(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.Monitor{})
	log.Println("CPU Deleted")
	return true, nil
}

func (r *mutationResolver) UpdateMonitor(ctx context.Context, input model.UpdatedMonitor) (*model.Monitor, error) {
	updatemonitor := &model.Monitor{
		ID:         input.ID,
		Name:       input.Name,
		Company:    input.Company,
		Resolution: input.Resolution,
		Hz:         input.Hz,
		Price:      input.Price,
	}

	r.DB.Model(&updatemonitor).Updates(&model.Monitor{
		Name:       input.Name,
		Company:    input.Company,
		Resolution: input.Resolution,
		Hz:         input.Hz,
		Price:      input.Price,
	})
	return updatemonitor, nil
}

func (r *mutationResolver) CreateOperatingSystem(ctx context.Context, input model.NewOs) (*model.OperatingSystem, error) {
	newOS := &model.OperatingSystem{
		ID:      rand.Int(),
		Name:    input.Name,
		Company: input.Company,
		Price:   int(input.Price),
	}

	r.DB.Create(&newOS)
	r.OSs = append(r.OSs, newOS)
	return newOS, nil
}

func (r *mutationResolver) RemoveOperatingSystem(ctx context.Context, input int) (bool, error) {
	r.DB.Where("id = ?", input).Delete(&model.OperatingSystem{})
	log.Println("CPU Deleted")
	return true, nil
}

func (r *mutationResolver) UpdateOperatingSystem(ctx context.Context, input model.UpdatedOs) (*model.OperatingSystem, error) {
	updateOS := &model.OperatingSystem{
		ID:      input.ID,
		Name:    input.Name,
		Company: input.Company,
		Price:   input.Price,
	}

	r.DB.Model(&updateOS).Updates(&model.OperatingSystem{
		Name:    input.Name,
		Company: input.Company,
		Price:   input.Price,
	})
	return updateOS, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Motherboards(ctx context.Context) ([]*model.Motherboard, error) {
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
	return r.OSs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Motherboard(ctx context.Context) ([]*model.Motherboard, error) {
	panic(fmt.Errorf("not implemented"))
}
