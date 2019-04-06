package golangexamples

import "github.com/ehteshamz/greetings"

func ConcatSlice(sliceToConcat []byte) string {
	lengthOfArray := len(sliceToConcat)
	var concatstring string
	
	for i := 0; i < lengthOfArray; i++ {
		concatstring += string(sliceToConcat[i])
		if i != lengthOfArray-1 {
			concatstring += "-"
		}
	}
	return concatstring
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	lengthOfArray := len(sliceToEncrypt)
	
	for i := 0; i < lengthOfArray; i++ {
		sliceToEncrypt[i] += byte(ceaserCount-48)
	}
}

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}