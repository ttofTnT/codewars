package backtracking

import (
	"sort"
)

// https://www.codewars.com/kata/5254ca2719453dcc0b00027d/train/go
func Permutations(input string) []string {
	// Convert the input string to a slice of characters for easy manipulation.
	inputChars := []rune(input)

	// Create a map to store unique permutations.
	uniquePermutations := make(map[string]bool)

	// Define a recursive function to generate permutations.
	var generatePermutations func(start int)
	generatePermutations = func(start int) {
		if start == len(inputChars) {
			// Convert the slice of characters back to a string.
			permutation := string(inputChars)
			uniquePermutations[permutation] = true
			return
		}

		for i := start; i < len(inputChars); i++ {
			// Swap characters at positions start and i.
			inputChars[start], inputChars[i] = inputChars[i], inputChars[start]
			// Recursively generate permutations for the remaining characters.
			generatePermutations(start + 1)
			// Restore the original order of characters (backtracking).
			inputChars[start], inputChars[i] = inputChars[i], inputChars[start]
		}
	}

	// Start generating permutations from index 0.
	generatePermutations(0)

	// Convert the unique permutations map into a slice of strings.
	permutations := make([]string, 0, len(uniquePermutations))
	for perm := range uniquePermutations {
		permutations = append(permutations, perm)
	}

	// Sort the permutations lexicographically.
	sort.Strings(permutations)

	return permutations
}

func Permutations1(s string) (a []string) {
	if len(s) < 2 {
		return []string{s}
	}
	m := map[string]bool{}
	for _, sub := range Permutations(s[1:]) {
		for i := 0; i <= len(sub); i++ {
			st := sub[0:i] + s[0:1] + sub[i:]
			if m[st] {
				continue
			}
			m[st] = true
			a = append(a, st)
		}
	}
	return
}
