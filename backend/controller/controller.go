package controller

import (
	"sync"

	"cazoodle.com/routes"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	*echo.Echo
}

var c *Controller
var once sync.Once

func GetInstance() *Controller {
	once.Do(func() {
		c = &Controller{
			echo.New()}
	})
	return c
}

func (c *Controller) StartOnPort(port string) error {
	err := c.Start(port)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) Init() {
}

func (c *Controller) InitRoute(routeData routes.RouteData) {
	for _, route := range routeData.Routes {
		if route.Group != "" {
			c.SetGroupRoute(route)
		} else {
			c.SetRoute(route)
		}
	}
}

func (c *Controller) SetGroupRoute(route routes.Route) {

}
func (c *Controller) SetRoute(route routes.Route) {
}
