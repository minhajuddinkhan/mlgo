package nlp

import (
	"fmt"
	"strings"
)

type tokenizor struct {
	rules rules
}

type rules struct {
	Surnames []string
}

func NewTokenizor() *tokenizor {

	r := rules{[]string{"Dr"}}
	return &tokenizor{r}

}
func (t *tokenizor) Tokenize(blob string) []string {

	fp := strings.Split(blob, "\n")
	fmt.Println((fp))
	paragraphs := []string{}

	for _, v := range fp {

		if len(v) > 0 {
			paragraphs = append(paragraphs, v)
		}
	}

	return t.Words(paragraphs)

}

func (r *rules) R1(w string, v string) (yep bool) {

	yep = false
	for _, v := range r.Surnames {
		if v == w {
			yep = true
		}
	}
	return

}

func (t *tokenizor) Words(sentences []string) []string {

	r := t.rules
	words := []string{}

	for _, v := range sentences {
		w := ""
		for i := 0; i < len(v); i++ {
			if string(v[i]) == " " {
				if !r.R1(w, string(v[i])) {
					if len(w) <= 0 {
						continue
					}

					words = append(words, w)
					w = ""
					continue
				}
			}

			if string(v[i]) == "." { // if a dot appears ? check if the prefix is Dr or Mrs or something?
				if !r.R1(w, string(v[i])) {
					if len(w) <= 0 {
						continue
					}
					words = append(words, w)
					w = ""
					continue
				}
			}

			w += string(v[i])
			if i == len(v)-1 {
				words = append(words, w)
			}
		}
	}
	return words
}
