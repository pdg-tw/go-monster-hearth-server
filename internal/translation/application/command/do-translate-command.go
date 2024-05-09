package command

import (
	"github.com/pdg-tw/go-monster-hearth-server/internal/translation/application/ports"
	"github.com/pdg-tw/go-monster-hearth-server/internal/translation/domain/entity"
	"github.com/pdg-tw/go-monster-hearth-server/internal/translation/domain/service"
)

type TranslateCommand struct {
	Translator            service.Translator
	TranslationRepository ports.TranslationRepository
	Translation           entity.Translation
}
