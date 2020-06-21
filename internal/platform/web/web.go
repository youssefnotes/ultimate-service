package web

import (
	"github.com/go-chi/chi"
	"log"
)

type App struct {
	mux *chi.Mux
	log *log.Logger
}

func NewApp(log *log.Logger) *App {
	return &App{mux: chi.NewRouter(), log: log}
}
