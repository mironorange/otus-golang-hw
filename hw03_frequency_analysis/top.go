package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type dimension struct {
	word  string
	count int
}

var reg = regexp.MustCompile(`\s+`)

func Top10(text string, strictCleaning bool) []string {
	if len(text) == 0 {
		return make([]string, 0)
	}

	cleaningWords := ClearAndSplit(text, strictCleaning)
	wordsAndRepetition := countWordsAndRepetition(cleaningWords)
	words := get10TopCountedWords(wordsAndRepetition)

	return words
}

func ClearAndSplit(text string, strictCleaning bool) []string {
	words := make([]string, 0)
	safeText := reg.ReplaceAllString(text, " ")
	for _, word := range strings.Split(safeText, " ") {
		if strictCleaning {
			if word == "-" {
				continue
			}
			word = strings.TrimFunc(word, func(character rune) bool {
				return character == '.' || character == '!' || character == ','
			})
			word = strings.ToLower(word)
		}
		words = append(words, word)
	}
	return words
}

func countWordsAndRepetition(words []string) []dimension {
	wordsAndRepetition := make([]dimension, 0)
	counter := make(map[string]int)
	for _, word := range words {
		counter[word]++
	}
	for value, key := range counter {
		d := dimension{
			word:  value,
			count: key,
		}
		wordsAndRepetition = append(wordsAndRepetition, d)
	}
	return wordsAndRepetition
}

func get10TopCountedWords(wordsAndRepetition []dimension) []string {
	sort.Slice(wordsAndRepetition, func(i, j int) bool {
		if wordsAndRepetition[i].count == wordsAndRepetition[j].count {
			return strings.Compare(wordsAndRepetition[i].word, wordsAndRepetition[j].word) < 0
		}
		return wordsAndRepetition[i].count > wordsAndRepetition[j].count
	})

	words := make([]string, 0)
	for i := 0; i < len(wordsAndRepetition) && i < 10; i++ {
		words = append(words, wordsAndRepetition[i].word)
	}

	return words
}
