package command

import (
	"context"
	"fmt"
)

func (cmd *TranslateCommand) Handle(ctx context.Context) error {
	translation, err := cmd.Translator.Translate(cmd.Translation)
	fmt.Println(translation)
	if err != nil {
		return err
	}

	err = cmd.TranslationRepository.Store(ctx, translation)
	if err != nil {
		return err
	}

	return nil
}
