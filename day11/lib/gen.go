package day11

type Passwd [8]byte

func Generate(input string) (int, string) {
	var passwd Passwd
	copy(passwd[:], input)
	for i := 0; ; i++ {
		inc(&passwd)
		incPastBanned(&passwd)
		s := string(passwd[:])
		if Validate(s) {
			return i, s
		}
	}
}

func inc(passwd *Passwd) {
	for i := len(passwd) - 1; i >= 0; i-- {
		if passwd[i] == 'z' {
			passwd[i] = 'a'
		} else {
			passwd[i]++
			break
		}
	}
}

func incPastBanned(passwd *Passwd) {
	var banned = false
	for i := 0; i < len(*passwd); i++ {
		if banned {
			passwd[i] = 'a'
		} else if isBannedChar(passwd[i]) {
			banned = true
			// Safe because all banned chars are spread out, and don't include 'z'
			passwd[i]++
		}
	}
}
