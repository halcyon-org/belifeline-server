package helper

import (
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"
	"github.com/halcyon-org/kizuna/gen/belifeline/v1/mainv1connect"
	"github.com/halcyon-org/kizuna/internal/di"
)

func NewClients(t *testing.T) ([]mainv1connect.AdminServiceClient, []mainv1connect.ProviderServiceClient, []mainv1connect.ExternalInformationServiceClient, []mainv1connect.KoyoServiceClient) {
	t.Helper()

	server := NewServer(t)

	adminServiceClients := newAdminServiceClient(t, server)
	providerServiceClients := newProviderServiceClient(t, server)
	externalInformationServiceClients := newExternalInformationServiceClient(t, server)
	koyoServiceClients := newKoyoServiceClient(t, server)

	return adminServiceClients, providerServiceClients, externalInformationServiceClients, koyoServiceClients
}

func NewServer(t *testing.T) *httptest.Server {
	t.Helper()

	controllers, err := di.InitializeControllerSet()
	if err != nil {
		t.Fatal(err)
	}

	server, err := controllers.BeLifelineController.TestServe()
	if err != nil {
		t.Fatal(err)
	}

	return server
}

func newAdminServiceClient(t *testing.T, server *httptest.Server) []mainv1connect.AdminServiceClient {
	t.Helper()

	connectClient := mainv1connect.NewAdminServiceClient(server.Client(), server.URL)
	grpcClient := mainv1connect.NewAdminServiceClient(server.Client(), server.URL, connect.WithGRPC())

	return []mainv1connect.AdminServiceClient{connectClient, grpcClient}
}

func newProviderServiceClient(t *testing.T, server *httptest.Server) []mainv1connect.ProviderServiceClient {
	t.Helper()

	connectClient := mainv1connect.NewProviderServiceClient(server.Client(), server.URL)
	grpcClient := mainv1connect.NewProviderServiceClient(server.Client(), server.URL, connect.WithGRPC())

	return []mainv1connect.ProviderServiceClient{connectClient, grpcClient}
}

func newExternalInformationServiceClient(t *testing.T, server *httptest.Server) []mainv1connect.ExternalInformationServiceClient {
	t.Helper()

	connectClient := mainv1connect.NewExternalInformationServiceClient(server.Client(), server.URL)
	grpcClient := mainv1connect.NewExternalInformationServiceClient(server.Client(), server.URL, connect.WithGRPC())

	return []mainv1connect.ExternalInformationServiceClient{connectClient, grpcClient}
}

func newKoyoServiceClient(t *testing.T, server *httptest.Server) []mainv1connect.KoyoServiceClient {
	t.Helper()

	connectClient := mainv1connect.NewKoyoServiceClient(server.Client(), server.URL)
	grpcClient := mainv1connect.NewKoyoServiceClient(server.Client(), server.URL, connect.WithGRPC())

	return []mainv1connect.KoyoServiceClient{connectClient, grpcClient}
}
