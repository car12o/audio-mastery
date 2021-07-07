package audio

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) AddAudio(audio *Audio) (*Audio, error) {
	return nil, nil
}

func (r *repository) GetAudio(id string) (*Audio, error) {
	return nil, nil
}

func (r *repository) ListAudios() ([]*Audio, error) {
	return nil, nil
}

func (r *repository) UpdateAudio(audio *Audio) (*Audio, error) {
	return nil, nil
}

func (r *repository) DeleteAudio(id string) error {
	return nil
}
