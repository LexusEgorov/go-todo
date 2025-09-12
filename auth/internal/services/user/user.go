package user

type UserRepository interface{}

type Service struct {
	storage UserRepository
}

func New(storage UserRepository) *Service {
	return &Service{
		storage: storage,
	}
}
