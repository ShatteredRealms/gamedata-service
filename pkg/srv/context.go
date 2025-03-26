package srv

import (
	"context"

	"github.com/ShatteredRealms/gamedata-service/pkg/config"
	commonsrv "github.com/ShatteredRealms/go-common-service/pkg/srv"
)

type GameDataContext struct {
	*commonsrv.Context
}

func NewGamedataContext(ctx context.Context) (*GameDataContext, error) {
	gamedataCtx := &GameDataContext{
		Context: commonsrv.NewContext(&config.GlobalConfig.BaseConfig, config.ServiceName),
	}
	ctx, span := gamedataCtx.Tracer.Start(ctx, "context.gamedata.new")
	defer span.End()

	// pg, err := commonrepo.ConnectDB(ctx, cconfig.DBPoolConfig{Master: config.GlobalConfig.Postgres}, config.GlobalConfig.Redis)
	// if err != nil {
	// 	return nil, fmt.Errorf("connect db: %w", err)
	// }
	//
	// migrater, err := commonrepo.NewPgxMigrater(ctx, config.GlobalConfig.Postgres.PostgresDSN(), config.GlobalConfig.MigrationPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("postgres migrater: %w", err)
	// }

	return gamedataCtx, nil
}

func (c *GameDataContext) Close() {
}
