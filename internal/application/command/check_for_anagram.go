package command

import (
	"context"
	"fmt"

	"github.com/coryt/anagram/internal/adapters"
	"github.com/coryt/anagram/internal/domain/anagram"
)

type PossibleAnagrams struct {
	FirstWord  string
	SecondWord string
}

type ValidateAnagramHandler struct {
	anagramRepository anagram.Repository
}

func NewValidateAnagramHandler(anagramRepository anagram.Repository) ValidateAnagramHandler {
	if anagramRepository == nil {
		panic("nil anagramRepository")
	}

	return ValidateAnagramHandler{anagramRepository: anagramRepository}
}

func (h ValidateAnagramHandler) Handle(ctx context.Context, possibleAnagram PossibleAnagrams) error {
	pair, err := anagram.ValidateAnagram(possibleAnagram.FirstWord, possibleAnagram.SecondWord)
	if err != nil {
		return err
	}

	if err := h.anagramRepository.SaveAnagram(ctx, pair); err != nil {
		return fmt.Errorf("ValidateAnagramHandler.Handle: %v %w", err, adapters.ErrSavingAnagram)
	}

	return nil
}
