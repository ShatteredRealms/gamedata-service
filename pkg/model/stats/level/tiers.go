package level

import (
	"github.com/ShatteredRealms/gamedata-service/pkg/model/profession"
)

type CharacterTierLevel int
type CharacterTier struct {
	Level     CharacterTierLevel
	Name      map[profession.Profession]string
	Threshold int
}

const (
	TierLevel1 CharacterTierLevel = iota + 1
	TierLevel2
	MaxTierLevel
)

var (
	Tiers = make(map[CharacterTierLevel]CharacterTier)
	Tier1 = registerTier(CharacterTier{
		Level: TierLevel1,
		Name: map[profession.Profession]string{
			profession.Necromancer: "Novice",
		},
		Threshold: 0,
	})
	Tier2 = registerTier(CharacterTier{
		Level: TierLevel2,
		Name: map[profession.Profession]string{
			profession.Necromancer: "Apprentice",
		},
		Threshold: 10,
	})
)

func registerTier(t CharacterTier) CharacterTier {
	Tiers[t.Level] = t
	return t
}

func TierLevelFromLevel(level int) CharacterTier {
	for tierLevel := TierLevel1; tierLevel < MaxTierLevel; tierLevel++ {
		if Tiers[tierLevel].Threshold > level {
			return Tiers[tierLevel]
		}
	}
	return CharacterTier{
		Level:     CharacterTierLevel(-1),
		Name:      make(map[profession.Profession]string),
		Threshold: -1,
	}
}
