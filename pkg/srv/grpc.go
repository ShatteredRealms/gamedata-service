package srv

import (
	"context"
	"errors"

	"github.com/ShatteredRealms/gamedata-service/pkg/config"
	"github.com/ShatteredRealms/gamedata-service/pkg/pb"
	"github.com/ShatteredRealms/go-common-service/pkg/util"
	"github.com/WilSimpson/gocloak/v13"
)

var (
	GamedataRoles = make([]*gocloak.Role, 0)

	RoleGameDataView = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("gamedata.view"),
		Description: gocloak.StringP("Allows getting gamedata"),
	}, &GamedataRoles)
)

var (
	CompositeGamedataRoles = make([]*gocloak.Role, 0)

	RoleManageGamedatasSelf = util.RegisterRole(&gocloak.Role{
		Name:        gocloak.StringP("gamedata.manage"),
		Description: gocloak.StringP("Allows managing all gamedata"),
		Composite:   gocloak.BoolP(true),
		Composites: &gocloak.CompositesRepresentation{
			Client: &map[string][]string{
				config.GlobalConfig.Keycloak.ClientId: {"gamedata.view"},
			},
		},
	}, &CompositeGamedataRoles)
)

var (
	ErrGamedataDoesNotExist = errors.New("CS-C-00")
	ErrGamedataCreate       = errors.New("CS-C-01")
	ErrGamedataDelete       = errors.New("CS-C-02")
	ErrGamedataGet          = errors.New("CS-C-03")
	ErrGamedataEdit         = errors.New("CS-C-04")
	ErrGamedataIdInvalid    = errors.New("CS-C-07")
)

type gamedataServiceServer struct {
	pb.UnimplementedGameDataServiceServer
	Context *GameDataContext
}

func NewGameDataServiceServer(ctx context.Context, srvCtx *GameDataContext) (pb.GameDataServiceServer, error) {
	err := srvCtx.CreateRoles(ctx, &GamedataRoles)
	if err != nil {
		return nil, err
	}
	err = srvCtx.CreateRoles(ctx, &CompositeGamedataRoles)
	if err != nil {
		return nil, err
	}

	s := &gamedataServiceServer{
		Context: srvCtx,
	}
	return s, nil
}
