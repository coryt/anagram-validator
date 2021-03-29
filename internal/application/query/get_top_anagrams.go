package query

import (
	"context"

	"github.com/coryt/anagram/internal/domain/anagram"
)

type GetTopAnagramsHandler struct {
	anagramRepository anagram.Repository
}

func NewGetTopAnagramsHandler(anagramRepository anagram.Repository) GetTopAnagramsHandler {
	if anagramRepository == nil {
		panic("nil anagramRepository")
	}

	return GetTopAnagramsHandler{anagramRepository: anagramRepository}
}

func (h GetTopAnagramsHandler) Handle(ctx context.Context) (anagram.TopAnagrams, error) {
	return h.anagramRepository.GetTopAnagrams(ctx)
}
