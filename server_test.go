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
