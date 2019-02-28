package Service

import (
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
	"github.com/umirode/prom-calendar-russia/src/Domain/Service/DTO"
)

type IAuthService interface {
	Login(authDTO *DTO.AuthDTO) (*Entity.User, error)
	Signup(authDTO *DTO.AuthDTO) (*Entity.User, error)
}
