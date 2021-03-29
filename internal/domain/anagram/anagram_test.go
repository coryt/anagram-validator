package anagram_test

import (
	"testing"

	"github.com/coryt/anagram/internal/domain/anagram"
	"github.com/stretchr/testify/assert"
)

func TestNewAnagramPair(t *testing.T) {
	expectedPair := &anagram.Pair{Word: "wolf", Anagram: "folw"}
	actualPair := anagram.NewAnagramPair("wolf", "folw")
	assert.Equal(t, expectedPair, actualPair)
}
