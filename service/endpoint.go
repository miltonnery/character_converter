package service

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"miltonnery/character_converter/dto"
)

func NewServiceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//Parsing the request
		serviceRequest := request.(dto.Request)

		//Executing the service
		serviceResponse, err := s.Translate(&serviceRequest)

		response := dto.InternalServiceResponse{
			Response: serviceResponse,
			Err:      err,
		}

		//Sending the response to the transport layer
		return response, nil
	}
}

type Endpoints struct {
	Service endpoint.Endpoint
}
