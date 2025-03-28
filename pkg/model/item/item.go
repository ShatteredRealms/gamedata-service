package item

import "github.com/ShatteredRealms/gamedata-service/pkg/model/stats/skills"

type Item struct {
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	Quality     uint32 `json:"quality"`
	Rarity      Rarity `json:"rarity"`
	Description string `json:"description"`
	EquipSlots  []Slot `json:"equip_slots"`

	MaxStack uint `json:"max_stack"`

	Requirements map[skills.SkillName]int `json:"requirements"`
	Enhancements map[skills.SkillName]int `json:"enhancements"`
}

type ItemInstance struct {
	Item
	Slot     Slot `json:"slot"`
	Quantity uint `json:"quantity"`
}
