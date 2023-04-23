package database

const base_goblin_hp int = 50
const base_goblin_dmg int = 10

type Goblin struct {
	ID               string  `json:"id"`
	HelthPoint       float32 `json:"health_point"`
	Level            int     `json:"level"`
	BaseAttackDamage float32 `json:"base_attack_damage"`
}

func NewEnemy(enemy_id string, level int) Goblin {
	goblin := Goblin{
		ID:               enemy_id,
		Level:            level,
		HelthPoint:       float32(base_goblin_hp) * (1.00 + float32(level/10)),
		BaseAttackDamage: float32(base_goblin_dmg) * (1.00 + float32(level/10)),
	}

	return goblin
}
