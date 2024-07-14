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

type HogeGet200JSONResponseContentType = string

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

// (GET /provider/clients)
func (Server) ClientsGetClient(ctx *gin.Context) {
	// TODO: Implement the functionality for getting provider clients
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (GET /provider/data)
func (Server) DataList(ctx *gin.Context, params DataListParams) {
	// TODO: Implement the functionality for listing provider data
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (GET /provider/data/{altorithm_id})
func (Server) EachDataGet(ctx *gin.Context, altorithmId AlgorithmAlgorithmId, params EachDataGetParams) {
	// TODO: Implement the functionality for getting data for a specific algorithm ID
	ctx.JSON(http.StatusNotImplemented, nil)
}

// (GET /status)
func (Server) StatusGet(ctx *gin.Context) {
	// TODO: Implement the functionality for getting status
	ctx.Writer.WriteHeader(http.StatusOK)
}
