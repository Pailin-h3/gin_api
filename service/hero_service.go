package service

import (
	"errors"
	"gin_api/database"
)

var hero database.Hero

func CreateNewHero() database.Hero {
	hero = database.Hero{
		ID:               "000",
		Level:            1,
		SkillPoint:       0,
		HelthPoint:       200.00,
		FocusPoint:       100,
		BaseAttackDamage: 5.00,
		Skills:           database.Hero_skills[0:2],
	}

	return hero
}

func GetHero() database.Hero {
	return hero
}

func LearnSkill(skill_id string) error {
	for _, skill := range hero.Skills {
		if skill.ID == skill_id {
			return errors.New("ALREADY LEARN")
		}
	}
	for i, skill := range database.Hero_skills {
		if skill.ID == skill_id {
			if hero.SkillPoint >= skill.SkillPoint {
				hero.Skills = append(hero.Skills, database.Hero_skills[i])
				hero.SkillPoint -= skill.SkillPoint
				return nil
			}
			return errors.New("INSUFFICIENT SKILL POINT")
		}
	}

	return errors.New("NO SKILL ID NOT FOUND")
}

func UseSkill(skill_id string) (float32, error) {
	for _, skill := range hero.Skills {
		if skill.ID == skill_id {
			return skill.Effect, nil
		}
	}

	return 0.00, errors.New("SKILL NOT FOUND")
}

func GotAttacked(atk_point float32) error {
	hero.HelthPoint -= atk_point
	return nil
}
