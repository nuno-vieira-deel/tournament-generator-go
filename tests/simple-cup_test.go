package test

/*
 * Dependencies
 */

import (
	"fmt"
	"testing"

	g "github.com/nunovieira220/tournament-generator-go/generators"
	m "github.com/nunovieira220/tournament-generator-go/models"
	s "github.com/nunovieira220/tournament-generator-go/structures"
)

/*
 * Simple cup tests
 */

func TestFailSimpleCupIfInvalidTeamName(t *testing.T) {
	teams := []string{"Porto", "Benfica", "Sporting", "TO_BE_DEFINED"}
	teamsArray := s.NewArray(teams)
	_, err := g.SimpleCupGenerate(teamsArray, m.GeneratorOptions{Type: "simple-cup"})

	if err == nil {
		t.Fatalf("should return an error")
	}

	if fmt.Sprint(err) != "invalid team names" {
		t.Fatalf("should return correct error message: invalid team names")
	}
}

func TestGenerateNoTeamSimpleCupCompetition(t *testing.T) {
	teamsArray := *new(s.Array[string])
	result, err := g.SimpleCupGenerate(teamsArray, m.GeneratorOptions{Type: "simple-cup"})

	if err != nil {
		t.Fatalf("should not return an error")
	}

	if result.Length() > 0 {
		t.Fatalf("should return empty result list")
	}
}

func TestGenerateFourTeamSimpleCupCompetition(t *testing.T) {
	teams := []string{"Porto", "Benfica", "Sporting", "Braga"}
	teamsArray := s.NewArray(teams)
	result, err := g.SimpleCupGenerate(teamsArray, m.GeneratorOptions{Type: "simple-cup"})

	if err != nil {
		t.Fatalf("should not return an error")
	}

	if result.Length() != 2 {
		t.Fatalf("should return 2 games: %d", result.Length())
	}

	counter := make(map[string]int)

	for _, game := range result.List() {
		counter[game.HomeTeam] = counter[game.HomeTeam] + 1
		counter[game.AwayTeam] = counter[game.AwayTeam] + 1
	}

	for _, v := range counter {
		if v != 1 {
			t.Fatalf("each team should appear only 1 time")
		}
	}
}

func TestGenerateSixteenTeamSimpleCupCompetition(t *testing.T) {
	teams := []string{
		"Porto",
		"Benfica",
		"Braga",
		"Sporting",
		"Rio Ave",
		"Famalicão",
		"Guimarães",
		"Moreirense",
		"Santa Clara",
		"Gil Vicente",
		"Marítimo",
		"Boavista",
		"Paços de Ferreira",
		"Tondela",
		"Belenenses",
		"Portimonense",
	}

	teamsArray := s.NewArray(teams)
	result, err := g.SimpleCupGenerate(teamsArray, m.GeneratorOptions{Type: "simple-cup"})

	if err != nil {
		t.Fatalf("should not return an error")
	}

	if result.Length() != 8 {
		t.Fatalf("should return 8 games: %d", result.Length())
	}
}

func TestGenerateTwelveTeamSimpleCupCompetition(t *testing.T) {
	teams := []string{
		"Porto",
		"Benfica",
		"Braga",
		"Sporting",
		"Rio Ave",
		"Famalicão",
		"Guimarães",
		"Moreirense",
		"Santa Clara",
		"Gil Vicente",
		"Marítimo",
		"Boavista",
	}

	teamsArray := s.NewArray(teams)
	result, err := g.SimpleCupGenerate(teamsArray, m.GeneratorOptions{Type: "simple-cup"})

	if err != nil {
		t.Fatalf("should not return an error")
	}

	lowerRoundGames := s.NewArray([]m.GeneratorGame{})
	toBeHomeNotDefinedGames := s.NewArray([]m.GeneratorGame{})
	for _, game := range result.List() {
		if game.Round == 1 {
			lowerRoundGames.Push(game)
		}

		if game.AwayTeam != "TO_BE_DEFINED" && game.HomeTeam == "TO_BE_DEFINED" {
			toBeHomeNotDefinedGames.Push(game)
		}
	}

	if toBeHomeNotDefinedGames.Length() > 0 {
		t.Fatalf("should not return games with home teams to be defined")
	}

	for _, game := range lowerRoundGames.List() {
		top := s.NewArray([]m.GeneratorGame{})

		for _, resultGame := range result.List() {
			if resultGame.CustomData["awayTeam"] == game.Id || resultGame.CustomData["homeTeam"] == game.Id {
				top.Push(resultGame)
			}
		}

		if top.Length() != 1 {
			t.Fatalf("should have only one game in the top round")
		}

		if top.Get(0).Round != 2 {
			t.Fatalf("should return game in round 2")
		}
	}
}

func TestGenerateThirteenTeamSimpleCupCompetition(t *testing.T) {
	teams := []string{
		"Porto",
		"Benfica",
		"Braga",
		"Sporting",
		"Rio Ave",
		"Famalicão",
		"Guimarães",
		"Moreirense",
		"Santa Clara",
		"Gil Vicente",
		"Marítimo",
		"Boavista",
		"Paços de Ferreira",
	}

	teamsArray := s.NewArray(teams)
	result, err := g.SimpleCupGenerate(teamsArray, m.GeneratorOptions{Type: "simple-cup"})

	if err != nil {
		t.Fatalf("should not return an error")
	}

	singleToBeDefinedGames := s.NewArray([]m.GeneratorGame{})
	bothToBeDefinedGames := s.NewArray([]m.GeneratorGame{})
	for _, game := range result.List() {
		if game.AwayTeam == "TO_BE_DEFINED" && game.HomeTeam == "TO_BE_DEFINED" {
			bothToBeDefinedGames.Push(game)
		}

		if game.AwayTeam != "TO_BE_DEFINED" && game.HomeTeam == "TO_BE_DEFINED" {
			singleToBeDefinedGames.Push(game)
		}
	}

	if bothToBeDefinedGames.Length() != 2 {
		t.Fatalf("should return 2 games with both teams to be defined")
	}

	if singleToBeDefinedGames.Length() != 1 {
		t.Fatalf("should return 1 game with just one team to be defined")
	}
}
