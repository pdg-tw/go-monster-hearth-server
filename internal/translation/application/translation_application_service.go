package application

import (
	"context"
	"fmt"

	"github.com/pdg-tw/go-monster-hearth-server/internal/translation/application/command"
	"github.com/pdg-tw/go-monster-hearth-server/internal/translation/application/ports"
	entity "github.com/pdg-tw/go-monster-hearth-server/internal/translation/domain/entity"
	"github.com/pdg-tw/go-monster-hearth-server/internal/translation/domain/service"
)

// TranslationUseCase -.
type TranslationUseCase struct {
	translationRepository ports.TranslationRepository
	translator            service.Translator
}

func NewWithDependencies(translationRepository ports.TranslationRepository, translator service.Translator) *TranslationUseCase {
	return &TranslationUseCase{
		translationRepository: translationRepository,
		translator:            translator,
	}
}

// History - getting translate history from store.
func (uc *TranslationUseCase) History(ctx context.Context) ([]entity.Translation, error) {
	translations, err := uc.translationRepository.GetHistory(ctx)
	if err != nil {
		return nil, fmt.Errorf("TranslationUseCase - History - s.translationRepository.GetHistory: %w", err)
	}

	return translations, nil
}

// Translate -.
func (uc *TranslationUseCase) Translate(ctx context.Context, t entity.Translation) (entity.Translation, error) {
	cmd := &command.TranslateCommand{
		Translator:            uc.translator,
		TranslationRepository: uc.translationRepository,
		Translation:           t,
	}

	if err := cmd.Handle(ctx); err != nil {
		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - cmd.Execute: %w", err)
	}

	return t, nil
}
