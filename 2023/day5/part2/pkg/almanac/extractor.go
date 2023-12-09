package almanac

import "day5/pkg/almanac/mapper"

func extractMap(i int, lines []string) (int, *mapper.Map) {
	top := i
	bottom := i
	for len(lines[bottom]) > 0 {
		bottom++
	}

	m := mapper.NewMap(lines[top:bottom])
	return bottom, m
}

func extractOptimizedMap(i int, lines []string) (int, *mapper.OptimizedMap) {
	top := i
	bottom := i
	for len(lines[bottom]) > 0 {
		bottom++
	}

	m := mapper.NewOptimizedMap(lines[top:bottom])
	return bottom, m
}
