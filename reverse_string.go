package reverse_string

func ReverseString(input string) (output string) {
	for _, character := range input {
		output = string(character) + output
	}
	return output
}
