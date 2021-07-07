package audios

import (
	"github.com/car12o/audio-mastery/api/generated/models"
	"github.com/car12o/audio-mastery/api/generated/restapi/operations/audios"
	"github.com/go-openapi/runtime/middleware"
)

type PostAudioHandler struct{}

func NewPostAudioHandler() *PostAudioHandler {
	return &PostAudioHandler{}
}

func (a *PostAudioHandler) Handle(params audios.PostAudioParams, principal *models.Principal) middleware.Responder {
	return audios.NewPostAudioAccepted()
}
