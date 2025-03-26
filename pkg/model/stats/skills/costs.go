package skills

import (
	"github.com/ShatteredRealms/gamedata-service/pkg/model/profession"
	"github.com/ShatteredRealms/gamedata-service/pkg/model/realm"
)

var (
	AbilityCosts = map[realm.Realm]map[AbilityName]float64{
		realm.Human: {
			Strength:   2.0,
			Stamina:    2.0,
			Agility:    2.0,
			Perception: 2.0,
			Intellect:  2.0,
			Willpower:  2.0,
		},
		realm.Cyborg: {
			Strength:   1.5,
			Stamina:    2.0,
			Agility:    4.0,
			Perception: 1.5,
			Intellect:  2.0,
			Willpower:  1.0,
		},
	}
	SkillCosts = map[profession.Profession]map[SkillName]float64{
		profession.Necromancer: {
			Endurance:           2.0,
			EnergyReserve:       2.0,
			CloseCombatEvasion:  2.0,
			RangedCombatEvasion: 2.0,
			Deflection:          2.0,
			MartialArts:         2.0,
			OneHanded:           2.0,
			TwoHanded:           2.0,
			Piercing:            2.0,
			WeaponSpeed:         2.0,
			MeleeSpeed:          2.0,
			SneakAttack:         2.0,
			QuickStrike:         2.0,
			PowerStrike:         2.0,
			RapidBlows:          2.0,
			BowSkill:            2.0,
			PistolSkill:         2.0,
			RifleSkill:          2.0,
			ShotgunSkill:        2.0,
			SniperSkill:         2.0,
			ThrowingSkill:       2.0,
			RangedSpeed:         2.0,
			SteadyShot:          2.0,
			QuickShot:           2.0,
			PowerShot:           2.0,
			RapidFire:           2.0,
			SilentShot:          2.0,
			Creation:            2.0,
			Metamorphosis:       2.0,
			Remedy:              2.0,
			Rejuvenation:        2.0,
			SpaceTime:           2.0,
			Destruction:         2.0,
			EnergySpeed:         2.0,
			Alchemy:             2.0,
			Programming:         2.0,
			Construction:        2.0,
			Engineering:         2.0,
			FirstAid:            2.0,
			Surgery:             2.0,
			Speed:               2.0,
			Flying:              2.0,
			Stealth:             2.0,
		},
	}
)
