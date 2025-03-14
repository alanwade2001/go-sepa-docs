package main

import (
	inf "github.com/alanwade2001/go-sepa-infra"
)

type App struct {
	Infra *inf.Infra
}

func NewApp() *App {
	app := &App{
		Infra: inf.NewInfra(),
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
