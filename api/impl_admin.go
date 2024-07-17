package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// (GET /admin/client)
func (Server) ClientList(ctx *gin.Context, params ClientListParams) {
	// TODO: Implement the functionality for listing clients
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (POST /admin/client)
func (Server) ClientCreate(ctx *gin.Context) {
	// TODO: Implement the functionality for creating a client
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (POST /admin/client/{client_id})
func (Server) ClientDelete(ctx *gin.Context, clientId ProviderClientId) {
	// TODO: Implement the functionality for deleting a client
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (POST /admin/client/{client_id}/revoke)
func (Server) ClientRevoke(ctx *gin.Context, clientId ProviderClientId) {
	// TODO: Implement the functionality for revoking a client
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (POST /admin/extinfo)
func (Server) ExtInfoCreate(ctx *gin.Context) {
	// TODO: Implement the functionality for creating extinfo
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (DELETE /admin/extinfo/{extinfo_id})
func (Server) ExtInfoDelete(ctx *gin.Context, extinfoId ExtInfoExtInfoId) {
	// TODO: Implement the functionality for deleting extinfo
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (POST /admin/extinfo/{extinfo_id}/revoke)
func (Server) ExtInfoRevoke(ctx *gin.Context, extinfoId ExtInfoExtInfoId) {
	// TODO: Implement the functionality for revoking extinfo
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (POST /admin/koyo)
func (Server) KoyoCreate(ctx *gin.Context) {
	// TODO: Implement the functionality for creating koyo information
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (DELETE /admin/koyo/{koyo_id})
func (Server) KoyoDelete(ctx *gin.Context, koyoId KoyoKoyoId) {
	// TODO: Implement the functionality for deleting koyo information
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (POST /admin/koyo/{koyo_id}/revoke)
func (Server) KoyoRevoke(ctx *gin.Context, koyoId KoyoKoyoId) {
	// TODO: Implement the functionality for revoking koyo api key
	ctx.JSON(http.StatusNotImplemented, nil)
}
