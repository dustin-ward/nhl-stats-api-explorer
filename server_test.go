package main

import (
	"dustin-ward/NHLGraphQLAPI/util"
	"fmt"
	"testing"
)

func TestGetTeams(t *testing.T) {
	teams, err := util.GetTeams()
	if err != nil {
		t.Error(err)
	}

	for _,t := range teams {
		fmt.Println("=================================")
		fmt.Println(t.Name)
		fmt.Println(t.Abbreviation)
	}

	if len(teams) != 32 {
		t.Errorf("wanted %d teams, got %d", 32, len(teams))
	}
}
