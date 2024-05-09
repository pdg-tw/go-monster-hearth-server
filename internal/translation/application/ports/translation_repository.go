package ports

import (
	"context"

	entity "github.com/pdg-tw/go-monster-hearth-server/internal/translation/domain/entity"
)

type TranslationRepository interface {
	Store(context.Context, entity.Translation) error
	GetHistory(context.Context) ([]entity.Translation, error)
}
