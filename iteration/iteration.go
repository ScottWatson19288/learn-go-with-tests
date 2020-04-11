package iteration

const repeatCount = 5

// Repeat takes a string and repeats it and then
// returns that output
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
