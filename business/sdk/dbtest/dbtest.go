package dbtest

import (
	"fmt"

	"github.com/bentenison/microservice/business/sdk/mongodb"
	"github.com/bentenison/microservice/business/sdk/redisdb"
	"github.com/bentenison/microservice/business/sdk/sqldb"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	DB        DS
	Log       *logger.CustomLogger
	BusDomain BusDomain
}
type DS struct {
	SQL *sqlx.DB
	MGO *mongo.Database
	Rdb *redis.Client
}

func New() (*Database, error) {
	db, err := sqldb.Open(sqldb.Config{
		User:         "epic",
		Password:     "admin#123",
		Host:         "localhost",
		Name:         "epic",
		MaxIdleConns: 10,
		MaxOpenConns: 10,
	})
	if err != nil {
		return nil, fmt.Errorf("connecting to db: %w", err)
	}

	defer db.Close()
	rdb, err := redisdb.OpenRDB(redisdb.Config{})
	if err != nil {
		return nil, err
	}
	// starting mongio db connection

	mongo, err := mongodb.InitializeMongo(mongodb.Config{
		Username:    "admin",
		Password:    "admin#123",
		AuthDB:      "admin",
		Host:        "localhost",
		Port:        "27017",
		DBName:      "EXECUTOR",
		AllowDirect: false,
	})
	if err != nil {
		return nil, err
	}
	ds := DS{
		MGO: mongo,
		SQL: db,
		Rdb: rdb,
	}
	log := logger.NewCustomLogger(map[string]interface{}{
		"service": "Test",
	})
	return &Database{
		DB:        ds,
		Log:       log,
		BusDomain: newBusDomain(log, ds),
	}, nil
}
