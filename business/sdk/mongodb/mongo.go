package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Host        string
	Port        string
	Username    string
	Password    string
	DBName      string
	AuthDB      string
	AllowDirect bool
}

func InitializeMongo(cfg Config) (*mongo.Database, error) {
	dsn := createDSNFromConfig(cfg)
	clientOptions := options.Client().ApplyURI(dsn)
	clientOptions.SetDirect(true)
	clientOptions.SetAuth(options.Credential{AuthSource: cfg.AuthDB, Username: cfg.Username, Password: cfg.Password})
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err

	}
	// if err := client.Connect(context.TODO()); err != nil {
	// 	log.Println("error connecting to mongoDB client reason:", err)
	// 	log.Fatal(err)
	// }
	db := client.Database(cfg.DBName)
	log.Println("connected to mongoDB!!!")
	return db, nil
}
func CheckStatus(db *mongo.Database) bool {
	if err := db.Client().Ping(context.TODO(), nil); err != nil {
		return false
	}
	return true
}

// createDSNFromConfig creates a db DSN from given config ex. "mongodb://usrname:password@host:port"
func createDSNFromConfig(conf Config) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBName)
}
