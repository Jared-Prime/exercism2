package scrabble

import "unicode"

var englishDictionary = map[int]string{
	1:  "AEIOULNRST",
	2:  "DG",
	3:  "BCMP",
	4:  "FHVWY",
	5:  "K",
	8:  "JX",
	10: "QZ",
}

var englishScoreMap = scoreMap(englishDictionary)

// Score assigns a value to word
func Score(letters string) int {
	var score int

	for _, letter := range letters {
		letter = unicode.ToUpper(letter)
		score = score + englishScoreMap[letter]
	}

	return score
}

func scoreMap(dictionary map[int]string) map[rune]int {
	scores := make(map[rune]int)

	for value, letters := range dictionary {
		for _, letter := range letters {
			scores[letter] = value
		}
	}

	return scores
}
