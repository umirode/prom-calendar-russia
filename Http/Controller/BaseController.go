package Controller

import (
	"github.com/labstack/echo"
	"github.com/umirode/prom-calendar-russia/Http/Error"
	"strconv"
)

type BaseController struct {}

func (c *BaseController) getParam(context echo.Context, key string, valueType string) (interface{}, error) {
	param := context.Param(key)
	if param == "" {
		return nil, Error.NewRequestParsingError()
	}

	switch valueType {
	case "int":
		result, _ := strconv.Atoi(param)

		return result, nil
		break
	case "uint":
		result, _ := strconv.Atoi(param)

		return uint(result), nil
		break
	case "string":
		return param, nil
		break
	}

	return nil, Error.NewRequestParsingError()
}