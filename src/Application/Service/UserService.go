package Service

import (
	"github.com/umirode/prom-calendar-russia/src/Domain/Error"
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
	"github.com/umirode/prom-calendar-russia/src/Domain/Repository"
)

type UserService struct {
	userRepository Repository.IUserRepository
}

func NewUserService(userRepository Repository.IUserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetOneById(id uint) (*Entity.User, error) {
	user, err := s.userRepository.FindOneByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, Error.NewNotFoundError()
	}

	return user, nil
}
