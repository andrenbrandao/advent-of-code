package seeds

import (
	"strconv"
	"strings"
)

type SeedSet struct {
	seedInput   string
	internalSet map[int]bool
}

func NewSeedSet(line string) *SeedSet {
	aSet := SeedSet{seedInput: line, internalSet: make(map[int]bool)}

	return &aSet
}

func (s *SeedSet) Has(src int) bool {
	s.calculateSet(src)

	val := s.internalSet[src]

	return val
}

func (ss *SeedSet) calculateSet(src int) {
	s := strings.Split(ss.seedInput, ":")
	fields := strings.Fields(s[1])

	for i := 0; i < len(fields); i += 2 {
		seed, _ := strconv.Atoi(fields[i])
		aRange, _ := strconv.Atoi(fields[i+1])

		if src >= seed && src <= seed+aRange-1 {
			ss.internalSet[src] = true
		}

	}
}
