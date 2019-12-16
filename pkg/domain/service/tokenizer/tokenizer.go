package tokenizer

import "strings"

type tokenizer struct {

}

func NewTokenizer() tokenizer {
	return tokenizer{}
}

func (t *tokenizer) Tokenize(sentence string) []string {
	return strings.Split(sentence, " ")
}