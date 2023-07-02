package endpoint

import (
	"app1/internal/app1/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	s *service.Service
}

func New() *Endpoint {
	return &Endpoint{}
}

func (e *Endpoint) Action(c echo.Context) error {
	ans := e.s.TimeUntil()
	err := c.String(http.StatusOK, ans)
	if err != nil {
		return err
	}
	return err
}
