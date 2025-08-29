package api

import (
	"apis_service/domain"
	"apis_service/usecases"
	"net"
	"os"

	"github.com/becash/apis/gen_go/swallow_channel_to_service"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Server is Structure implementation wrapper
// all dependency injection goes here.
type Server struct {
	swallow_channel_to_service.UnimplementedServiceToSwallowServer
	log      *zap.SugaredLogger
	useCases *usecases.UseCases
}

func ListenAndServe(
	log *zap.SugaredLogger,
	addrGrpc string,
	stop <-chan os.Signal,
	useCases *usecases.UseCases,
	cfg *domain.Config,
	listener net.Listener,
) {
	grpcServer := grpc.NewServer()

	serviceServer := &Server{
		log:      log,
		useCases: useCases,
	}

	swallow_channel_to_service.RegisterServiceToSwallowServer(grpcServer, serviceServer)

	go func() {
		<-stop
		log.Infow("api: attempting graceful shutdown")
		grpcServer.GracefulStop()
		log.Info("api: clean shutdown")
	}()

	log.Infof("api available on %s", addrGrpc)

	err := grpcServer.Serve(listener)
	if err != nil {
		log.Panic(errors.Wrap(err, "ListenAndServe"))
	}

}
