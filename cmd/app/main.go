package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	genGrpcTest "github.com/MShilenko/gen-test/gen/grpc-server/test"
	genGrpcTest2 "github.com/MShilenko/gen-test/gen/grpc-server/test2"
	genHttpTest "github.com/MShilenko/gen-test/gen/http-server/test"
	gen2HttpTest "github.com/MShilenko/gen-test/gen/http-server/test2"
	testGrpcTransport "github.com/MShilenko/gen-test/internal/transport/grpc/test"
	test2GrpcTransport "github.com/MShilenko/gen-test/internal/transport/grpc/test2"
	testHttpTransport "github.com/MShilenko/gen-test/internal/transport/http/test"
	test2HttpTransport "github.com/MShilenko/gen-test/internal/transport/http/test2"
)

func main() {
	fmt.Println("App started")

	testGrpcTransport := testGrpcTransport.New()
	test2GrpcTransport := test2GrpcTransport.New()
	testHttpTransport := testHttpTransport.New()
	test2HttpTransport := test2HttpTransport.New()

	// gRPC
	lis, err := net.Listen("tcp", "localhost:8087")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	go grpcServer.Serve(lis)

	genGrpcTest.RegisterTestServiceServer(grpcServer, testGrpcTransport)
	genGrpcTest2.RegisterTest2ServiceServer(grpcServer, test2GrpcTransport)

	// http
	r := mux.NewRouter()
	genHttpTest.HandlerFromMux(genHttpTest.NewStrictHandler(testHttpTransport, nil), r)
	gen2HttpTest.HandlerFromMux(gen2HttpTest.NewStrictHandler(test2HttpTransport, nil), r)

	httpServer := &http.Server{
		Handler: r,
		Addr:    "localhost:8082",
	}
	go func() {
		if err = httpServer.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println("Start http server error: ", err)
			return
		}
	}()

	// shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)

	<-sigCh

	fmt.Println("App stopped")
}
