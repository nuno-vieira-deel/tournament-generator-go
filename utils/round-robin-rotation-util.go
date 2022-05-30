package utils

import (
	m "github.com/nunovieira220/tournament-generator-go/models"
	s "github.com/nunovieira220/tournament-generator-go/structures"
)

/*
 * Round robin rotation method
 */

func ExecuteRoundRobinRotation(teams s.Array[string], isDouble bool) (s.Array[m.GeneratorGame], error) {
	oddExtraId, err := GetUniqueValue()
	firstRound := s.NewArray([]m.GeneratorGame{})
	secondRound := s.NewArray([]m.GeneratorGame{})
	isOdd := false

	if err != nil {
		return firstRound, err
	}

	if teams.Length()%2 == 1 {
		isOdd = true
		teams.Push(oddExtraId)
	}

	numberOfTeams := teams.Length()
	teamList := Shuffle(teams.List())
	homeTeams := s.NewArray(teamList[0 : numberOfTeams/2])
	awayTeams := s.NewArray(teamList[numberOfTeams/2 : numberOfTeams])

	for i := 0; i < numberOfTeams-1; i++ {
		for j := 0; j < homeTeams.Length(); j++ {
			teams := Shuffle([]string{homeTeams.Get(j), awayTeams.Get(j)})
			firstId, _ := GetUniqueValue()
			round := i + 1

			// 1st round game
			firstRound.Push(m.GeneratorGame{
				AwayTeam: teams[1],
				HomeTeam: teams[0],
				Id:       firstId,
				Round:    round,
			})

			if isDouble {
				// 2nd round game
				secondId, _ := GetUniqueValue()

				secondRound.Push(m.GeneratorGame{
					AwayTeam: teams[0],
					HomeTeam: teams[1],
					Id:       secondId,
					Round:    round + numberOfTeams - 1,
				})
			}
		}

		// Rotation
		homeFixedTeam := homeTeams.Shift()
		homeTeams.Unshift(awayTeams.Shift())
		homeTeams.Unshift(homeFixedTeam)
		awayTeams.Push(homeTeams.Pop())
	}

	fixtures := append(firstRound.List(), secondRound.List()...)

	if !isOdd {
		return s.NewArray(fixtures), nil
	}

	result := s.NewArray([]m.GeneratorGame{})

	for _, game := range fixtures {
		if game.HomeTeam != oddExtraId && game.AwayTeam != oddExtraId {
			result.Push(game)
		}
	}

	return result, nil
}
