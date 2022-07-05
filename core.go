package core

/*
 * Dependencies
 */

import (
	"errors"

	"github.com/nunovieira220/tournament-generator-go/generators"
	m "github.com/nunovieira220/tournament-generator-go/models"
	s "github.com/nunovieira220/tournament-generator-go/structures"
)

/*
 * Generators
 */

var generatorMap map[string]interface{} = map[string]interface{}{
	"double-round": generators.DoubleRoundGenerate,
	"simple-cup":   generators.SimpleCupGenerate,
	"single-round": generators.SingleRoundGenerate,
}

/*
 * Main generate method
 */

func Generate(teams []string, options map[string]string) ([]m.GeneratorGame, error) {
	generatorOptions := m.GeneratorOptions{ToBeDefinedValue: options["ToBeDefinedValue"], Type: options["Type"]}
	generator := generatorMap[generatorOptions.Type]

	if generator == nil {
		return nil, errors.New("unsupported generator type")
	}

	teamsArray := s.NewArray(teams)
	method := generator.(func(s.Array[string], m.GeneratorOptions) (s.Array[m.GeneratorGame], error))
	result, err := method(teamsArray, generatorOptions)

	return result.List(), err
}
