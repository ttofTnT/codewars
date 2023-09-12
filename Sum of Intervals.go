package codewars

import "sort"

func SumOfIntervals(intervals [][2]int) int {
	// Sort intervals by first data
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergedIntervals := make([][2]int, 0)
	mergedIntervals = append(mergedIntervals, intervals[0])
	// Merge overlapping intervals
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= mergedIntervals[len(mergedIntervals)-1][1] {
			mergedIntervals[len(mergedIntervals)-1][1] = max(intervals[i][1], mergedIntervals[len(mergedIntervals)-1][1])
		} else {
			mergedIntervals = append(mergedIntervals, intervals[i])
		}
	}

	// Calculate the sum of interval lengths
	sum := 0
	for _, v := range mergedIntervals {
		sum += v[1] - v[0]
	}

	return sum
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
