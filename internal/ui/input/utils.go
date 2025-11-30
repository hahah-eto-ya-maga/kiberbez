package input

func IsAlphaCyrillic(s string) bool {
	for _, r := range s {
		if !((r >= 'а' && r <= 'я') ||
			(r >= 'А' && r <= 'Я')) {
			return false
		}
	}
	return true
}
