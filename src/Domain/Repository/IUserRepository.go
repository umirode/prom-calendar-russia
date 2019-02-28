package Repository

import (
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
)

type IUserRepository interface {
	Save(user *Entity.User) error

	FindOneByID(id uint) (*Entity.User, error)
	FindOneByEmail(email string) (*Entity.User, error)
	FindOneByEmailAndPassword(email string, password string) (*Entity.User, error)
}
