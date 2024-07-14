package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// (GET /algorithm)
func (Server) AlgorithmList(ctx *gin.Context, params AlgorithmListParams) {
	// TODO: Implement the functionality for listing algorithms
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (GET /algorithm/{algorithm_id})
func (Server) EachAlgorithmGet(ctx *gin.Context, algorithmId AlgorithmAlgorithmId) {
	// TODO: Implement the functionality for getting an algorithm by ID
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (PUT /algorithm/{algorithm_id})
func (Server) EachAlgorithmUpdate(ctx *gin.Context, algorithmId AlgorithmAlgorithmId) {
	// TODO: Implement the functionality for updating an algorithm by ID
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (GET /algorithm/{algorithm_id}/data)
func (Server) EachAlgorithmDataGet(ctx *gin.Context, algorithmId AlgorithmAlgorithmId, params EachAlgorithmDataGetParams) {
	// TODO: Implement the functionality for getting algorithm data by ID
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (PUT /algorithm/{algorithm_id}/data)
func (Server) EachAlgorithmDataUpdate(ctx *gin.Context, algorithmId AlgorithmAlgorithmId, params EachAlgorithmDataUpdateParams) {
	// TODO: Implement the functionality for updating algorithm data by ID
	ctx.JSON(http.StatusNotImplemented, nil)
}
