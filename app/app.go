package app

import (
	"fmt"
	"net/http"

	"github.com/arisnotargon/t_write_test/service"
)

type App struct {
	route   *Route
	service *service.Service
}

func NewApp() *App {
	a := &App{}
	a.service = service.NewService()
	a.route = NewRoute(a)

	return a
}

func (a *App) Run() {
	a.route.InitRoute()
	fmt.Println("Listening")
	http.ListenAndServe(":8081", nil)
}
