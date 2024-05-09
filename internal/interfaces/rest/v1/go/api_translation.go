/*
 * Go Clean Template API
 *
 * Using a translation service as an example
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/pdg-tw/go-monster-hearth-server/internal/translation/application"
	entity "github.com/pdg-tw/go-monster-hearth-server/internal/translation/domain/entity"
	"github.com/pdg-tw/go-monster-hearth-server/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Translator struct {
	translationUseCase *application.TranslationUseCase
	log                *logger.Logger
}

var singletonTranslator *Translator

func NewTranslator(translationUseCase *application.TranslationUseCase, log *logger.Logger) *Translator {
	singletonTranslator = &Translator{
		translationUseCase: translationUseCase,
		log:                log,
	}
	return singletonTranslator
}

func DoTranslate(c *gin.Context) {
	singletonTranslator.DoTranslate(c)
}

func History(c *gin.Context) {
	singletonTranslator.History(c)
}

func (t *Translator) DoTranslate(c *gin.Context) {

	log := t.log
	translationUseCase := t.translationUseCase

	var request TranslateRequestObject
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	translation, err := translationUseCase.Translate(
		c.Request.Context(),
		entity.Translation{
			Source:      request.Source,
			Destination: request.Destination,
			Original:    request.Original,
		},
	)

	if err != nil {
		log.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusInternalServerError, "translation service problems")
		return
	}

	translationResponseObject := translationToResponseObject(translation)

	c.JSON(http.StatusOK, translationResponseObject)
}

func (t *Translator) History(c *gin.Context) {

	log := t.log
	translationUseCase := t.translationUseCase

	translations, err := translationUseCase.History(c.Request.Context())
	if err != nil {
		log.Error(err, "http - v1 - history")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	translationResponseObjects := translationsToResponseObjects(translations)

	c.JSON(http.StatusOK, HistoryResponseObject{
		History: translationResponseObjects,
	})
}

func translationsToResponseObjects(translations []entity.Translation) []TranslationResponseObject {
	var translationResponseObjects = []TranslationResponseObject{}

	for _, translation := range translations {
		translationResponseObjects = append(translationResponseObjects, translationToResponseObject(translation))
	}
	return translationResponseObjects
}

func translationToResponseObject(translation entity.Translation) TranslationResponseObject {
	return TranslationResponseObject{
		Destination: translation.Destination,
		Original:    translation.Original,
		Source:      translation.Source,
		Translation: translation.Translation,
	}
}
