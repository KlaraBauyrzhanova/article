package article

import (
	"github.com/labstack/echo/v4"
)

type Command interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
}
