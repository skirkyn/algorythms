package algorythms

func strongPasswordChecker(s string) int {

	var prevRune rune
	times := 0
	lCaseCount := 0
	uCaseCount := 0
	var changes int
	if len(s) < 6 {
		changes = 6 - len(s)
	} else if len(s) > 20 {
		changes = len(s) - 20
	}
	for i, r := range s {
		if i == 0 {
			times++
			prevRune = r
		} else if prevRune == r {
			times++
			if times > 3 && i > 19 {
				changes++
			}
		} else {
			prevRune = r
			times = 0
		}
	}
	if uCaseCount == 0 && len(s) >= 6 && len(s)<=20 {
		changes++
	}
	if lCaseCount == 0 && len(s) >= 6 && len(s)<=20{
		changes++
	}

	return changes
}
