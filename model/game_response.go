package model

import "gin_api/database"

type StartGameResponse struct {
	Hero  database.Hero   `json:"hero"`
	Enemy database.Goblin `json:"enemy"`
}
