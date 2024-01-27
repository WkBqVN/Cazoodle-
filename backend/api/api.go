package api

import (
	"sync"

	"cazoodle.com/service"
)

type API struct {
	FormSerivce *service.FormService
}

var a *API
var once sync.Once

func GetApiInstance() *API {
	once.Do(func() {
		a = &API{
			FormSerivce: service.GetFormServiceInstance(),
		}
	})
	return a
}