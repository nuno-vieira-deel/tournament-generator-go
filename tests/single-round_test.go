package test

/*
 * Dependencies
 */

import (
	"testing"

	g "github.com/nunovieira220/tournament-generator-go/generators"
	m "github.com/nunovieira220/tournament-generator-go/models"
	s "github.com/nunovieira220/tournament-generator-go/structures"
)

/*
 * Single round tests
 */

func TestGenerateNoTeamSingleRoundCompetition(t *testing.T) {
	teamsArray := *new(s.Array[string])
	result, err := g.SingleRoundGenerate(teamsArray, m.GeneratorOptions{Type: "single-round"})

	if err != nil {
		t.Errorf("should not return an error")
	}

	if result.Length() > 0 {
		t.Errorf("should return empty result list")
	}
}

func TestGenerateFourTeamSingleRoundCompetition(t *testing.T) {
	teams := []string{"Porto", "Benfica", "Sporting", "Braga"}
	teamsArray := s.NewArray(teams)
	result, err := g.SingleRoundGenerate(teamsArray, m.GeneratorOptions{Type: "single-round"})

	if err != nil {
		t.Errorf("should not return an error")
	}

	if result.Length() != 6 {
		t.Errorf("should return 6 games: %d", result.Length())
	}

	counter := make(map[string]int)

	for _, game := range result.List() {
		counter[game.HomeTeam] = counter[game.HomeTeam] + 1
		counter[game.AwayTeam] = counter[game.AwayTeam] + 1
	}

	for _, v := range counter {
		if v != 3 {
			t.Errorf("each team should appear only 3 times")
		}
	}

}

func TestGenerateEigthteenTeamSingleRoundCompetition(t *testing.T) {
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
		"Setubal",
		"Portimonense",
		"Aves",
	}

	teamsArray := s.NewArray(teams)
	result, err := g.SingleRoundGenerate(teamsArray, m.GeneratorOptions{Type: "single-round"})

	if err != nil {
		t.Errorf("should not return an error")
	}

	if result.Length() != 153 {
		t.Errorf("should return 153 games %d", result.Length())
	}
}

func TestGenerateNineteenTeamSingleRoundCompetition(t *testing.T) {
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
		"Setubal",
		"Portimonense",
		"Aves",
		"Tottenham",
	}

	teamsArray := s.NewArray(teams)
	result, err := g.SingleRoundGenerate(teamsArray, m.GeneratorOptions{Type: "single-round"})

	if err != nil {
		t.Errorf("should not return an error")
	}

	if result.Length() != 171 {
		t.Errorf("should return 171 games: %d", result.Length())
	}
}
