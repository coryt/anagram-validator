package ports

import (
	"github.com/coryt/anagram/internal/application"
	"github.com/coryt/anagram/internal/application/command"
	"github.com/gofiber/fiber/v2"
)

func GetTopAnagrams(application application.Application) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		top, err := application.Queries.TopAnagrams.Handle(ctx.Context())
		if err != nil {
			return ctx.JSON(map[string]interface{}{
				"message": err.Error(),
			})
		}
		list := make([]AnagramCount, 0)
		for _, value := range top {
			list = append(list, AnagramCount{value.Word, value.Anagram, value.Count})
		}
		resp := BuildTopAnagramResponse(list)
		return ctx.JSON(resp)
	}
}

func ValidateAnagram(application application.Application) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		reqModel := new(ValidateAnagramRequest)
		if err := ctx.BodyParser(reqModel); err != nil {
			return fiber.ErrBadRequest
		}

		cmd := command.PossibleAnagrams{
			FirstWord:  reqModel.FirstWord,
			SecondWord: reqModel.SecondWord,
		}
		err := application.Commands.ValidateAnagram.Handle(ctx.Context(), cmd)
		if err != nil {
			return ctx.JSON(&ValidateAnagramResponse{
				Valid: false,
			})
		}
		return ctx.JSON(&ValidateAnagramResponse{
			Valid: true,
		})
	}
}
