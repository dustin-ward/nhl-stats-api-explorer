package main

import (
	"dustin-ward/NHLGraphQLAPI/util"
	"testing"
)

func TestGetTeams(t *testing.T) {
	teams, err := util.GetTeams()
	if err != nil {
		t.Error(err)
		return
	}

	for _,tm := range teams {
		t.Log("=================================")
		t.Logf("Team Name: %s\n", tm.Name)
		t.Logf("Team Abbr: %s\n", tm.Abbreviation)
	}

	if len(teams) != 32 {
		t.Errorf("wanted %d teams, got %d", 32, len(teams))
		return
	}

	const (
		ID = 22
		NAME = "Edmonton Oilers"
		LOCATION = "Edmonton"
		TEAMNAME = "Oilers"
		SHORTNAME = "Edmonton"
		ABBREVIATION = "EDM"
		SITEURL = "http://www.edmontonoilers.com/"
		FIRSTYEAR = 1979
		DIVISION_ID = 15
		CONFERENCE_ID = 5
		VENUE_ID = 5100
		FRANCHISE_ID = 25
		ACTIVE = true
	)

	team, ok := teams[ID]
	if !ok {
		t.Errorf("could not find team id=%d", ID)
		return
	}

	if team.Name != NAME {
		t.Errorf("team.Name != %s, actually: %s", NAME, team.Name)
		return
	}
	if team.LocationName != LOCATION {
		t.Errorf("team.LocationName != %s, actually: %s", LOCATION, team.LocationName)
		return
	}
	if team.TeamName != TEAMNAME {
		t.Errorf("team.TeamName != %s, actually: %s", TEAMNAME, team.TeamName)
		return
	}
	if team.ShortName != SHORTNAME {
		t.Errorf("team.ShortName != %s, actually: %s", SHORTNAME, team.ShortName)
		return
	}
	if team.Abbreviation != ABBREVIATION {
		t.Errorf("team.Abbreviation != %s, actually: %s", ABBREVIATION, team.Abbreviation)
		return
	}
	if team.OfficialSiteURL != SITEURL {
		t.Errorf("team.OfficialSiteURL != %s, actually: %s", SITEURL, team.OfficialSiteURL)
		return
	}
	if team.FirstYearOfPlay != FIRSTYEAR {
		t.Errorf("team.FirstYearOfPlay != %d, actually: %d", FIRSTYEAR, team.FirstYearOfPlay)
		return
	}
	if team.DivisionID != DIVISION_ID {
		t.Errorf("team.DivisionID != %d, actually: %d", DIVISION_ID, team.DivisionID)
		return
	}
	if team.ConferenceID != CONFERENCE_ID {
		t.Errorf("team.ConferenceID != %d, actually: %d", CONFERENCE_ID, team.ConferenceID)
		return
	}
	if team.VenueID != VENUE_ID {
		t.Errorf("team.VenueID != %d, actually: %d", VENUE_ID, team.VenueID)
		return
	}
	if team.FranchiseID != FRANCHISE_ID {
		t.Errorf("team.FranchiseID != %d, actually: %d", FRANCHISE_ID, team.FranchiseID)
		return
	}
	if team.Active != ACTIVE {
		t.Errorf("team.Active != %v, actually: %v", ACTIVE, team.Active)
		return
	}
}

func TestGetVenues(t *testing.T) {
	venues, err := util.GetVenues()
	if err != nil {
		t.Error(err)
		return
	}

	for _,v := range venues {
		t.Log("=================================")
		t.Logf("Venue Name: %s\n", v.Name)
	}

	// For some reason, not all venues exist at this endpoint...
	if len(venues) != 17 {
		t.Errorf("wanted %d venues, got %d", 17, len(venues))
		return
	}

	const (
		ID = 5076
		NAME = "Enterprise Center"
	)

	venue, ok := venues[ID]
	if !ok {
		t.Errorf("could not find venue id=%d", ID)
		return
	}

	if venue.Name != NAME {
		t.Errorf("venue.Name != %s, actually: %s", NAME, venue.Name)
		return
	}
}

func TestGetConferences(t *testing.T) {
	conferences, err := util.GetConferences()
	if err != nil {
		t.Error(err)
		return
	}

	for _,v := range conferences {
		t.Log("=================================")
		t.Logf("Conference Name: %s\n", v.Name)
		t.Logf("Conference Abbr: %s\n", v.Abbreviation)
		t.Logf("Conference Active: %v\n", v.Active)
	}

	if len(conferences) != 2 {
		t.Errorf("wanted %d conferences, got %d", 2, len(conferences))
		return
	}

	const (
		ID = 5
		NAME = "Western"
		ABBREVIATION = "W"
		SHORTNAME = "West"
		ACTIVE = true
	)

	conference, ok := conferences[ID]
	if !ok {
		t.Errorf("could not find conference id=%d", ID)
		return
	}

	if conference.Name != NAME {
		t.Errorf("conference.Name != %s, actually: %s", NAME, conference.Name)
		return
	}
	if conference.Abbreviation != ABBREVIATION {
		t.Errorf("conference.Abbreviation != %s, actually: %s", ABBREVIATION, conference.Abbreviation)
		return
	}
	if conference.NameShort != SHORTNAME {
		t.Errorf("conference.NameShort != %s, actually: %s", SHORTNAME, conference.NameShort)
		return
	}
	if conference.Active != ACTIVE {
		t.Errorf("conference.Active != %v, actually: %v", ACTIVE, conference.Active)
		return
	}
}

func TestGetDivisions(t *testing.T) {
	divisions, err := util.GetDivisions()
	if err != nil {
		t.Error(err)
		return
	}

	for _,v := range divisions {
		t.Log("=================================")
		t.Logf("Division Name: %s\n", v.Name)
		t.Logf("Division Abbr: %s\n", v.Abbreviation)
		t.Logf("Division Active: %v\n", v.Active)
	}

	if len(divisions) != 4 {
		t.Errorf("wanted %d divisions, got %d", 4, len(divisions))
		return
	}

	const (
		ID = 15
		NAME = "Pacific"
		ABBREVIATION = "P"
		SHORTNAME = "PAC"
		CONFID = 5
		ACTIVE = true
	)

	division, ok := divisions[ID]
	if !ok {
		t.Errorf("could not find division id=%d", ID)
		return
	}

	if division.Name != NAME {
		t.Errorf("division.Name != %s, actually: %s", NAME, division.Name)
		return
	}
	if division.Abbreviation != ABBREVIATION {
		t.Errorf("division.Abbreviation != %s, actually: %s", ABBREVIATION, division.Abbreviation)
		return
	}
	if division.NameShort != SHORTNAME {
		t.Errorf("division.NameShort != %s, actually: %s", SHORTNAME, division.NameShort)
		return
	}
	if division.ConferenceID != CONFID {
		t.Errorf("division.ConferenceID != %d, actually: %d", CONFID, division.ConferenceID)
		return
	}
	if division.Active != ACTIVE {
		t.Errorf("division.Active != %v, actually: %v", ACTIVE, division.Active)
		return
	}
}

func TestGetPlayer(t *testing.T) {
	const (
		ID = 8476454
		FULLNAME = "Ryan Nugent-Hopkins"
		FIRSTNAME = "Ryan"
		LASTNAME = "Nugent-Hopkins"
		NUMBER = "93"
		BIRTHDATE = "1993-04-12"
		CURAGE = 30
		BIRTHCITY = "Burnaby"
		BIRTHPROVINCE = "BC"
		BIRTHCOUNTRY = "CAN"
		NATIONALITY = "CAN"
		HEIGHT = "6' 0\""
		WEIGHT = 195
		ACTIVE = true
		CAPTAIN = false
		ALTERNATE = true
		ROOKIE = false
		ROSTER = "Y"
		SHOOTS = "L"
		CUR_TEAM = 22
		POS_CODE = "C"
		POS_NAME = "Center"
		POS_TYPE = "Forward"
		POS_ABBR = "C"
	)

	player, err := util.GetPlayer(ID)
	if err != nil {
		t.Errorf("unable to get player %d", ID)
	}

	if player.FullName != FULLNAME {
		t.Errorf("player.FullName != %s, actually: %s", player.FullName, FULLNAME)
	}
	if player.FirstName != FIRSTNAME {
		t.Errorf("player.FirstName != %s, actually: %s", player.FirstName, FIRSTNAME)
	}
	if player.LastName != LASTNAME {
		t.Errorf("player.LastName != %s, actually: %s", player.LastName, LASTNAME)
	}
	if player.CurrentAge != CURAGE {
		t.Errorf("player.CurrentAge != %d, actually: %d", player.CurrentAge, CURAGE)
	}
	if player.PrimaryNumber != NUMBER {
		t.Errorf("player.PrimaryNumber != %s, actually: %s", player.PrimaryNumber, NUMBER)
	}
	if player.BirthDate != BIRTHDATE {
		t.Errorf("player.BirthDate != %s, actually: %s", player.BirthDate, BIRTHDATE)
	}
	if *player.BirthCity != BIRTHCITY {
		t.Errorf("player.BirthCity != %s, actually: %s", *player.BirthCity, BIRTHCITY)
	}
	if *player.BirthStateProvince != BIRTHPROVINCE {
		t.Errorf("player.BirthStateProvince != %s, actually: %s", *player.BirthStateProvince, BIRTHPROVINCE)
	}
	if *player.BirthCountry != BIRTHCOUNTRY {
		t.Errorf("player.BirthCountry != %s, actually: %s", *player.BirthCountry, BIRTHCOUNTRY)
	}
	if player.Nationality != NATIONALITY {
		t.Errorf("player.Nationality != %s, actually: %s", player.Nationality, NATIONALITY)
	}
	if player.Height != HEIGHT {
		t.Errorf("player.Height != %s, actually: %s", player.Height, HEIGHT)
	}
	if player.Weight != WEIGHT {
		t.Errorf("player.Weight != %d, actually: %d", player.Weight, WEIGHT)
	}
	if player.Active != ACTIVE {
		t.Errorf("player.Active != %v, actually: %v", player.Active, ACTIVE)
	}
	if *player.Captain != CAPTAIN {
		t.Errorf("player.Captain != %v, actually: %v", *player.Captain, CAPTAIN)
	}
	if *player.AlternateCaptain != ALTERNATE {
		t.Errorf("player.AlternateCaptain != %v, actually: %v", *player.AlternateCaptain, ALTERNATE)
	}
	if player.Rookie != ROOKIE {
		t.Errorf("player.Rookie != %v, actually: %v", player.Rookie, ROOKIE)
	}
	if player.RosterStatus != ROSTER {
		t.Errorf("player.RosterStatus != %s, actually: %s", player.RosterStatus, ROSTER)
	}
	if *player.CurrentTeam != CUR_TEAM {
		t.Errorf("player.CurrentTeam != %d, actually: %d", *player.CurrentTeam, CUR_TEAM)
	}
	if player.Position.Code != POS_CODE {
		t.Errorf("player.Position.Code != %s, actually: %s", player.Position.Code, POS_CODE)
	}
	if player.Position.Name != POS_NAME {
		t.Errorf("player.Position.Name != %s, actually: %s", player.Position.Name, POS_NAME)
	}
	if player.Position.Type != POS_TYPE {
		t.Errorf("player.Position.Type != %s, actually: %s", player.Position.Type, POS_TYPE)
	}
}
