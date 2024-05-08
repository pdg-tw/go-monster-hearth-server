package ports

import (
	"context"

	"github.com/pdg-tw/go-monster-hearth-server/internal/translation/domain/translation/entity"
)

type TranslationRepository interface {
	Store(context.Context, entity.Translation) error
	GetHistory(context.Context) ([]entity.Translation, error)
}
