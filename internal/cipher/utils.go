package cipher

func ShiftRune(r rune, shift int) rune {
	switch {
	case 'а' <= r && r <= 'я':
		return 'а' + rune((int(r-'а')+shift+32)%32)
	case 'А' <= r && r <= 'Я':
		return 'А' + rune((int(r-'А')+shift+32)%32)
	}

	switch {
	case 'a' <= r && r <= 'z':
		return 'a' + rune((int(r-'a')+shift+26)%26)
	case 'A' <= r && r <= 'Z':
		return 'A' + rune((int(r-'A')+shift+26)%26)
	}

	return r
}

func GetShift(r rune) int {
	switch {
	case 'A' <= r && r <= 'Z':
		return int(r - 'A')
	case 'a' <= r && r <= 'z':
		return int(r - 'a')
	case 'А' <= r && r <= 'Я':
		return int(r - 'А')
	case 'а' <= r && r <= 'я':
		return int(r - 'а')
	default:
		return 0
	}
}
