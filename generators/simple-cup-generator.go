package generators

/*
 * Dependencies
 */

import (
	"errors"
	"math"

	m "github.com/nunovieira220/tournament-generator-go/models"
	s "github.com/nunovieira220/tournament-generator-go/structures"
	u "github.com/nunovieira220/tournament-generator-go/utils"
)

/*
 * Constants
 */

const TO_BE_DEFINED_CONSTANT = "TO_BE_DEFINED"

/*
 * Export generator
 */

func SimpleCupGenerate(teams s.Array[string], options m.GeneratorOptions) (s.Array[m.GeneratorGame], error) {
	toBeDefined := options.ToBeDefinedValue

	if toBeDefined == "" {
		toBeDefined = TO_BE_DEFINED_CONSTANT
	}

	data := s.NewArray([]m.GeneratorGame{})

	for _, v := range teams.List() {
		if v == toBeDefined {
			return data, errors.New("invalid team names")
		}
	}

	teamList := u.Shuffle(teams.List())
	length := teams.Length()
	logLength := math.Log2(float64(length))
	topRoundTeamsNumber := int(math.Pow(2, math.Floor(logLength)))
	var multiRounds bool = math.Mod(logLength, 1) != 0
	var lowRoundIndex int = topRoundTeamsNumber

	for i := 0; i < topRoundTeamsNumber; i += 2 {
		var customData = map[string]any{}
		var homeTeam string = teamList[i]
		var awayTeam string = teamList[i+1]

		if multiRounds {
			// 1st round game for home spot
			if lowRoundIndex < length {
				id, _ := u.GetUniqueValue()

				data.Push(m.GeneratorGame{
					AwayTeam: teamList[lowRoundIndex],
					HomeTeam: homeTeam,
					Id:       id,
					Round:    1,
				})

				customData["homeTeam"] = id
				homeTeam = toBeDefined
				lowRoundIndex++
			}

			// 2nd round game for away spot
			if lowRoundIndex < length {
				id, _ := u.GetUniqueValue()

				data.Push(m.GeneratorGame{
					AwayTeam: awayTeam,
					HomeTeam: teamList[lowRoundIndex],
					Id:       id,
					Round:    1,
				})

				customData["awayTeam"] = id
				awayTeam = toBeDefined
				lowRoundIndex++
			}
		}

		// 2nd round game
		id, _ := u.GetUniqueValue()
		round := 1
		if multiRounds {
			round = 2
		}

		data.Push(m.GeneratorGame{
			AwayTeam:   awayTeam,
			CustomData: customData,
			Id:         id,
			HomeTeam:   homeTeam,
			Round:      round,
		})
	}

	return data, nil
}
