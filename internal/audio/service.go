package audio

import (
	"log"

	"github.com/car12o/audio-mastery/internal/grammarian"
)

type service struct {
	repository Repository
	grammarian grammarian.Service
}

func NewService(repository Repository, grammarian grammarian.Service) Service {
	return &service{repository, grammarian}
}

func (s *service) AddAudio(audio *Audio) error {
	// validate necessary audio fields
	// and return error in case something is not valid
	// otherwise process and store audio async
	go s.processAndStore(audio)
	return nil
}

func (s *service) GetAudio(id string) (*Audio, error) {
	return nil, nil
}

func (s *service) ListAudios() ([]*Audio, error) {
	return nil, nil
}

func (s *service) UpdateAudio(audio *Audio) error {
	return nil
}

func (s *service) DeleteAudio(id string) error {
	return nil
}

func (s *service) processAndStore(audio *Audio) {
	improvedRedord, err := s.grammarian.Process(audio.Record)
	if err != nil {
		log.Println(err)
		// notify user that audit failed to be proccessed
		return
	}

	audio.Record = improvedRedord

	if _, err := s.repository.AddAudio(audio); err != nil {
		log.Println(err)
		// notify user that audit failed to be stored
	}
}
