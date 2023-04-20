package database

type Hero struct {
	ID               string  `json:"id"`
	HelthPoint       float32 `json:"health_point"`
	FocusPoint       int     `json:"focus_point"`
	BaseAttackDamage float32 `json:"base_attack_damage"`
	Skills           []Skill `json:"skills"`
}

type Skill struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Effect       float32 `json:"effect"`
	FocusPoint   int     `json:"focus_point"`
	CooldownTime int     `json:"cooldown_time"`
}

var Hero_skills = []Skill{
	{ID: "S01", Name: "Light Attack", Effect: 1.00, FocusPoint: 0, CooldownTime: 0},
	{ID: "S02", Name: "Heavy Attack", Effect: 1.50, FocusPoint: 0, CooldownTime: 2},
	{ID: "S03", Name: "Heavier Attack", Effect: 2.00, FocusPoint: 10, CooldownTime: 2},
	{ID: "S04", Name: "Very Heavy Attack", Effect: 2.50, FocusPoint: 20, CooldownTime: 3},
	{ID: "S05", Name: "Fucking Heavy Attack", Effect: 3.00, FocusPoint: 30, CooldownTime: 5},
	{ID: "S06", Name: "The Most Fucking Heavy Attack", Effect: 5.00, FocusPoint: 50, CooldownTime: 8},
}
