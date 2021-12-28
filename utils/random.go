package utils

import (
	"math/rand"
	"strings"
	"time"
)

func RandomString(choice int) string {
	var output strings.Builder
	rand.Seed(time.Now().Unix())

	if choice == 0 {
		//Only lowercase
		charSet := "abcdedfghijklmnopqrst"
		length := 10
		for i := 0; i < length; i++ {
			random := rand.Intn(len(charSet))
			randomChar := charSet[random]
			output.WriteString(string(randomChar))
		}
		//output.Reset()
		return output.String()
	} else {
		//Lowercase and Uppercase Both
		charSet := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"
		length := 20
		for i := 0; i < length; i++ {
			random := rand.Intn(len(charSet))
			randomChar := charSet[random]
			output.WriteString(string(randomChar))
		}
		return output.String()
	}
}
