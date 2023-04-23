package service

import (
	"gin_api/database"
	"gin_api/model"
)

var goblin database.Goblin

func StartGame() model.StartGameResponse {
	hero = CreateNewHero()
	goblin = database.NewEnemy("E01", 1)

	response := model.StartGameResponse{
		Hero:  hero,
		Enemy: goblin,
	}

	return response
}

func EndTurn(skill_id string) model.StartGameResponse {
	hero_dmg, err := UseSkill(skill_id)
	if err != nil {
		return model.StartGameResponse{}
	}

	goblin.HelthPoint -= hero_dmg

	return model.StartGameResponse{
		Hero:  hero,
		Enemy: goblin,
	}
}
