package anagram

type Pair struct {
	Word    string
	Anagram string
}

func NewAnagramPair(word, anagram string) *Pair {
	return &Pair{word, anagram}
}
