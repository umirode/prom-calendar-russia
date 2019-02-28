package Service

import (
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
)

type IUserService interface {
	GetOneById(id uint) (*Entity.User, error)
}
