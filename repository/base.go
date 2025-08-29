package repository

import (
	"context"
	"crypto/tls"
	"time"

	"apis_service/domain"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	timeoutInit = 10 * time.Second
	timeoutPing = 2 * time.Second
)

func GetMongoDB(cfg *domain.MongoConfig, secondary bool) *mongo.Database {
	ctx, cancelCtx := context.WithTimeout(context.Background(), timeoutInit)
	defer cancelCtx()

	ctxPing, cancelCtxPing := context.WithTimeout(context.Background(), timeoutPing)
	defer cancelCtxPing()

	var cliOpts *options.ClientOptions
	if cfg.URI != "" {
		cliOpts = options.Client().ApplyURI(cfg.URI).SetTLSConfig(&tls.Config{}) // TLS enabled
	} else {

		cliOpts = options.Client().SetHosts(cfg.Hosts)

		if cfg.Tls {
			cliOpts = cliOpts.SetTLSConfig(&tls.Config{}) // TLS enabled
		}

		if cfg.ReplicaSetName == "" {
			cliOpts = cliOpts.SetDirect(true)
		}
		// If username/password are provided, set credentials
		if cfg.Username != "" && cfg.Password != "" {
			cred := options.Credential{
				Username:   cfg.Username,
				Password:   cfg.Password,
				AuthSource: "admin",
			}
			cliOpts = cliOpts.SetAuth(cred)
		}

		var maxPoolSize uint64 = 200

		var minPoolSize uint64 = 100

		cliOpts.SetMaxPoolSize(maxPoolSize)
		cliOpts.SetMinPoolSize(minPoolSize)

	}

	cli, err := mongo.Connect(ctx, cliOpts)
	if err != nil {
		panic(err)
	}

	err = cli.Ping(ctxPing, readpref.Primary())
	if err != nil {
		panic(err)
	}

	dbOpts := options.Database()

	if secondary {
		secondary := readpref.SecondaryPreferred()
		dbOpts = dbOpts.SetReadPreference(secondary)
	}

	return cli.Database(cfg.DBName, dbOpts)
}
