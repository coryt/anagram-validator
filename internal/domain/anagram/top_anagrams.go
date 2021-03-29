package anagram

type AnagramCount struct {
	Pair
	Count int
}

type TopAnagrams []AnagramCount

func BuildTopAnagramList(ac []AnagramCount) TopAnagrams {
	top := TopAnagrams{}
	top = append(top, ac...)
	return top
}
