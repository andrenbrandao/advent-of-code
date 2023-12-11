package race

import (
	"day6/pkg/race/boat"
	"day6/pkg/race/types"
)

type Race struct {
	maxTime        types.Time
	recordDistance types.Distance
}

func NewRace(maxTime types.Time, recordDistance types.Distance) *Race {
	return &Race{maxTime: maxTime, recordDistance: recordDistance}
}

func (r *Race) WaysToWin() int {
	result := 0
	for i := 1; i < int(r.maxTime)-1; i++ {
		charge := types.Charge(i)
		remainingTime := r.maxTime - types.Time(i)

		boat := boat.NewBoat(charge)
		if boat.Distance(remainingTime) > r.recordDistance {
			result++
		}
	}

	return result
}
