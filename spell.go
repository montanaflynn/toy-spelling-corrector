package toyspellingcorrector

import (
	"io/ioutil"
	"regexp"
	"strings"
)

// ToySpellcheck holds the words used for finding corrections
type ToySpellcheck struct {
	words map[string]int
}

// Train takes a text file and splits it into words and places
// the words into the ToySpellcheck struct words dictionary
func (s *ToySpellcheck) Train(path string) {
	s.words = make(map[string]int)
	pattern := regexp.MustCompile("[a-z]+")
	content, err := ioutil.ReadFile(path)
	if err == nil {
		lc := strings.ToLower(string(content))
		for _, w := range pattern.FindAllString(lc, -1) {
			s.words[w]++
		}
	}
}

// edits1 runs a series of mutations against a given word
// including deletions, transpositions, alterations and
// insertions totalling 54n+25 where n is len(word)
func (s *ToySpellcheck) edits1(word string) (out []string) {

	// break up the alphabet into a slice of characters
	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz", "")

	// Delete single letters from word
	for i := 0; i < len(word); i++ {
		out = append(out, word[:i]+word[i+1:])
	}

	// Transpose letters meaning swap adjacent letters
	for i := 0; i < len(word)-1; i++ {
		out = append(out, word[:i]+word[i+1:i+2]+word[i:i+1]+word[i+2:])
	}

	// Alter letters meaning replace letter with another
	for i := 0; i < len(word); i++ {
		for _, char := range alphabet {
			out = append(out, word[:i]+char+word[i+1:])
		}
	}

	// Insert letters
	for i := 0; i <= len(word); i++ {
		for _, char := range alphabet {
			out = append(out, word[:i]+char+word[i:])
		}
	}

	return out
}

// Run a second round of edits against output of first round
// and limit the output to only known words from the dictionary
func (s *ToySpellcheck) knownEdits2(word string) (out []string) {
	firstRound := s.edits1(word)
	for _, variation := range firstRound {
		secondRound := s.edits1(variation)
		out = append(out, s.known(secondRound)...)
	}
	return out
}

// Find any known words in a slice and return them
func (s *ToySpellcheck) known(words []string) []string {
	out := []string{}
	for _, word := range words {
		_, found := s.words[word]
		if found {
			out = append(out, word)
		}
	}
	return out
}

// Correct tries to find the best correct spelling of
// a given word, falls back to giving the same word if
// none can be found
func (s *ToySpellcheck) Correct(word string) string {

	if _, found := s.words[word]; found {
		return word
	}

	firstRounnd := s.known(s.edits1(word))
	if result := s.bestCandidate(firstRounnd); result != "" {
		return result
	}

	secondRound := s.knownEdits2(word)
	if result := s.bestCandidate(secondRound); result != "" {
		return result
	}

	return word
}

// bestCandidate returns the word with the highest count of use in
// the word list or simply the first word if there is only one
func (s *ToySpellcheck) bestCandidate(words []string) (result string) {
	if len(words) > 0 {
		if len(words) == 1 {
			return words[0]
		}
		highCount := 0
		for _, word := range words {
			if s.words[word] > highCount {
				highCount = s.words[word]
				result = word
			}
		}
		return result
	}
	return ""
}
