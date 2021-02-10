package service

import (
	"miltonnery/character_converter/configuration"
	"miltonnery/character_converter/dto"
	"miltonnery/character_converter/log"
	"miltonnery/character_converter/translator"
	"strings"
)

type Service interface {
	Translate(serviceRequest *dto.Request) (serviceResponse *dto.Response, err error)
}

// SERVICE INTERFACE IMPLEMENTATION ------------------------------------------------------------------------------------/

const (
	TEXT   = "TEXT"
	BINARY = "BINARY"
	MORSE  = "MORSE"
)

type CharacterTranslator struct {
	configuration configuration.Configuration
	logger        log.SeriviceLog
}

func NewImpl(configuration configuration.Configuration, logger log.SeriviceLog) *CharacterTranslator {
	return &CharacterTranslator{
		configuration: configuration,
		logger:        logger,
	}
}

func (ct CharacterTranslator) Translate(serviceRequest *dto.Request) (serviceResponse *dto.Response, err error) {
	ct.logger.InfoLite("Service", "Starting the service translation layer")

	var t translator.Translator
	input := serviceRequest.TextoATraducir
	destino := serviceRequest.FormatoDestino

	switch serviceRequest.FormatoOrigen {
	case TEXT:
		{
			if strings.EqualFold(destino, BINARY) {
				t = translator.NewTextBinaryAdapter(input)
				break
			}
			if strings.EqualFold(destino, MORSE) {
				t = translator.NewTextMorseAdapter(input)
			}
		}
	case BINARY:
		{
			if strings.EqualFold(destino, TEXT) {
				t = translator.NewBinaryTextAdapter(input)
				break
			}
			if strings.EqualFold(destino, MORSE) {
				t = translator.NewBinaryMorseAdapter(input)
				break
			}
		}

	case MORSE:
		{
			if strings.EqualFold(destino, TEXT) {
				t = translator.NewMorseTextAdapter(input)
				break
			}
			if strings.EqualFold(destino, BINARY) {
				t = translator.NewMorseBinaryAdapter(input)
			}
		}
	}

	//Send to translate
	output, translationErr := t.Translate()
	if err != nil {
		err = translationErr
		return nil, err
	}

	serviceResponse = &dto.Response{TextoTraducido: output}
	return
}

// Auxiliary functions -------------------------------------------------------------------------------------------------/
