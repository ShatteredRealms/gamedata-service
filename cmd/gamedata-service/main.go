package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/ShatteredRealms/gamedata-service/pkg/config"
	"github.com/ShatteredRealms/gamedata-service/pkg/model/stats/level"
	"github.com/ShatteredRealms/gamedata-service/pkg/pb"
	"github.com/ShatteredRealms/gamedata-service/pkg/srv"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	commonpb "github.com/ShatteredRealms/go-common-service/pkg/pb"
	commonsrv "github.com/ShatteredRealms/go-common-service/pkg/srv"
	"github.com/ShatteredRealms/go-common-service/pkg/telemetry"
	"github.com/ShatteredRealms/go-common-service/pkg/util"
	"github.com/WilSimpson/gocloak/v13"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	defer func() {
		log.Logger.Info("Server stopped.")
	}()

	interuptCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ctx := context.Background()

	// Load configuration and setup server context
	if config.GlobalConfigErr != nil {
		log.Logger.WithContext(ctx).Errorf("loading config: %v", config.GlobalConfigErr)
		return
	}
	cfg := config.GlobalConfig

	srvCtx, err := srv.NewGamedataContext(ctx)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("creating gamedata context: %v", err)
		return
	}
	defer srvCtx.Close()

	ctx, span := srvCtx.Tracer.Start(ctx, "main")
	defer span.End()

	log.Logger.WithContext(ctx).Infof("Starting %s service", config.ServiceName)

	// OpenTelemetry setup
	otelShutdown, err := telemetry.SetupOTelSDK(ctx, "gamedata", config.Version, cfg.OpenTelemtryAddress)
	defer otelShutdown(ctx)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("connecting to otel: %v", err)
		return
	}

	// Common gRPC server setup
	keycloakClient := gocloak.NewClient(cfg.Keycloak.BaseURL)
	grpcServer, gwmux := util.InitServerDefaults(keycloakClient, cfg.Keycloak.Realm)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Health Service
	commonpb.RegisterHealthServiceServer(grpcServer, commonsrv.NewHealthServiceServer())
	err = commonpb.RegisterHealthServiceHandlerFromEndpoint(ctx, gwmux, cfg.Server.Address(), opts)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("register health service handler endpoint: %v", err)
		return
	}

	// Gamedata Service
	gamedataService, err := srv.NewGameDataServiceServer(ctx, srvCtx)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("creating gamedata service: %v", err)
		return
	}
	pb.RegisterGameDataServiceServer(grpcServer, gamedataService)
	// err = pb.RegisterGameDataServiceHandlerFromEndpoint(ctx, gwmux, cfg.Server.Address(), opts)
	// if err != nil {
	// 	log.Logger.WithContext(ctx).Errorf("register gamedata service handler endpoint: %v", err)
	// 	return
	// }

	// Setup Complete
	log.Logger.WithContext(ctx).Info("Initializtion complete")
	span.End()
	srv, srvErr := util.StartServer(ctx, grpcServer, gwmux, cfg.Server.Address())
	defer srv.Shutdown(ctx)

	_ = level.XpData

	select {
	case err := <-srvErr:
		if err != nil {
			log.Logger.Error(err)
		}

	case <-interuptCtx.Done():
		log.Logger.Info("Server canceled by user input.")
		stop()
	}
}
