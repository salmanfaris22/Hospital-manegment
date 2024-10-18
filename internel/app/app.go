package app

import "main.go/internel/router"

type App interface {
	Start()
}

type impl struct {
	r router.Router
}

func (i *impl) Start() {
	i.r.Start() // Starts the router and server
}

func NewApp(rout router.Router) App {
	return &impl{
		r: rout,
	}
}
