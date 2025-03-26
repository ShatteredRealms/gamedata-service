package srv_test

import (
	"io"
	"testing"

	"github.com/ShatteredRealms/gamedata-service/pkg/srv"
	"github.com/ShatteredRealms/go-common-service/pkg/auth"
	"github.com/ShatteredRealms/go-common-service/pkg/config"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	commonsrv "github.com/ShatteredRealms/go-common-service/pkg/srv"
	mock_gocloak "github.com/WilSimpson/gocloak/v13/mocks"
	iauth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.opentelemetry.io/otel"
	"go.uber.org/mock/gomock"
)

var (
	srvCtx             *srv.GameDataContext
	mockConrol         *gomock.Controller
	keycloakClientMock *mock_gocloak.MockGoCloakIface
	authFunc           iauth.AuthFunc
)

const (
	kcServiceName = "gamedata"
)

func TestSrv(t *testing.T) {
	BeforeEach(func() {
		log.Logger.Out = io.Discard
		mockConrol = gomock.NewController(GinkgoT())
		keycloakClientMock = mock_gocloak.NewMockGoCloakIface(mockConrol)
		srvCtx = &srv.GameDataContext{
			Context: &commonsrv.Context{
				Config: &config.BaseConfig{
					Keycloak: config.KeycloakConfig{
						ClientId: kcServiceName,
					},
				},
				KeycloakClient: keycloakClientMock,
				Tracer:         otel.Tracer("TestGamedataContext"),
			},
		}
		authFunc = auth.AuthFunc(keycloakClientMock, "realm")
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "Srv Suite")
}
