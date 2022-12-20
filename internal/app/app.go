package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"service/internal"
	"service/internal/controller"
	"service/internal/service"
)

type App struct {
	port    string
	router  *chi.Mux
	service service.ServiceInterface
}

func NewApp(conf internal.Config) *App {
	return &App{
		port:    conf.Port,
		router:  chi.NewRouter(),
		service: service.NewService(conf.Db),
	}
}

func (a *App) Run() {
	a.registerRoutes()
	url := fmt.Sprintf("localhost:%s", a.port)
	err := http.ListenAndServe(url, a.router)
	if err != nil {
		fmt.Printf("can't start http service: %s\n", err.Error())
	}
}

func (a *App) registerRoutes() {
	a.router.Use(middleware.Logger)

	a.router.Post("/create", controller.Create(a.service))
	a.router.Post("/makeFriends", controller.MakeFriends(a.service))
	a.router.Get("/getAll", controller.GetAll(a.service))
	a.router.Get("/friends/{id}", controller.GetFriends(a.service))
	a.router.Delete("/user", controller.DeleteUser(a.service))
	a.router.Put("/user/{id}", controller.UpdateAge(a.service))
}
