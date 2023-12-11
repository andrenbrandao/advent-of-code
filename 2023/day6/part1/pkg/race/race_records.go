package race

import (
	"day6/pkg/race/types"
	"strconv"
	"strings"
)

type RaceRecords struct {
	input string
}

func NewRaceRecords(input string) *RaceRecords {
	return &RaceRecords{input}
}

func (rr *RaceRecords) WaysToWinAllRaces() []int {
	lines := strings.Split(rr.input, "\n")
	maxTimes := []int{}
	recordDistances := []int{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "Time:") {
			maxTimes = rr.parseTimes(line)
		}

		if strings.HasPrefix(line, "Distance:") {
			recordDistances = rr.parseDistances(line)
		}
	}

	races := []*Race{}

	for i := 0; i < len(maxTimes); i++ {
		time := types.Time(maxTimes[i])
		distance := types.Distance(recordDistances[i])
		race := NewRace(time, distance)
		races = append(races, race)
	}

	waysToWin := []int{}
	for i := 0; i < len(races); i++ {
		waysToWin = append(waysToWin, races[i].WaysToWin())
	}

	return waysToWin
}

func (*RaceRecords) parseTimes(line string) []int {
	maxTimes := []int{}
	s := strings.Split(line, ":")
	fields := strings.Fields(s[1])

	for _, f := range fields {
		time, _ := strconv.Atoi(f)
		maxTimes = append(maxTimes, time)
	}

	return maxTimes
}

func (*RaceRecords) parseDistances(line string) []int {
	distances := []int{}
	s := strings.Split(line, ":")
	fields := strings.Fields(s[1])

	for _, f := range fields {
		time, _ := strconv.Atoi(f)
		distances = append(distances, time)
	}

	return distances
}
