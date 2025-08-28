package main

import (
	"apis_service/domain"
	"net"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	locGrpc "apis_service/grpc"
	"apis_service/repository"
	"apis_service/usecases"
)

func main() {
	cfg := domain.NewConfig("")
	log := cfg.Log

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	stopGrpc := make(chan os.Signal, 1)
	stopMainLoop := make(chan os.Signal, 1)

	dbConnection := repository.GetMongoDB(cfg.Mongo, false)
	repoAutoIncrement := repository.NewRepoAutoIncrement(
		dbConnection,
		log.Named("repoAutoIncrement"),
		"auto_increments",
	)
	repoProducts := repository.NewRepoProducts(dbConnection, log.Named("repoProducts"))

	useCases := usecases.NewUseCases(
		cfg,
		log.Named("usecases"),
		repoAutoIncrement,
		repoProducts,
	)

	grpcCon, err := net.Listen("tcp", "0.0.0.0:"+cfg.GrpcPort)
	if err != nil {
		log.Panic("grpc", err)
	}

	go locGrpc.ListenAndServe(
		log.Named("grpc"),
		"0.0.0.0:"+cfg.GrpcPort,
		stopGrpc,
		useCases,
		cfg,
		grpcCon,
	)

	sig := <-sigCh
	log.Infof("got signal %v shutting down services", sig)
	stopGrpc <- sig
	stopMainLoop <- sig
	//repoCatDisp.Producer.Close()
	log.Info("all down, bye!")
}
