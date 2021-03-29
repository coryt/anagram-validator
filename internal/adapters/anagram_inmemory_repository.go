package adapters

import (
	"context"
	"fmt"

	"github.com/bluele/gcache"
	"github.com/coryt/anagram/internal/domain/anagram"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

const defaultTopListThreshold int = 10

type AnagramModel struct {
	Word    string
	Anagram string
}

type InMemoryAnagramRepository struct {
	cache gcache.Cache
	debug bool
}

func NewInMemoryAnagramRepository(debug bool) *InMemoryAnagramRepository {
	cache := gcache.New(defaultTopListThreshold).
		LFU().
		EvictedFunc(func(key, value interface{}) {
			if debug {
				fmt.Printf("evicted key: %s value: %d", key, value)
			}
		}).
		Build()
	return &InMemoryAnagramRepository{
		cache: cache,
		debug: debug,
	}
}

func (r *InMemoryAnagramRepository) GetTopAnagrams(ctx context.Context) (anagram.TopAnagrams, error) {
	emptyTopList := anagram.TopAnagrams{}
	top := r.cache.GetALL(true)
	if r.debug {
		spew.Dump(top)
	}

	cachedResults := make([]anagram.AnagramCount, 0)
	for key, value := range top {
		model := key.(AnagramModel)
		pair, err := unmarshalAnagram(model)
		if err != nil {
			return emptyTopList, err
		}
		anagramCount := anagram.AnagramCount{
			Pair:  *pair,
			Count: value.(int),
		}
		cachedResults = append(cachedResults, anagramCount)
	}

	return anagram.BuildTopAnagramList(cachedResults), nil
}

func (r *InMemoryAnagramRepository) SaveAnagram(ctx context.Context, pair *anagram.Pair) error {
	model := marshalAnagram(pair)
	value, err := r.cache.Get(model)
	if err == nil {
		if count, ok := value.(int); ok {
			r.cache.Set(model, count+1)
			return nil
		}
		return ErrSavingAnagram
	}
	err = r.cache.Set(model, 1)
	if err != nil {
		return ErrSavingAnagram
	}
	return nil
}

func marshalAnagram(a *anagram.Pair) AnagramModel {
	anagramModel := AnagramModel{
		Word:    a.Word,
		Anagram: a.Anagram,
	}

	return anagramModel
}

func unmarshalAnagram(key interface{}) (*anagram.Pair, error) {
	model, ok := key.(AnagramModel)
	if !ok {
		return nil, errors.Errorf("Adapters.unmarshalAnagram: unable to unmarshal AnagramModel:  %+v", key)
	}

	return anagram.NewAnagramPair(model.Word, model.Anagram), nil
}
