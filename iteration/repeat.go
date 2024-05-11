package iteration

func Repeat(character string) string {
	var repeatd string
	for i := 0; i < 5; i++ {
		repeatd = repeatd + character
	}
	return repeatd
}
