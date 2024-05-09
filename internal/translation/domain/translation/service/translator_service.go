package service

import entity "github.com/pdg-tw/go-monster-hearth-server/internal/translation/domain"

type Translator interface {
	Translate(translation entity.Translation) (entity.Translation, error)
}
