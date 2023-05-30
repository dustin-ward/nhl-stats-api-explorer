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
			ID: id,
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
			ID: id,
			Name: v_json["name"].(string),
		}
	}

	return venues, nil
}

func GetConferences() (map[int]*model.Conference, error) {
	conferences := make(map[int]*model.Conference, 0)

	data, err := queryNHLStats("/api/v1/conferences")
	if err != nil {
		return conferences, err
	}

	// Parse raw JSON for each conference
	conferences_json := data.(map[string]any)["conferences"]
	for _,v := range conferences_json.([]any) {
		v_json := v.(map[string]any)
		id := (int)(v_json["id"].(float64))
		conferences[id] = &model.Conference{
			ID: id,
			Name: v_json["name"].(string),
			Abbreviation: v_json["abbreviation"].(string),
			NameShort: v_json["shortName"].(string),
			Active: v_json["active"].(bool),
		}
	}

	return conferences, nil
}

func GetDivisions() (map[int]*model.Division, error) {
	divisions := make(map[int]*model.Division, 0)

	data, err := queryNHLStats("/api/v1/divisions")
	if err != nil {
		return divisions, err
	}

	// Parse raw JSON for each division
	divisions_json := data.(map[string]any)["divisions"]
	for _,v := range divisions_json.([]any) {
		v_json := v.(map[string]any)
		id := (int)(v_json["id"].(float64))
		divisions[id] = &model.Division{
			ID: id,
			Name: v_json["name"].(string),
			Abbreviation: v_json["abbreviation"].(string),
			NameShort: v_json["nameShort"].(string),
			ConferenceID: (int)(v_json["conference"].(map[string]any)["id"].(float64)),
			Active: v_json["active"].(bool),
		}
	}

	return divisions, nil
}

func GetPlayer(id int) (*model.Player, error) {
	data, err := queryNHLStats("/api/v1/people/" + strconv.Itoa(id))
	if err != nil {
		return &model.Player{}, nil
	}

	json := data.(map[string]any)["people"].([]any)[0].(map[string]any)
	player := &model.Player{
		ID: id,
		FullName: json["fullName"].(string),
		FirstName: json["firstName"].(string),
		LastName: json["lastName"].(string),
		PrimaryNumber: json["primaryNumber"].(string),
		CurrentAge: (int)(json["currentAge"].(float64)),
		BirthDate: json["birthDate"].(string),
		Nationality: json["nationality"].(string),
		Height: json["height"].(string),
		Weight: (int)(json["weight"].(float64)),
		Active: json["active"].(bool),
		Rookie: json["rookie"].(bool),
		RosterStatus: json["rosterStatus"].(string),
		Position: &model.Position{
			Code: json["primaryPosition"].(map[string]any)["code"].(string),
			Name: json["primaryPosition"].(map[string]any)["name"].(string),
			Type: json["primaryPosition"].(map[string]any)["type"].(string),
			Abbreviation: json["primaryPosition"].(map[string]any)["abbreviation"].(string),
		},
	}

	// Annoying type conversions cant be done inline?
	b_city := json["birthCity"].(string)
	player.BirthCity = &b_city
	b_sp := json["birthStateProvince"].(string)
	player.BirthStateProvince = &b_sp
	b_cnt := json["birthCountry"].(string)
	player.BirthCountry = &b_cnt
	cptn := json["captain"].(bool)
	player.Captain = &cptn
	alt := json["alternateCaptain"].(bool)
	player.AlternateCaptain = &alt

	// Can be nil
	if json["currentTeam"] != nil {
		t_id := (int)(json["currentTeam"].(map[string]any)["id"].(float64))
		player.CurrentTeam = &t_id
	}

	return player, nil
}
