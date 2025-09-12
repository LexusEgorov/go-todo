package token

type TokenRepository interface{}

type Service struct {
	storage TokenRepository
}

func New(storage TokenRepository) *Service {
	return &Service{
		storage: storage,
	}
}
