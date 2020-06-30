package authGen

import (
	"math/rand"
)


func GetAPIKey() string{

	APIKEY := "xxxxxx-xxxxxx-xxxx-xxxxxx"

	returnKey := editKey(APIKEY)

	return returnKey

}

func editKey(apiKey string) string{
	runes := []rune(apiKey)
	for i := 0; i < len(apiKey); i++ {
		if runes[i] == 'x' {
			randomInt := rand.Intn (127);

			if randomInt <= 63 {
				newChar := getInt(randomInt)
				runes[i] = newChar
			} else{
				newChar := getChar(randomInt)
				runes[i] = newChar
			}
		}
	}
	return string(runes)
}

func getInt (index int) rune{
	 num := index % 9
	 return rune(num + 48)
}

func getChar (index int) rune {
		if index <= 113 {
			num := index % 25
			return rune(num + 97)
		} else {
			leftOver := 127 - index
			return rune(leftOver + 97)
		}
}