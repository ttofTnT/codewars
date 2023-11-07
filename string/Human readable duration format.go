package string

import (
	"fmt"
	"strings"
)

// https://www.codewars.com/kata/52742f58faf5485cae000b9a/train/go
func FormatDuration(seconds int64) string {
	if seconds == 0 {
		return "now"
	}

	components := []string{}
	years := seconds / 31536000
	if years > 0 {
		components = append(components, pluralize(years, "year"))
		seconds -= years * 31536000
	}

	days := seconds / 86400
	if days > 0 {
		components = append(components, pluralize(days, "day"))
		seconds -= days * 86400
	}

	hours := seconds / 3600
	if hours > 0 {
		components = append(components, pluralize(hours, "hour"))
		seconds -= hours * 3600
	}

	minutes := seconds / 60
	if minutes > 0 {
		components = append(components, pluralize(minutes, "minute"))
		seconds -= minutes * 60
	}

	if seconds > 0 {
		components = append(components, pluralize(seconds, "second"))
	}

	if len(components) == 1 {
		return components[0]
	}

	last := components[len(components)-1]
	components = components[:len(components)-1]
	return strings.Join(components, ", ") + " and " + last
}

func pluralize(n int64, unit string) string {
	if n > 1 {
		return fmt.Sprintf("%d %ss", n, unit)
	}
	return fmt.Sprintf("%d %s", n, unit)
}
