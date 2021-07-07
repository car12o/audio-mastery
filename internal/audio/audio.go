package audio

import "time"

type Audio struct {
	ID        string
	User      string
	Record    []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	AddAudio(audio *Audio) error
	GetAudio(id string) (*Audio, error)
	ListAudios() ([]*Audio, error)
	UpdateAudio(audio *Audio) error
	DeleteAudio(id string) error
}

type Repository interface {
	AddAudio(audio *Audio) (*Audio, error)
	GetAudio(id string) (*Audio, error)
	ListAudios() ([]*Audio, error)
	UpdateAudio(audio *Audio) (*Audio, error)
	DeleteAudio(id string) error
}
