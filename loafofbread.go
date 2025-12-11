package piscine

func LoafOfBread(str string) string {
	startString := ""
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "invalid Output\n"
	}
	ct := 0
	for _, ch := range str {
		if ch != ' ' && ct != 5 {
			startString += string(ch)
			ct++
		} else if ct == 5 {
			startString += " "
			ct = 0
		}
	}
	if len(startString) > 0 && startString[len(startString)-1] == ' ' {
		startString = startString[:len(startString)-1]
	}
	return startString + "\n"
}
