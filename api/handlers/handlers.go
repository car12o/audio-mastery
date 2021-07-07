package handlers

import (
	"github.com/car12o/audio-mastery/api/generated/restapi/operations"
	"github.com/car12o/audio-mastery/api/handlers/audios"
)

func Load(api *operations.AudioMasteryAPI) {
	// audios handlers
	api.AudiosPostAudioHandler = audios.NewPostAudioHandler()
}
