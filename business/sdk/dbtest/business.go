package dbtest

import (
	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/executorbus"
	"github.com/bentenison/microservice/business/domain/executorbus/stores/executordb"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/docker/docker/client"
)

type BusDomain struct {
	Delegate *delegate.Delegate
	ExecBus  *executorbus.Business
}

func newBusDomain(log *logger.CustomLogger, db DS) BusDomain {
	delegate := delegate.New(log)
	ds := mux.DataSource{
		MGO: db.MGO,
		SQL: db.SQL,
		RDB: db.Rdb,
	}
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	execbus := executorbus.NewBusiness(log, delegate, executordb.NewStore(log, ds), cli)
	return BusDomain{
		Delegate: delegate,
		ExecBus:  execbus,
	}
}
