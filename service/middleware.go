package service

import (
	"context"
	"miltonnery/character_converter/configuration"
	"miltonnery/character_converter/dto"
	"miltonnery/character_converter/log"
	"miltonnery/character_converter/validator"
)

type middlewareImpl struct {
	configuration configuration.Configuration
	logger        log.SeriviceLog
	next          Service
}

func NewMiddlewareImpl(configuration configuration.Configuration, logger log.SeriviceLog, serviceImpl Service) *middlewareImpl {
	return &middlewareImpl{
		configuration: configuration,
		logger:        logger,
		next:          serviceImpl,
	}
}

func (m middlewareImpl) Translate(serviceRequest *dto.Request) (serviceResponse *dto.Response, err error) {

	ctx := context.Background()
	log.FillContextForLogs(ctx)
	m.logger.InfoLite("Middleware", "Beginning of middleware layer")

	//Request validation
	m.logger.InfoLite("Middleware", "Request content validation")
	rv := validator.NewRequestValidator(serviceRequest)
	err = rv.OK()
	if err != nil {
		return
	}

	//TODO: Set the middleware logic here
	serviceResponse, err = m.next.Translate(serviceRequest)
	if err != nil {
		return nil, err
	}
	return
}
