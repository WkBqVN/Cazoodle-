package controller

import (
	"sync"

	"cazoodle.com/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Controller struct {
	Echo *echo.Echo
	API  *api.API
}

var c *Controller
var once sync.Once

func GetInstance() *Controller {
	once.Do(func() {
		c = &Controller{
			Echo: echo.New(),
			API:  api.GetApiInstance(),
		}
	})
	return c
}

func (c *Controller) StartOnPort(port string) error {
	err := c.Echo.Start(port)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) Init() error {
	c.Echo.Use(middleware.CORS())
	c.InitRoute()
	return nil
}
func (c *Controller) InitRoute() {
	g := c.Echo.Group("/survey")
	g.GET("/:survey_id/:form_id", c.API.GetSurvey)
	g.POST("/:survey_id/:form_id", c.API.PostSurvey)
}
