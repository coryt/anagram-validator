package anagram

import (
	"strings"
	"unicode/utf8"
)

type AnagramValidationRule = func(firstWord, secondWord string) error

func ValidateAnagram(firstWord, secondWord string) (*Pair, error) {
	validationRules := []AnagramValidationRule{
		pairCannotBeEmpty,
		pairMustBeEqualLength,
		pairCannotBeSingleCharacter,
		pairMustBeAnagrams,
	}

	for _, ruleFn := range validationRules {
		err := ruleFn(firstWord, secondWord)
		if err != nil {
			return nil, err
		}
	}
	return NewAnagramPair(firstWord, secondWord), nil
}

func pairCannotBeSingleCharacter(firstWord, secondWord string) error {
	if utf8.RuneCountInString(firstWord) == 1 || utf8.RuneCountInString(secondWord) == 1 {
		return ErrNotAValidAnagram
	}
	return nil
}

func pairCannotBeEmpty(firstWord, secondWord string) error {
	if strings.TrimSpace(firstWord) == "" || strings.TrimSpace(secondWord) == "" {
		return ErrNotAValidAnagram
	}
	return nil
}

func pairMustBeEqualLength(firstWord, secondWord string) error {
	if utf8.RuneCountInString(firstWord) != utf8.RuneCountInString(secondWord) {
		return ErrNotAValidAnagram
	}
	return nil
}

func pairMustBeAnagrams(firstWord, secondWord string) error {
	freq := make(map[rune]int, 0)
	for _, runeValue := range firstWord {
		freq[runeValue]++
	}
	for _, runeValue := range secondWord {
		if freq[runeValue] == 0 {
			return ErrNotAValidAnagram
		}
		freq[runeValue]--
	}
	return nil
}
