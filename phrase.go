package pinyin

import (
	"strings"
)

func cutWords(p string) []string {
	s := make([]string, 0)
	for _, r := range p {
		s = append(s, string(r))
	}
	return s
}

func pinyinPhrase(s string) string {
	words := cutWords(s)
	for _, word := range words {
		match := phraseDict[word]
		if match == "" {
			match = phraseDictAddition[word]
		}

		match = toFixed(match, paragraphOption)
		if match != "" {
			s = strings.Replace(s, word, " "+match+" ", 1)
		}
	}

	return s
}
