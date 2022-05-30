package generators

/*
 * Dependencies
 */

import (
	m "github.com/nunovieira220/tournament-generator-go/models"
	s "github.com/nunovieira220/tournament-generator-go/structures"
	u "github.com/nunovieira220/tournament-generator-go/utils"
)

/*
 * Export generator
 */

func SingleRoundGenerate(teams s.Array[string], _ m.GeneratorOptions) (s.Array[m.GeneratorGame], error) {
	return u.ExecuteRoundRobinRotation(teams, false)
}
