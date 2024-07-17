package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// (GET /extinfo)
func (Server) ExtInfoList(ctx *gin.Context, params ExtInfoListParams) {
	// TODO: Implement the functionality for listing extinfo
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (GET /extinfo/{extinfo_id})
func (Server) ExtInfoGet(ctx *gin.Context, extinfoId ExtInfoExtInfoId) {
	// TODO: Implement the functionality for getting extinfo by ID
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (GET /extinfo/example_id/data)
func (Server) ExampleInfoGet(ctx *gin.Context, params ExampleInfoGetParams) {
	// TODO: Implement the functionality for getting exmample data
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (POST /extinfo/example_id/data)
func (Server) ExampleInfoPost(ctx *gin.Context) {
	// TODO: Implement the functionality for creating exmample data
	ctx.JSON(http.StatusNotImplemented, nil)
}
