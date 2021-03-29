package anagram

import (
	"context"
)

type Repository interface {
	GetTopAnagrams(ctx context.Context) (TopAnagrams, error)
	SaveAnagram(ctx context.Context, anagram *Pair) error
}
