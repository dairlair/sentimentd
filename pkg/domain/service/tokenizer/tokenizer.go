package tokenizer

import (
	"regexp"
	"strings"
)

type tokenizer struct {

}

func NewTokenizer() tokenizer {
	return tokenizer{}
}

func (t *tokenizer) Tokenize(sentence string) []string {
	return tokenize(sentence)
}

// meaninglessWords are words which have very little meaning
var meaninglessWords = map[string]struct{}{
	"i": struct{}{}, "me": struct{}{}, "my": struct{}{}, "myself": struct{}{}, "we": struct{}{}, "our": struct{}{}, "ours": struct{}{},
	"ourselves": struct{}{}, "you": struct{}{}, "your": struct{}{}, "yours": struct{}{}, "yourself": struct{}{}, "yourselves": struct{}{},
	"he": struct{}{}, "him": struct{}{}, "his": struct{}{}, "himself": struct{}{}, "she": struct{}{}, "her": struct{}{}, "hers": struct{}{},
	"herself": struct{}{}, "it": struct{}{}, "its": struct{}{}, "itself": struct{}{}, "they": struct{}{}, "them": struct{}{}, "their": struct{}{},
	"theirs": struct{}{}, "themselves": struct{}{}, "what": struct{}{}, "which": struct{}{}, "who": struct{}{}, "whom": struct{}{}, "this": struct{}{},
	"that": struct{}{}, "these": struct{}{}, "those": struct{}{}, "am": struct{}{}, "is": struct{}{}, "are": struct{}{}, "was": struct{}{},
	"were": struct{}{}, "be": struct{}{}, "been": struct{}{}, "being": struct{}{}, "have": struct{}{}, "has": struct{}{}, "had": struct{}{},
	"having": struct{}{}, "do": struct{}{}, "does": struct{}{}, "did": struct{}{}, "doing": struct{}{}, "a": struct{}{}, "an": struct{}{},
	"the": struct{}{}, "and": struct{}{}, "but": struct{}{}, "if": struct{}{}, "or": struct{}{}, "because": struct{}{}, "as": struct{}{},
	"until": struct{}{}, "while": struct{}{}, "of": struct{}{}, "at": struct{}{}, "by": struct{}{}, "for": struct{}{}, "with": struct{}{},
	"about": struct{}{}, "against": struct{}{}, "between": struct{}{}, "into": struct{}{}, "through": struct{}{}, "during": struct{}{},
	"before": struct{}{}, "after": struct{}{}, "above": struct{}{}, "below": struct{}{}, "to": struct{}{}, "from": struct{}{}, "up": struct{}{},
	"down": struct{}{}, "in": struct{}{}, "out": struct{}{}, "on": struct{}{}, "off": struct{}{}, "over": struct{}{}, "under": struct{}{},
	"again": struct{}{}, "further": struct{}{}, "then": struct{}{}, "once": struct{}{}, "here": struct{}{}, "there": struct{}{}, "when": struct{}{},
	"where": struct{}{}, "why": struct{}{}, "how": struct{}{}, "all": struct{}{}, "any": struct{}{}, "both": struct{}{}, "each": struct{}{},
	"few": struct{}{}, "more": struct{}{}, "most": struct{}{}, "other": struct{}{}, "some": struct{}{}, "such": struct{}{}, "no": struct{}{},
	"nor": struct{}{}, "not": struct{}{}, "only": struct{}{}, "same": struct{}{}, "so": struct{}{}, "than": struct{}{}, "too": struct{}{},
	"very": struct{}{}, "can": struct{}{}, "will": struct{}{}, "just": struct{}{}, "don't": struct{}{}, "should": struct{}{}, "should've": struct{}{},
	"now": struct{}{}, "aren't": struct{}{}, "couldn't": struct{}{}, "didn't": struct{}{}, "doesn't": struct{}{}, "hasn't": struct{}{}, "haven't": struct{}{},
	"isn't": struct{}{}, "shouldn't": struct{}{}, "wasn't": struct{}{}, "weren't": struct{}{}, "won't": struct{}{}, "wouldn't": struct{}{},
}

func isMeaninglessWord(w string) bool {
	_, ok := meaninglessWords[w]
	return ok
}

// Cleanup remove none-alphanumeric characters and convert them to lowercase
func cleanup(sentence string) string {
	re := regexp.MustCompile("[^a-zA-Z 0-9]+")
	return re.ReplaceAllString(strings.ToLower(sentence), "")
}

// Tokenize create an array of words from a sentence
func tokenize(sentence string) []string {
	s := cleanup(sentence)
	words := strings.Fields(s)
	var tokens []string
	for _, w := range words {
		if !isMeaninglessWord(w) {
			tokens = append(tokens, w)
		}
	}
	return tokens
}
