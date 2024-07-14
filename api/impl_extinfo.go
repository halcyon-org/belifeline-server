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

// (GET /extinfo/HOGE_ID/data)
func (Server) HogeGet(ctx *gin.Context, params HogeGetParams) {
	// TODO: Implement the functionality for getting HOGE data
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (GET /extinfo/{extinfo_id})
func (Server) EachExtInfoGet(ctx *gin.Context, extinfoId ExtInfoExtInfoId) {
	// TODO: Implement the functionality for getting extinfo by ID
	ctx.JSON(http.StatusNotImplemented, nil)
}
