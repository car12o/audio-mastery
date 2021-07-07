package api

import (
	"fmt"

	"github.com/car12o/audio-mastery/api/generated/restapi"
	"github.com/car12o/audio-mastery/api/generated/restapi/operations"
	"github.com/car12o/audio-mastery/api/middleware"
	"github.com/car12o/audio-mastery/pkg/logger"
	"github.com/go-openapi/loads"
)

func NewServer(host string, port int, log logger.Service) (*restapi.Server, error) {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "2.0")
	if err != nil {
		return nil, err
	}

	api := operations.NewAudioMasteryAPI(swaggerSpec)
	api.Logger = func(s string, i ...interface{}) {
		log.Info(fmt.Sprintf(s, i...))
	}
	handler := middleware.Load(api, log)

	server := restapi.NewServer(api)
	server.Host = host
	server.Port = port
	server.ConfigureAPI()
	server.SetHandler(handler)

	return server, nil
}
