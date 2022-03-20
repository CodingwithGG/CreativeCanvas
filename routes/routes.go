package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

//InitServer is used to initialize server routes

type routes struct {
	router *gin.Engine
}

func (r routes) Run(addr ...string) error {
	err := r.router.Run(addr...)
	return errors.Wrap(err, "Unable to run server")
}

func NewRoutes() routes {
	r := routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("/v1")
	r.user(v1)

	return r
}
