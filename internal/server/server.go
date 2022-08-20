package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/odpf/kay/config"
	"github.com/odpf/kay/internal/store/postgres"
	"github.com/odpf/salt/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

func Start(config *config.Config) error {
	logger := log.NewLogrus(log.LogrusWithLevel(config.Log))

	grpcServer := grpc.NewServer()

	// init http proxy
	// timeoutGrpcDialCtx, grpcDialCancel := context.WithTimeout(context.Background(), time.Second*5)
	// defer grpcDialCancel()

	// gwmux := runtime.NewServeMux(
	// 	runtime.WithErrorHandler(runtime.DefaultHTTPErrorHandler),
	// 	runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
	// 		MarshalOptions: protojson.MarshalOptions{
	// 			UseProtoNames: true,
	// 		},
	// 		UnmarshalOptions: protojson.UnmarshalOptions{
	// 			DiscardUnknown: true,
	// 		},
	// 	}),
	// )

	address := fmt.Sprintf(":%d", config.App.Port)
	// grpcConn, err := grpc.DialContext(
	// 	timeoutGrpcDialCtx,
	// 	address,
	// 	grpc.WithInsecure(),
	// 	grpc.WithDefaultCallOptions(),
	// )
	// if err != nil {
	// 	return err
	// }

	// runtimeCtx, runtimeCancel := context.WithCancel(context.Background())
	// defer runtimeCancel()

	baseMux := http.NewServeMux()
	baseMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	server := &http.Server{
		Handler:      grpcHandlerFunc(grpcServer, baseMux),
		Addr:         address,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	logger.Info(fmt.Sprintf("server running on port: %d", config.App.Port))
	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			return err
		}
	}

	db, err := postgres.NewClient(config.DB)
	if err != nil {
		return err
	}
	defer func() {
		fmt.Println("closing db connection...")
		db.Close()
	}()
	return nil
}

// grpcHandlerFunc routes http1 calls to baseMux and http2 with grpc header to grpcServer.
// Using a single port for proxying both http1 & 2 protocols will degrade http performance
// but for our usecase the convenience per performance tradeoff is better suited
// if in future, this does become a bottleneck(which I highly doubt), we can break the service
// into two ports, default port for grpc and default+1 for grpc-gateway proxy.
// We can also use something like a connection multiplexer
// https://github.com/soheilhy/cmux to achieve the same.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
