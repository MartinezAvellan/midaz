//go:build wireinject
// +build wireinject

package gen

import (
	"fmt"
	"github.com/LerianStudio/midaz/common/mgrpc"
	"github.com/LerianStudio/midaz/components/transaction/internal/adapters/grpc"
	"sync"

	"github.com/LerianStudio/midaz/common"
	"github.com/LerianStudio/midaz/common/mmongo"
	"github.com/LerianStudio/midaz/common/mpostgres"
	"github.com/LerianStudio/midaz/common/mzap"
	"github.com/LerianStudio/midaz/components/transaction/internal/adapters/database/postgres"
	"github.com/LerianStudio/midaz/components/transaction/internal/app/command"
	"github.com/LerianStudio/midaz/components/transaction/internal/app/query"
	o "github.com/LerianStudio/midaz/components/transaction/internal/domain/operation"
	t "github.com/LerianStudio/midaz/components/transaction/internal/domain/transaction"
	"github.com/LerianStudio/midaz/components/transaction/internal/ports"
	httpHandler "github.com/LerianStudio/midaz/components/transaction/internal/ports/http"
	"github.com/LerianStudio/midaz/components/transaction/internal/service"
	"github.com/google/wire"
)

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

func setupGRPCConnection(cfg *service.Config) *mgrpc.GRPCConnection {
	addr := fmt.Sprintf("%s:%s", cfg.LedgerGRPCAddr, cfg.LedgerGRPCPort)

	return &mgrpc.GRPCConnection{
		Addr: addr,
	}
}

var (
	serviceSet = wire.NewSet(
		common.InitLocalEnvConfig,
		mzap.InitializeLogger,
		setupPostgreSQLConnection,
		setupMongoDBConnection,
		setupGRPCConnection,
		service.NewConfig,
		httpHandler.NewRouter,
		service.NewServer,
		postgres.NewTransactionPostgreSQLRepository,
		postgres.NewOperationPostgreSQLRepository,
		grpc.NewAccountGRPC,
		wire.Struct(new(ports.TransactionHandler), "*"),
		wire.Struct(new(command.UseCase), "*"),
		wire.Struct(new(query.UseCase), "*"),
		wire.Bind(new(t.Repository), new(*postgres.TransactionPostgreSQLRepository)),
		wire.Bind(new(o.Repository), new(*postgres.OperationPostgreSQLRepository)),
	)

	svcSet = wire.NewSet(
		wire.Struct(new(service.Service), "Server", "Logger"),
	)
)

// InitializeService the setup the dependencies and returns a new *service.Service instance
func InitializeService() *service.Service {
	wire.Build(serviceSet, svcSet)

	return nil
}
