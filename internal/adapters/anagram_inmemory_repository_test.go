package adapters_test

import (
	"context"
	"testing"

	"github.com/coryt/anagram/internal/adapters"
	"github.com/coryt/anagram/internal/domain/anagram"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const debugModeDisabled = false

func TestNewInMemoryAnagramRepository_SaveAnagram(t *testing.T) {
	testCases := map[string]struct {
		expectedAnagramCount anagram.AnagramCount
		fixtureSetup         func() *adapters.InMemoryAnagramRepository
	}{
		"New Pair Has Count of 1": {
			expectedAnagramCount: anagram.AnagramCount{
				Pair:  anagram.Pair{Word: "wolf", Anagram: "folw"},
				Count: 1,
			},
			fixtureSetup: func() *adapters.InMemoryAnagramRepository {
				return adapters.NewInMemoryAnagramRepository(debugModeDisabled)
			},
		},
		"Same Pair Increments Count to 2": {
			expectedAnagramCount: anagram.AnagramCount{
				Pair:  anagram.Pair{Word: "wolf", Anagram: "folw"},
				Count: 2,
			},
			fixtureSetup: func() *adapters.InMemoryAnagramRepository {
				repo := adapters.NewInMemoryAnagramRepository(debugModeDisabled)
				repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "wolf", Anagram: "folw"})
				return repo
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			repo := tc.fixtureSetup()

			expectedAnagram := tc.expectedAnagramCount

			err := repo.SaveAnagram(ctx, &expectedAnagram.Pair)
			require.NoError(t, err)

			assertPersistedAnagramEquals(t, repo, expectedAnagram)
		})
	}
}

func assertPersistedAnagramEquals(t *testing.T, repo *adapters.InMemoryAnagramRepository, anagramCount anagram.AnagramCount) {
	topAnagrams, err := repo.GetTopAnagrams(context.Background())
	require.NoError(t, err)
	assert.Contains(t, topAnagrams, anagramCount)
	for _, item := range topAnagrams {
		if item == anagramCount {
			assert.Equal(t, anagramCount.Pair, item.Pair)
			return
		}
	}
	assert.Fail(t, "AnagramCount was not found in top list")
}

func TestNewInMemoryAnagramRepository_MaxAnagramsKept(t *testing.T) {
	repo := adapters.NewInMemoryAnagramRepository(debugModeDisabled)
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "wolf", Anagram: "folw"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "wolf", Anagram: "folw"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "wow", Anagram: "oww"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "wow", Anagram: "oww"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "bat", Anagram: "tab"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "bat", Anagram: "tab"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "cat", Anagram: "tac"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "cat", Anagram: "tac"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "mat", Anagram: "tam"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "mat", Anagram: "tam"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "sad", Anagram: "das"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "sad", Anagram: "das"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "knee", Anagram: "keen"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "knee", Anagram: "keen"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "dog", Anagram: "god"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "dog", Anagram: "god"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "ricat", Anagram: "artic"})
	repo.SaveAnagram(context.Background(), &anagram.Pair{Word: "ricat", Anagram: "artic"})

	expectedAnagramCountToBeEvicted := anagram.AnagramCount{
		Pair:  anagram.Pair{Word: "f1", Anagram: "1f"},
		Count: 1,
	}
	repo.SaveAnagram(context.Background(), &expectedAnagramCountToBeEvicted.Pair)

	ctx := context.Background()
	err := repo.SaveAnagram(ctx, &anagram.Pair{Word: "f1", Anagram: "1f"})
	require.NoError(t, err)

	assertPersistedAnagramNotEquals(t, repo, expectedAnagramCountToBeEvicted)
}

func assertPersistedAnagramNotEquals(t *testing.T, repo *adapters.InMemoryAnagramRepository, anagramCount anagram.AnagramCount) {
	topAnagrams, err := repo.GetTopAnagrams(context.Background())
	require.NoError(t, err)
	assert.NotContains(t, topAnagrams, anagramCount)
}
