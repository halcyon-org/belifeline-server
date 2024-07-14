package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// (POST /admin/algorithm)
func (Server) AlgorithmCreate(ctx *gin.Context) {
	// TODO: Implement the functionality for creating an algorithm
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (DELETE /admin/algorithm/{algorithm_id})
func (Server) AlgorithmDelete(ctx *gin.Context, algorithmId AlgorithmAlgorithmId) {
	// TODO: Implement the functionality for deleting an algorithm by ID
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (POST /admin/client)
func (Server) ClientCreate(ctx *gin.Context) {
	// TODO: Implement the functionality for creating a client
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (POST /admin/client/{client_id})
func (Server) ClientDelete(ctx *gin.Context, clientId ProviderClientId) {
	// TODO: Implement the functionality for deleting a client by ID
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (POST /admin/extinfo)
func (Server) ExtInfoCreate(ctx *gin.Context) {
	// TODO: Implement the functionality for creating extinfo
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (DELETE /admin/extinfo/{extinfo_id})
func (Server) ExtInfoDelete(ctx *gin.Context, extinfoId ExtInfoExtInfoId) {
	// TODO: Implement the functionality for deleting extinfo by ID
	ctx.JSON(http.StatusNotImplemented, nil)
}
