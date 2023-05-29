package main

import (
	"dustin-ward/NHLGraphQLAPI/util"
	"fmt"
	"log"
)

func main() {
	teams, err := util.GetTeams()
	if err != nil {
		log.Fatal(err)
	}

	for _,t := range teams {
		fmt.Println("=================================")
		fmt.Println(t.Name)
		fmt.Println(t.Abbreviation)
	}
}
