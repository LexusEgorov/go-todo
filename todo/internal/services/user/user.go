package user

import "github.com/LexusEgorov/todo/internal/models/dto"

type UserRepository interface {
	Create() error
	Read() error
	Update() error
	Delete() error
}

type Service struct {
	storage UserRepository
}

// Register implements user.UserService.
func (s Service) Register(data dto.Register) (dto.Tokens, error) {
	//Add to db
	//Send to auth service
	//Recieve tokens
	//Send tokens
	panic("unimplemented")
}

// Auth implements user.UserService.
func (s Service) Auth(data dto.Auth) (dto.Tokens, error) {
	//Send to auth service
	//Recieve tokens
	//Send tokens
	panic("unimplemented")
}

// Delete implements user.UserService.
func (s Service) Delete(uID int) error {
	//Just delete user
	panic("unimplemented")
}

// Get implements user.UserService.
func (s Service) Get(uID int) (dto.User, error) {
	//Just get user
	panic("unimplemented")
}

// Update implements user.UserService.
func (s Service) Update(user dto.UserUpdate) (dto.User, error) {
	//Just update user
	panic("unimplemented")
}

func New() *Service {
	return &Service{}
}
