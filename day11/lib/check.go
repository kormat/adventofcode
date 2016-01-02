package day11

func Validate(passwd string) bool {
	run := false
	pairs := false
	for i := 0; i < len(passwd); i++ {
		if isBannedChar(passwd[i]) {
			return false
		}
		if !run {
			run = findRun(passwd[i:])
		}
		if !pairs && (i <= len(passwd)-4) && (passwd[i] == passwd[i+1]) {
			pairs = findPair(passwd[i+2:])
		}
	}
	if run && pairs {
		return true
	} else {
		return false
	}
}

func isBannedChar(c byte) bool {
	switch c {
	case 'i', 'o', 'l':
		return true
	default:
		return false
	}
}

func findRun(input string) bool {
	if len(input) < 3 {
		return false
	}
	if (input[0]+1 == input[1]) && (input[1]+1 == input[2]) {
		return true
	}
	return false
}

func findPair(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return true
		}
	}
	return false
}
