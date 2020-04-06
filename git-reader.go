package main

import (
	"github.com/skhatri/api-router-go/router"
	"github.com/skhatri/api-router-go/starter"
	"github.com/skhatri/git-reader/controller"
	"os"
)

func main() {
	starter.StartApp(os.Args, 6200, func(cfg router.ApiConfigurer) {
		controller.Configure(cfg)
	})
}
