package util

import (
	"dustin-ward/NHLGraphQLAPI/graph/model"
	"encoding/json"
	"io"
	"net/http"
)

const (
	NHL_API_LINK = "https://statsapi.web.nhl.com"
)

func queryNHLStats(endpoint string) (any, error) {
	var data any
	
	resp, err := http.Get(NHL_API_LINK + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetTeams() ([]*model.Team, error) {
	teams := make([]*model.Team, 0)

	r, err := queryNHLStats("/api/v1/teams")
	if err != nil {
		return teams, err
	}
	data := r.(map[string]any)

	teams_json := data["teams"]
	for _,t := range teams_json.([]any) {
		teams = append(teams, &model.Team{
			Name: t.(map[string]any)["name"].(string),
			Abbreviation: t.(map[string]any)["abbreviation"].(string),
		})
	}

	return teams, nil
}
