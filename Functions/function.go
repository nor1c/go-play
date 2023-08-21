package main

// shorten parameters type
func shortenParametersType(x, y int) int {
	return x + y
}

// return multiple values (string, string)
func returnMultipleValues(message string, name string) (string, string) {
	return message, name
}

func main() {
	shortenParametersType(1, 2)

	returnMultipleValues("hello", "fauzi")
}
