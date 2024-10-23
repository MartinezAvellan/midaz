// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package gen

import (
	"fmt"
	"github.com/LerianStudio/midaz/common"
	"github.com/LerianStudio/midaz/common/mcasdoor"
	"github.com/LerianStudio/midaz/common/mgrpc"
	"github.com/LerianStudio/midaz/common/mlog"
	"github.com/LerianStudio/midaz/common/mmongo"
	"github.com/LerianStudio/midaz/common/mpostgres"
	"github.com/LerianStudio/midaz/common/mzap"
	"github.com/LerianStudio/midaz/components/transaction/internal/adapters/database/mongodb"
	"github.com/LerianStudio/midaz/components/transaction/internal/adapters/database/postgres"
	"github.com/LerianStudio/midaz/components/transaction/internal/adapters/grpc"
	"github.com/LerianStudio/midaz/components/transaction/internal/app/command"
	"github.com/LerianStudio/midaz/components/transaction/internal/app/query"
	"github.com/LerianStudio/midaz/components/transaction/internal/domain/account"
	"github.com/LerianStudio/midaz/components/transaction/internal/domain/assetrate"
	"github.com/LerianStudio/midaz/components/transaction/internal/domain/metadata"
	"github.com/LerianStudio/midaz/components/transaction/internal/domain/operation"
	"github.com/LerianStudio/midaz/components/transaction/internal/domain/transaction"
	"github.com/LerianStudio/midaz/components/transaction/internal/ports/http"
	"github.com/LerianStudio/midaz/components/transaction/internal/service"
	"github.com/google/wire"
	"sync"
)

// Injectors from inject.go:

// InitializeService the setup the dependencies and returns a new *service.Service instance
func InitializeService() *service.Service {
	config := service.NewConfig()
	logger := mzap.InitializeLogger()
	casdoorConnection := setupCasdoorConnection(config, logger)
	postgresConnection := setupPostgreSQLConnection(config, logger)
	transactionPostgreSQLRepository := postgres.NewTransactionPostgreSQLRepository(postgresConnection)
	grpcConnection := setupGRPCConnection(config, logger)
	accountGRPCRepository := grpc.NewAccountGRPC(grpcConnection)
	operationPostgreSQLRepository := postgres.NewOperationPostgreSQLRepository(postgresConnection)
	assetRatePostgreSQLRepository := postgres.NewAssetRatePostgreSQLRepository(postgresConnection)
	mongoConnection := setupMongoDBConnection(config, logger)
	metadataMongoDBRepository := mongodb.NewMetadataMongoDBRepository(mongoConnection)
	useCase := &command.UseCase{
		TransactionRepo: transactionPostgreSQLRepository,
		AccountGRPCRepo: accountGRPCRepository,
		OperationRepo:   operationPostgreSQLRepository,
		AssetRateRepo:   assetRatePostgreSQLRepository,
		MetadataRepo:    metadataMongoDBRepository,
	}
	queryUseCase := &query.UseCase{
		TransactionRepo: transactionPostgreSQLRepository,
		AccountGRPCRepo: accountGRPCRepository,
		OperationRepo:   operationPostgreSQLRepository,
		MetadataRepo:    metadataMongoDBRepository,
	}
	transactionHandler := &http.TransactionHandler{
		Command: useCase,
		Query:   queryUseCase,
	}
	operationHandler := &http.OperationHandler{
		Command: useCase,
		Query:   queryUseCase,
	}
	assetRateHandler := &http.AssetRateHandler{
		Command: useCase,
		Query:   queryUseCase,
	}
	app := http.NewRouter(logger, casdoorConnection, transactionHandler, operationHandler, assetRateHandler)
	server := service.NewServer(config, app, logger)
	serviceService := &service.Service{
		Server: server,
		Logger: logger,
	}
	return serviceService
}

// inject.go:

var onceConfig sync.Once

const prdEnvName = "production"

func setupPostgreSQLConnection(cfg *service.Config, log mlog.Logger) *mpostgres.PostgresConnection {
	connStrPrimary := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PrimaryDBHost, cfg.PrimaryDBUser, cfg.PrimaryDBPassword, cfg.PrimaryDBName, cfg.PrimaryDBPort)

	connStrReplica := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.ReplicaDBHost, cfg.ReplicaDBUser, cfg.ReplicaDBPassword, cfg.ReplicaDBName, cfg.ReplicaDBPort)

	return &mpostgres.PostgresConnection{
		ConnectionStringPrimary: connStrPrimary,
		ConnectionStringReplica: connStrReplica,
		PrimaryDBName:           cfg.PrimaryDBName,
		ReplicaDBName:           cfg.ReplicaDBName,
		Component:               "transaction",
		Logger:                  log,
	}
}

func setupMongoDBConnection(cfg *service.Config, log mlog.Logger) *mmongo.MongoConnection {
	connStrSource := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		cfg.MongoDBUser, cfg.MongoDBPassword, cfg.MongoDBHost, cfg.MongoDBPort)

	return &mmongo.MongoConnection{
		ConnectionStringSource: connStrSource,
		Database:               cfg.MongoDBName,
		Logger:                 log,
	}
}

func setupCasdoorConnection(cfg *service.Config, log mlog.Logger) *mcasdoor.CasdoorConnection {
	casdoor := &mcasdoor.CasdoorConnection{
		JWKUri:           cfg.JWKAddress,
		Endpoint:         cfg.CasdoorAddress,
		ClientID:         cfg.CasdoorClientID,
		ClientSecret:     cfg.CasdoorClientSecret,
		OrganizationName: cfg.CasdoorOrganizationName,
		ApplicationName:  cfg.CasdoorApplicationName,
		EnforcerName:     cfg.CasdoorEnforcerName,
		Logger:           log,
	}

	return casdoor
}

func setupGRPCConnection(cfg *service.Config, log mlog.Logger) *mgrpc.GRPCConnection {
	addr := fmt.Sprintf("%s:%s", cfg.LedgerGRPCAddr, cfg.LedgerGRPCPort)

	return &mgrpc.GRPCConnection{
		Addr:   addr,
		Logger: log,
	}
}

var (
	serviceSet = wire.NewSet(common.InitLocalEnvConfig, mzap.InitializeLogger, setupPostgreSQLConnection,
		setupMongoDBConnection,
		setupCasdoorConnection,
		setupGRPCConnection, service.NewConfig, http.NewRouter, service.NewServer, postgres.NewTransactionPostgreSQLRepository, postgres.NewOperationPostgreSQLRepository, postgres.NewAssetRatePostgreSQLRepository, mongodb.NewMetadataMongoDBRepository, grpc.NewAccountGRPC, wire.Struct(new(http.TransactionHandler), "*"), wire.Struct(new(http.OperationHandler), "*"), wire.Struct(new(http.AssetRateHandler), "*"), wire.Struct(new(command.UseCase), "*"), wire.Struct(new(query.UseCase), "*"), wire.Bind(new(transaction.Repository), new(*postgres.TransactionPostgreSQLRepository)), wire.Bind(new(operation.Repository), new(*postgres.OperationPostgreSQLRepository)), wire.Bind(new(assetrate.Repository), new(*postgres.AssetRatePostgreSQLRepository)), wire.Bind(new(account.Repository), new(*grpc.AccountGRPCRepository)), wire.Bind(new(metadata.Repository), new(*mongodb.MetadataMongoDBRepository)),
	)

	svcSet = wire.NewSet(wire.Struct(new(service.Service), "Server", "Logger"))
)
