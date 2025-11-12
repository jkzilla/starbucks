package graph

import (
	"sync"

	"starbucks/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	coffees map[string]*model.Coffee
	mutex   sync.RWMutex
}

func NewResolver() *Resolver {
	return &Resolver{
		coffees: make(map[string]*model.Coffee),
	}
}

func (r *Resolver) AddCoffee(coffee *model.Coffee) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.coffees[coffee.ID] = coffee
}
