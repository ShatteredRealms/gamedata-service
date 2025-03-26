package level

import (
	"fmt"
	"os"

	"github.com/ShatteredRealms/go-common-service/pkg/log"
)

const (
	MaxLevel     = 20
	GainIncrXp   = 22 * 4
	GainIncrMult = 2.5

	BaseNeededXp   = 900
	NeededIncrXp   = 100
	NeededIncrMult = 1.05

	BaseKillXp   = 70
	KillIncrXp   = 0
	KillIncrMult = 0.8

	EasyKillXpGain = 0.3
	MedKillXpGain  = 1.0
	HardKillXpGain = 1.2

	// XP gained from killing a mob
	//		Difficulty gain: easy to hard -> 0.3, 1.0, 1.2
	//		Level lower gain: (levelDiff + 1)^-2
	//		Max gain: 10% of needed XP
	//		Min gain: 1% of needed XP

	// Average Time to Kill (TTK) for med difficulty is 15s
)

type LevelXpData struct {
	XpGained   int
	XpNeeded   int
	BaseKillXp int
}

var (
	XpData = [MaxLevel - 1]LevelXpData{}
)

func init() {
	for level := 0; level < MaxLevel-1; level++ {
		XpData[level] = LevelXpData{
			XpGained:   int(GainIncrMult * GainIncrXp * float64(level)),
			XpNeeded:   BaseNeededXp,
			BaseKillXp: BaseKillXp,
		}
		if level > 0 {
			XpData[level].XpNeeded = XpData[level-1].XpNeeded + int(NeededIncrMult*NeededIncrXp*float64(level))
			XpData[level].BaseKillXp = XpData[level-1].BaseKillXp + int(KillIncrMult*KillIncrXp*float64(level))
		}
	}

	file, err := os.Create("xp_data.csv")
	if err != nil {
		log.Logger.Errorf("creating xp_data.csv: %v", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("Level,Xp Needed,Xp Gained,Kill Xp\n"); err != nil {
		log.Logger.Errorf("writing xp_data.csv: %v", err)
		return
	}
	for level, data := range XpData {
		if _, err := file.WriteString(fmt.Sprintf("%d,%d,%d,%d\n", level+1, data.XpNeeded, data.XpGained, data.BaseKillXp)); err != nil {
			log.Logger.Errorf("writing xp_data.csv: %v", err)
			return
		}
	}
}
