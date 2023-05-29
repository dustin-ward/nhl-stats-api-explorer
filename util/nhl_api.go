package util

import (
	"dustin-ward/NHLGraphQLAPI/graph/model"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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

func GetTeams() (map[int]*model.Team, error) {
	teams := make(map[int]*model.Team, 0)

	data, err := queryNHLStats("/api/v1/teams")
	if err != nil {
		return teams, err
	}

	// Parse raw JSON for each team
	teams_json := data.(map[string]any)["teams"]
	for _,t := range teams_json.([]any) {
		t_json := t.(map[string]any)
		id := (int)(t_json["id"].(float64))
		teams[id] = &model.Team{
			ID: (int)(t_json["id"].(float64)),
			Name: t_json["name"].(string),
			TeamName: t_json["teamName"].(string),
			LocationName: t_json["locationName"].(string),
			ShortName: t_json["shortName"].(string),
			Abbreviation: t_json["abbreviation"].(string),
			OfficialSiteURL: t_json["officialSiteUrl"].(string),
			VenueID: -1,
			DivisionID: (int)(t_json["division"].(map[string]any)["id"].(float64)),
			ConferenceID: (int)(t_json["conference"].(map[string]any)["id"].(float64)),
			FranchiseID: (int)(t_json["franchise"].(map[string]any)["franchiseId"].(float64)),
			Active: t_json["active"].(bool),
		}

		// Annoyingly stored as a string instead of int...
		teams[id].FirstYearOfPlay, _ = strconv.Atoi(t_json["firstYearOfPlay"].(string))

		// Venue id can be nil
		v := t_json["venue"].(map[string]any)["id"]
		if v != nil {
			teams[id].VenueID = (int)(v.(float64))
		}
	}

	return teams, nil
}

func GetVenues() (map[int]*model.Venue, error) {
	venues := make(map[int]*model.Venue, 0)

	data, err := queryNHLStats("/api/v1/venues")
	if err != nil {
		return venues, err
	}

	// Parse raw JSON for each venue
	venues_json := data.(map[string]any)["venues"]
	for _,v := range venues_json.([]any) {
		v_json := v.(map[string]any)
		id := (int)(v_json["id"].(float64))
		venues[id] = &model.Venue{
			ID: (int)(v_json["id"].(float64)),
			Name: v_json["name"].(string),
		}
	}

	return venues, nil
}
