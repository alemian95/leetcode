package first_occourence_string

func First_occourence_in_string(haystack string, needle string) int {
	for i := range haystack {
		if i+len(needle) <= len(haystack) && haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
