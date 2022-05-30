package test

/*
 * Dependencies
 */

import (
	"sort"
	"testing"

	g "github.com/nunovieira220/tournament-generator-go/generators"
	m "github.com/nunovieira220/tournament-generator-go/models"
	s "github.com/nunovieira220/tournament-generator-go/structures"
)

/*
 * Double round tests
 */

func TestGenerateNoTeamDoubleRoundCompetition(t *testing.T) {
	teamsArray := *new(s.Array[string])
	result, err := g.DoubleRoundGenerate(teamsArray, m.GeneratorOptions{Type: "double-round"})

	if err != nil {
		t.Errorf("should not return an error")
	}

	if result.Length() > 0 {
		t.Errorf("should return empty result list")
	}
}

func TestGenerateFourTeamDoubleRoundCompetition(t *testing.T) {
	teams := []string{"Porto", "Benfica", "Sporting", "Braga"}
	teamsArray := s.NewArray(teams)
	result, err := g.DoubleRoundGenerate(teamsArray, m.GeneratorOptions{Type: "double-round"})

	if err != nil {
		t.Errorf("should not return an error")
	}

	if result.Length() != 12 {
		t.Errorf("should return 12 games: %d", result.Length())
	}

	sortList := result.List()

	sort.Slice(sortList, func(i, j int) bool {
		if sortList[i].HomeTeam > sortList[j].HomeTeam {
			return false
		}

		if sortList[i].HomeTeam == sortList[j].HomeTeam && sortList[i].AwayTeam > sortList[j].AwayTeam {
			return false
		}

		return true
	})

	expectedGames := []m.GeneratorGame{
		{AwayTeam: "Braga", HomeTeam: "Benfica"},
		{AwayTeam: "Porto", HomeTeam: "Benfica"},
		{AwayTeam: "Sporting", HomeTeam: "Benfica"},
		{AwayTeam: "Benfica", HomeTeam: "Braga"},
		{AwayTeam: "Porto", HomeTeam: "Braga"},
		{AwayTeam: "Sporting", HomeTeam: "Braga"},
		{AwayTeam: "Benfica", HomeTeam: "Porto"},
		{AwayTeam: "Braga", HomeTeam: "Porto"},
		{AwayTeam: "Sporting", HomeTeam: "Porto"},
		{AwayTeam: "Benfica", HomeTeam: "Sporting"},
		{AwayTeam: "Braga", HomeTeam: "Sporting"},
		{AwayTeam: "Porto", HomeTeam: "Sporting"},
	}

	for i, game := range sortList {
		if game.AwayTeam != expectedGames[i].AwayTeam || game.HomeTeam != expectedGames[i].HomeTeam {
			t.Errorf("should have same teams on game: %s->%s %s->%s", game.HomeTeam, expectedGames[i].HomeTeam, game.AwayTeam, expectedGames[i].AwayTeam)
		}
	}
}

func TestGenerateEigthteenTeamDoubleRoundCompetition(t *testing.T) {
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
	result, err := g.DoubleRoundGenerate(teamsArray, m.GeneratorOptions{Type: "double-round"})

	if err != nil {
		t.Errorf("should not return an error")
	}

	if result.Length() != 306 {
		t.Errorf("should return 306 games: %d", result.Length())
	}
}

func TestGenerateNineteenTeamDoubleRoundCompetition(t *testing.T) {
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
	result, err := g.DoubleRoundGenerate(teamsArray, m.GeneratorOptions{Type: "double-round"})

	if err != nil {
		t.Errorf("should not return an error")
	}

	if result.Length() != 342 {
		t.Errorf("should return 342 games: %d", result.Length())
	}
}
