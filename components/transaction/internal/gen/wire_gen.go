// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package gen

import (
	"fmt"
	"github.com/LerianStudio/midaz/common"
	"github.com/LerianStudio/midaz/common/mmongo"
	"github.com/LerianStudio/midaz/common/mpostgres"
	"github.com/LerianStudio/midaz/common/mzap"
	"github.com/LerianStudio/midaz/components/transaction/internal/adapters/database/postgres"
	"github.com/LerianStudio/midaz/components/transaction/internal/app/command"
	"github.com/LerianStudio/midaz/components/transaction/internal/app/query"
	"github.com/LerianStudio/midaz/components/transaction/internal/domain/operation"
	"github.com/LerianStudio/midaz/components/transaction/internal/domain/transaction"
	"github.com/LerianStudio/midaz/components/transaction/internal/ports"
	"github.com/LerianStudio/midaz/components/transaction/internal/ports/http"
	"github.com/LerianStudio/midaz/components/transaction/internal/service"
	"github.com/google/wire"
	"sync"
)

// Injectors from inject.go:

// InitializeService the setup the dependencies and returns a new *service.Service instance
func InitializeService() *service.Service {
	config := service.NewConfig()
	postgresConnection := setupPostgreSQLConnection(config)
	transactionPostgreSQLRepository := postgres.NewTransactionPostgreSQLRepository(postgresConnection)
	useCase := &command.UseCase{
		TransactionRepo: transactionPostgreSQLRepository,
	}
	queryUseCase := &query.UseCase{
		TransactionRepo: transactionPostgreSQLRepository,
	}
	transactionHandler := &ports.TransactionHandler{
		Command: useCase,
		Query:   queryUseCase,
	}
	app := http.NewRouter(transactionHandler)
	logger := mzap.InitializeLogger()
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

func setupPostgreSQLConnection(cfg *service.Config) *mpostgres.PostgresConnection {
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
	}
}

func setupMongoDBConnection(cfg *service.Config) *mmongo.MongoConnection {
	connStrSource := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		cfg.MongoDBUser, cfg.MongoDBPassword, cfg.MongoDBHost, cfg.MongoDBPort)

	return &mmongo.MongoConnection{
		ConnectionStringSource: connStrSource,
	}
}

var (
	serviceSet = wire.NewSet(common.InitLocalEnvConfig, mzap.InitializeLogger, setupPostgreSQLConnection,
		setupMongoDBConnection, service.NewConfig, http.NewRouter, service.NewServer, postgres.NewTransactionPostgreSQLRepository, postgres.NewOperationPostgreSQLRepository, wire.Struct(new(ports.TransactionHandler), "*"), wire.Struct(new(command.UseCase), "*"), wire.Struct(new(query.UseCase), "*"), wire.Bind(new(transaction.Repository), new(*postgres.TransactionPostgreSQLRepository)), wire.Bind(new(operation.Repository), new(*postgres.OperationPostgreSQLRepository)),
	)

	svcSet = wire.NewSet(wire.Struct(new(service.Service), "Server", "Logger"))
)
