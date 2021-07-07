package grammarian

type Service interface {
	Process(record []byte) ([]byte, error)
}
