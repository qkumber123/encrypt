package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	// enMsg, hiddenKey :=
	encryptMsg("Hello World")
	// decryptMsg(enMsg, hiddenKey)
}

func makeLetterArray() []rune {
	var letters []rune

	for i := 32; i <= 126; i++ {
		letters = append(letters, rune(i))
	}

	return letters
}

func arrayShuffle() map[rune]rune {
	letters := makeLetterArray()
	keyMap := make(map[rune]rune)
	var randLetters []rune

	// copy(letters, randLetters)
	randLetters = append(randLetters, letters...)

	rand.Shuffle(len(randLetters), func(i, j int) {
		randLetters[i], randLetters[j] = randLetters[j], randLetters[i]
	})

	sb := strings.Builder{}

	// divining runes
	for i, randRune := range randLetters {
		keyMap[letters[i]] = randRune
		sb.WriteRune(randRune)
	}

	fmt.Println("Your key is:")
	fmt.Println(sb.String())

	return keyMap

}

func encryptMsg(msg string) (string, map[rune]rune) {
	encMsg := strings.Builder{}
	keyMap := arrayShuffle()

	// encrypt msg
	for i := 0; i < len(msg); i++ {
		encMsg.WriteRune(keyMap[rune(msg[i])])
	}

	encryptedMsg := encMsg.String()

	fmt.Println(msg)
	fmt.Println(encryptedMsg)

	return encryptedMsg, keyMap

}

// func decryptMsg(encodedMsg string, key map[rune]rune) {
// 	letters := makeLetterArray()
// 	decoderMap := make(map[rune]rune)

// 	for i := 0-; i < len(letters); i++ {
// 		decodermap = key[i]:letters[i]
// 	}

// }
