package audiomastery

import "github.com/car12o/audio-mastery/pkg/logger"

type App struct {
	Log logger.Service
}

func NewApp() *App {
	log := logger.New()
	return &App{log}
}
