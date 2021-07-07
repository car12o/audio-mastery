package audiomastery

import (
	"github.com/car12o/audio-mastery/internal/audio"
	"github.com/car12o/audio-mastery/internal/grammarian"
	"github.com/car12o/audio-mastery/pkg/logger"
)

type App struct {
	Log      logger.Service
	Services Services
}

type Services struct {
	Audio audio.Service
}

func NewApp() *App {
	log := logger.New()

	audioService := audio.NewService(
		audio.NewRepository(),
		grammarian.NewService(),
	)

	return &App{
		Log: log,
		Services: Services{
			Audio: audioService,
		},
	}
}
