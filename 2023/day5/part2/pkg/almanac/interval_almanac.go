package almanac

// type IntervalAlmanac struct {
// 	seedsInput string
// 	maps       []
// }

// func NewIntervalAlmanac(input string) *IntervalAlmanac {
// 	lines := strings.Split(input, "\n")
// 	var seedsInput string
// 	maps := []*OptimizedMap{}

// 	for i := 0; i < len(lines); i++ {
// 		line := lines[i]

// 		if len(line) == 0 {
// 			continue
// 		}

// 		if strings.HasPrefix(line, "seeds:") {
// 			seedsInput = line
// 		}

// 		if strings.Contains(line, "map:") {
// 			var aMap *OptimizedMap
// 			i, aMap = extractOptimizedMap(i+1, lines)
// 			maps = append(maps, aMap)
// 		}
// 	}

// 	return &IntervalAlmanac{
// 		seedsInput: seedsInput,
// 		maps:       maps,
// 	}
// }
