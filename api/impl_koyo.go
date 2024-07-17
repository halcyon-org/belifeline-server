package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// (GET /koyo)
func (Server) KoyoList(ctx *gin.Context, params KoyoListParams) {
	// TODO: Implement the functionality for listing koyos
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (GET /koyo/{koyo_id})
func (Server) EachKoyoGet(ctx *gin.Context, koyoId KoyoKoyoId) {
	// TODO: Implement the functionality for getting an koyo by ID
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (PUT /koyo/{koyo_id})
func (Server) EachKoyoUpdate(ctx *gin.Context, koyoId KoyoKoyoId) {
	// TODO: Implement the functionality for updating an koyo by ID
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (POST /koyo/{koyo_id}/data)
func (Server) EachKoyoDataNew(ctx *gin.Context, koyoId KoyoKoyoId) {
	// TODO: Implement the functionality for updating koyo data by ID
	ctx.JSON(http.StatusNotImplemented, nil)
}
