package srv

import (
	"context"

	"github.com/ShatteredRealms/go-common-service/pkg/auth"
	commonsrv "github.com/ShatteredRealms/go-common-service/pkg/srv"
	"github.com/WilSimpson/gocloak/v13"
)

func (c *gamedataServiceServer) validateUserPermissions(ctx context.Context, ownerId string, selfRole, otherRole *gocloak.Role) error {
	claims, ok := auth.RetrieveClaims(ctx)
	if !ok {
		return commonsrv.ErrPermissionDenied
	}
	if !claims.HasResourceRole(selfRole, c.Context.Config.Keycloak.ClientId) {
		return commonsrv.ErrPermissionDenied
	}
	if claims.Subject != ownerId && !claims.HasResourceRole(otherRole, c.Context.Config.Keycloak.ClientId) {
		return commonsrv.ErrPermissionDenied
	}
	return nil
}

func (c *gamedataServiceServer) validateRole(ctx context.Context, role *gocloak.Role) error {
	claims, ok := auth.RetrieveClaims(ctx)
	if !ok {
		return commonsrv.ErrPermissionDenied
	}
	if !claims.HasResourceRole(role, c.Context.Config.Keycloak.ClientId) {
		return commonsrv.ErrPermissionDenied
	}
	return nil
}
