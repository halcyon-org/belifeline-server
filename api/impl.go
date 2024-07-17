package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// optional code omitted

type Server struct{}

func NewServer() Server {
	return Server{}
}

type ExampleInfoGet200JSONResponseContentType = string

// (GET /status)
func (Server) StatusGet(ctx *gin.Context) {
	// TODO: Implement the functionality for getting status
	ctx.Writer.WriteHeader(http.StatusOK)
}
