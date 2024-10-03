package api

import "github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"

type BeLifelineServerImpl struct{}

func NewBeLifelineServiceHandler() mainv1connect.BeLifelineServiceHandler {
	return &BeLifelineServerImpl{}
}
