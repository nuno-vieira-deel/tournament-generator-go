package models

/*
 * Types
 */

type GeneratorError struct {
	CustomData interface{} `json:"customData,omitempty"`
	Message    string      `json:"message"`
	Status     int         `json:"status"`
}

type GeneratorGame struct {
	AwayTeam   string                 `json:"awayTeam"`
	CustomData map[string]interface{} `json:"customData,omitempty"`
	HomeTeam   string                 `json:"homeTeam"`
	Id         string                 `json:"id,omitempty"`
	Round      int                    `json:"round"`
	Score      string                 `json:"score,omitempty"`
}

type GeneratorOptions struct {
	ToBeDefinedValue string `json:"toBeDefinedValue"`
	Type             string `json:"type"`
}
