package main

func main() {

}

func RotateRight(text string) string {
	size := len(text)
	rotatedText := ""
	for fromIndex := range text {
		toIndex := (size - 1 + fromIndex) % size
		char := string(text[toIndex])
		rotatedText = rotatedText + char
	}
	return rotatedText
}
