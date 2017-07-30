package app

import (
	"github.com/yundream/contextlogging/handler"
	"net/http"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (a *Application) Run(addr string) {
	h := handler.New()
	h.Regist()
	err := http.ListenAndServe(addr, h.Router())
	if err != nil {
		panic(err)
	}
}
