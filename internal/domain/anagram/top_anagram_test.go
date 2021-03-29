package anagram_test

import (
	"testing"

	"github.com/coryt/anagram/internal/domain/anagram"
	"github.com/stretchr/testify/assert"
)

func TestBuildTopAnagramList(t *testing.T) {
	expectedTopAnagrams := []anagram.AnagramCount{
		{
			Pair:  anagram.Pair{Word: "wolf", Anagram: "folw"},
			Count: 10,
		},
		{
			Pair:  anagram.Pair{Word: "bob", Anagram: "obb"},
			Count: 5,
		},
		{
			Pair:  anagram.Pair{Word: "artic", Anagram: "rictar"},
			Count: 1,
		},
	}

	actualTopAnagrams := anagram.BuildTopAnagramList(expectedTopAnagrams)
	assert.Len(t, actualTopAnagrams, len(expectedTopAnagrams))
	assert.ElementsMatch(t, actualTopAnagrams, expectedTopAnagrams)
}
