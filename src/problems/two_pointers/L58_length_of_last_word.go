package two_pointers

func lengthOfLastWord(s string) int {
	last := len(s) - 1
	for last >= 0 && s[last] == ' ' {
		last = last - 1
	}
	if last == -1 {
		return 0
	}
	first := last

	for first >= 0 && s[first] != ' ' {
		first = first - 1
	}
	return last - first
}
