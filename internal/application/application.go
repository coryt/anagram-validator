package application

import (
	"context"
	"fmt"

	"github.com/coryt/anagram/internal/adapters"
	"github.com/coryt/anagram/internal/application/command"
	"github.com/coryt/anagram/internal/application/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ValidateAnagram command.ValidateAnagramHandler
}

type Queries struct {
	TopAnagrams query.GetTopAnagramsHandler
}

func NewApplication(ctx context.Context, debugMode bool) Application {
	fmt.Printf("Initializing Application with debugMode: %t\n", debugMode)

	anagramRepository := adapters.NewInMemoryAnagramRepository(debugMode)

	return Application{
		Commands: Commands{
			ValidateAnagram: command.NewValidateAnagramHandler(anagramRepository),
		},
		Queries: Queries{
			TopAnagrams: query.NewGetTopAnagramsHandler(anagramRepository),
		},
	}
}
