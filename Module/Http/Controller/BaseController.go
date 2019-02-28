package Controller

import (
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/umirode/prom-calendar-russia/Module/Http/Error"
	"github.com/umirode/prom-calendar-russia/Module/Http/Response"
	"github.com/umirode/prom-calendar-russia/src/Domain/Model/Entity"
	"github.com/umirode/prom-calendar-russia/src/Domain/Service"
)

type BaseController struct {
	UserService Service.IUserService
}

func (c *BaseController) GetToken(context echo.Context) (*jwt.Token, error) {
	token, ok := context.Get("user").(*jwt.Token)
	if !ok {
		return nil, Error.NewInvalidTokenError()
	}

	return token, nil
}

func (c *BaseController) GetTokenClaims(context echo.Context) (jwt.MapClaims, error) {
	token, err := c.GetToken(context)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, Error.NewInvalidTokenError()
	}

	return claims, nil
}

func (c *BaseController) GetParam(context echo.Context, key string, valueType string) (interface{}, error) {
	param := context.Param(key)
	if param == "" {
		return nil, Error.NewRequestParsingError()
	}

	switch valueType {
	case "int":
		result, _ := strconv.Atoi(param)

		return result, nil
	case "uint":
		result, _ := strconv.Atoi(param)

		return uint(result), nil
	case "string":
		return param, nil
	}

	return nil, Error.NewRequestParsingError()
}

func (c *BaseController) GetCurrentUser(context echo.Context) (*Entity.User, error) {
	claims, err := c.GetTokenClaims(context)
	if err != nil {
		return nil, err
	}
	userID := uint(claims["user_id"].(float64))

	return c.UserService.GetOneById(userID)
}

func (c *BaseController) Response(context echo.Context, status int, data interface{}, message string) error {
	return context.JSON(status, Response.NewResponse(status, data, message))
}
