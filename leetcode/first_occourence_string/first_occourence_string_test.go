package first_occourence_string

import (
	"testing"
)

func Test_first_occourence_in_string(t *testing.T) {
	haystacks := []string{"leetcode", "sadbutsad"}
	needles := []string{"leeto", "sad"}
	results := []int{-1, 0}

	for i := range haystacks {
		str := haystacks[i]
		needle := needles[i]
		r := First_occourence_in_string(str, needle)

		if results[i] != r {
			t.Error("Test Failed")
		}
	}
}
