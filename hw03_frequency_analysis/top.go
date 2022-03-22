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

// Change to true if needed.
var taskWithAsteriskIsCompleted = true
var reg = regexp.MustCompile(`\s+`)

func Top10(text string) []string {
	if len(text) <= 0 {
		return make([]string, 0)
	}

	counter := make(map[string]int)

	safe := reg.ReplaceAllString(text, " ")

	for _, v := range strings.Split(safe, " ") {

		if taskWithAsteriskIsCompleted {
			if v == "-" {
				continue
			}
			v = strings.Trim(v, ".")
			v = strings.ToLower(v)
		}

		counter[v]++
	}

	ds := make([]dimension, 0)
	for value, key := range counter {
		d := dimension{
			word:  value,
			count: key,
		}
		ds = append(ds, d)
	}

	sort.Slice(ds, func(i, j int) bool {
		if ds[i].count == ds[j].count {
			return strings.Compare(ds[i].word, ds[j].word) < 0
		}
		return ds[i].count > ds[j].count
	})

	words := make([]string, 0)

	// fmt.Println(ds)

	for i := 0; i < len(ds) && i < 10; i++ {
		words = append(words, ds[i].word)
	}

	return words
}
