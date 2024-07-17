package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// (GET /provider/data)
func (Server) DataList(ctx *gin.Context, params DataListParams) {
	// TODO: Implement the functionality for listing provider data
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (GET /provider/data/{altorithm_id})
func (Server) EachDataGet(ctx *gin.Context, altorithmId KoyoKoyoId, params EachDataGetParams) {
	// TODO: Implement the functionality for getting data for a specific koyo ID
	ctx.JSON(http.StatusNotImplemented, nil)
}
