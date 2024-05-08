package service

import "github.com/pdg-tw/go-monster-hearth-server/internal/translation/domain/translation/entity"

type Translator interface {
	Translate(translation entity.Translation) (entity.Translation, error)
}
