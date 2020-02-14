package service

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	DoStuffEndpoint endpoint.Endpoint
}

func (e Endpoints) DoStuff(ctx context.Context, value int) (int, error) {
	resp, err := e.DoStuffEndpoint(ctx, DoStuffRequest{Value: value})
	if err != nil {
		return 0, err
	}

	response, ok := resp.(DoStuffResponse)
	if !ok {
		return 0, errors.New("Invalid response structure")
	}

	return response.Result, response.Err
}

func MakeEndpoints(p Service) Endpoints {
	return Endpoints{
		DoStuffEndpoint: MakeDoStuffEndpoint(p),
	}
}

func MakeDoStuffEndpoint(p Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(DoStuffRequest)
		if !ok {
			return nil, errors.New("Invalid request structure")
		}

		v, err := p.DoStuff(ctx, req.Value)
		return DoStuffResponse{v, err}, nil
	}
}
