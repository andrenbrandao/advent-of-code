package almanac

func extractMap(i int, lines []string) (int, *Map) {
	top := i
	bottom := i
	for len(lines[bottom]) > 0 {
		bottom++
	}

	m := NewMap(lines[top:bottom])
	return bottom, m
}

func extractOptimizedMap(i int, lines []string) (int, *OptimizedMap) {
	top := i
	bottom := i
	for len(lines[bottom]) > 0 {
		bottom++
	}

	m := NewOptimizedMap(lines[top:bottom])
	return bottom, m
}
