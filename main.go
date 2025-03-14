package main

import (
	"github.com/alanwade2001/go-sepa-docs/internal/handler"
	"github.com/alanwade2001/go-sepa-docs/internal/repository"
	"github.com/alanwade2001/go-sepa-docs/internal/service"
	inf "github.com/alanwade2001/go-sepa-infra"
)

type App struct {
	Infra      *inf.Infra
	Repository *repository.Document
	Service    *service.Document
	Handler    *handler.Document
}

func NewApp() *App {
	infra := inf.NewInfra()
	repos := repository.NewDocument(infra.Persist)
	svc := service.NewDocument(repos)
	hdlr := handler.NewDocument(svc, infra.Router)

	app := &App{
		Infra:      infra,
		Repository: repos,
		Service:    svc,
		Handler:    hdlr,
	}

	return app
}

func (a *App) Run() {
	a.Infra.Run()
}

func main() {
	app := NewApp()
	app.Run()
}
