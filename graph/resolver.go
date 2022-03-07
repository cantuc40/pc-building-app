//go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	"github.com/cantuc40/gqlgen-todos/graph/model"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users         []*model.User
	motherboards  []*model.Motherboard
	OSs           []*model.OperatingSystem
	cpus          []*model.CPU
	storages      []*model.Storage
	memories      []*model.Memory
	powersupplies []*model.PowerSupply
	graphiccards  []*model.GraphicsCard
	cases         []*model.Case
	monitors      []*model.Monitor
	//parts model.PartsDb
	DB *gorm.DB
}
