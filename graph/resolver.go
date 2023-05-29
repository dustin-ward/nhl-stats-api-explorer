package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import "dustin-ward/NHLGraphQLAPI/graph/model"

type Resolver struct {
	teams       map[int]*model.Team
	franchises  map[int]*model.Franchise
	players     map[int]*model.Player
	venues      map[int]*model.Venue
	conferences map[int]*model.Conference
	divisions   map[int]*model.Division
}
