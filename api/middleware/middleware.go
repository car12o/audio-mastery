package middleware

import (
	"net/http"

	"github.com/car12o/audio-mastery/api/generated/restapi/operations"
	"github.com/car12o/audio-mastery/pkg/logger"
)

func Load(api *operations.AudioMasteryAPI, log logger.Service) http.Handler {
	setupAuths(api)
	handler := api.Serve(nil)
	requestLog := NewRequestLog(log)
	return requestLog.Handle(handler)
}
