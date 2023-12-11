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
	maxTimes := []types.Time{}
	recordDistances := []types.Distance{}

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
		time := maxTimes[i]
		distance := recordDistances[i]
		race := NewRace(time, distance)
		races = append(races, race)
	}

	waysToWin := []int{}
	for i := 0; i < len(races); i++ {
		waysToWin = append(waysToWin, races[i].WaysToWin())
	}

	return waysToWin
}

func (rr *RaceRecords) MultipliedNumberWays() int {
	waysToWin := rr.WaysToWinAllRaces()

	if len(waysToWin) == 0 {
		return 0
	}

	mul := 1
	for _, ways := range waysToWin {
		mul *= ways
	}

	return mul
}

func (*RaceRecords) parseTimes(line string) []types.Time {
	maxTimes := []types.Time{}
	s := strings.Split(line, ":")
	fields := strings.Fields(s[1])
	v := strings.Join(fields, "")

	intVal, _ := strconv.Atoi(v)
	maxTimes = append(maxTimes, types.Time(intVal))

	return maxTimes
}

func (*RaceRecords) parseDistances(line string) []types.Distance {
	distances := []types.Distance{}
	s := strings.Split(line, ":")
	fields := strings.Fields(s[1])
	v := strings.Join(fields, "")

	intVal, _ := strconv.Atoi(v)
	distances = append(distances, types.Distance(intVal))

	return distances
}
