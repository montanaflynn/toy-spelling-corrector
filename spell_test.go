package toyspellingcorrector

import (
	"fmt"
	"testing"
	"time"
)

var spellcheck *ToySpellcheck

func init() {
	spellcheck = &ToySpellcheck{}
	spellcheck.Train("./big.txt")
}

func TestEdits1(t *testing.T) {
	variations := spellcheck.edits1("something")
	if len(variations) != 511 {
		t.Fatal("Found variations:", len(variations), "!= 511")
	}
}

type results struct {
	n       int
	bad     int
	unknown int
	time    time.Duration
}

func TestSpell(t *testing.T) {
	c := developmentCorpus()
	testCorrections(c.tests1, t)
	testCorrections(c.tests2, t)
}

func testCorrections(words map[string][]string, t *testing.T) {
	r := results{}
	start := time.Now()
	for target, wrongs := range words {
		for _, wrong := range wrongs {
			r.n++
			w := spellcheck.Correct(wrong)
			if w != target {
				r.bad++
				if _, found := spellcheck.words[target]; !found {
					r.unknown++
				}
			}
		}
	}

	duration := time.Since(start)
	msg := "n: %d, bad: %d, unknown: %d, pct: %d, secs: %v\n"
	pct := int(100 - (float64(r.bad) / float64(r.n) * 100))
	fmt.Printf(msg, r.n, r.bad, r.unknown, pct, duration)
}
