package anagram_test

import (
	"testing"

	"github.com/coryt/anagram/internal/domain/anagram"
	"github.com/stretchr/testify/require"
)

func TestAnagram_ValidateAnagram_with_valid_ascii_anagram(t *testing.T) {
	expectedAnagram := anagram.NewAnagramPair("wolf", "flow")
	actualAnagram, err := anagram.ValidateAnagram("wolf", "flow")
	require.NoError(t, err)
	require.NotNil(t, actualAnagram)
	require.Equal(t, expectedAnagram, actualAnagram, "expected words to be an anagram")
}

func TestAnagram_ValidateAnagram_with_invalid_ascii_anagram(t *testing.T) {
	actualAnagram, err := anagram.ValidateAnagram("wolf", "owl")
	require.EqualError(t, err, anagram.ErrNotAValidAnagram.Error())
	require.Nil(t, actualAnagram)
}

func TestAnagram_ValidateAnagram_with_valid_utf8_anagram(t *testing.T) {
	expectedAnagram := anagram.NewAnagramPair("Ångström", "strömÅng")
	actualAnagram, err := anagram.ValidateAnagram("Ångström", "strömÅng")
	require.NoError(t, err)
	require.NotNil(t, actualAnagram)
	require.Equal(t, expectedAnagram, actualAnagram, "expected words to be an anagram")
}

func TestAnagram_ValidateAnagram_with_invalid_utf8_anagram(t *testing.T) {
	actualAnagram, err := anagram.ValidateAnagram("Ångström", "strömang")
	require.EqualError(t, err, anagram.ErrNotAValidAnagram.Error())
	require.Nil(t, actualAnagram)
}

func TestAnagram_ValidateAnagram_with_empty_words(t *testing.T) {
	actualAnagram, err := anagram.ValidateAnagram("", "")
	require.EqualError(t, err, anagram.ErrNotAValidAnagram.Error())
	require.Nil(t, actualAnagram)
}

func TestAnagram_ValidateAnagram_with_nonmatching_length(t *testing.T) {
	actualAnagram, err := anagram.ValidateAnagram("a", "ab")
	require.EqualError(t, err, anagram.ErrNotAValidAnagram.Error())
	require.Nil(t, actualAnagram)
}

func TestAnagram_ValidateAnagram_with_single_character(t *testing.T) {
	actualAnagram, err := anagram.ValidateAnagram("a", "a")
	require.EqualError(t, err, anagram.ErrNotAValidAnagram.Error())
	require.Nil(t, actualAnagram)
}

func TestAnagram_ValidateAnagram_with_spaces(t *testing.T) {
	actualAnagram, err := anagram.ValidateAnagram(" ", "  ")
	require.EqualError(t, err, anagram.ErrNotAValidAnagram.Error())
	require.Nil(t, actualAnagram)
}

func TestAnagram_ValidateAnagram_with_spaces_and_single_character(t *testing.T) {
	actualAnagram, err := anagram.ValidateAnagram(" a", " a ")
	require.EqualError(t, err, anagram.ErrNotAValidAnagram.Error())
	require.Nil(t, actualAnagram)
}

func TestAnagram_ValidateAnagram_with_mixed_case(t *testing.T) {
	actualAnagram, err := anagram.ValidateAnagram("aRtIc", "rictar")
	require.EqualError(t, err, anagram.ErrNotAValidAnagram.Error())
	require.Nil(t, actualAnagram)
}
