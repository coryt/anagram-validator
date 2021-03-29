package ports

import "sort"

type ValidateAnagramResponse struct {
	Valid bool `json:"valid"`
}

type TopAnagramsResponse []AnagramCount

func (t TopAnagramsResponse) Len() int {
	return len(t)
}

func (t TopAnagramsResponse) Less(i, j int) bool {
	return t[i].Count > t[j].Count
}

func (t TopAnagramsResponse) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type AnagramCount struct {
	Word    string `json:"word"`
	Anagram string `json:"anagram"`
	Count   int    `json:"count"`
}

func BuildTopAnagramResponse(ac []AnagramCount) TopAnagramsResponse {
	resp := TopAnagramsResponse{}
	resp = append(resp, ac...)
	sort.Sort(resp)
	return resp
}
