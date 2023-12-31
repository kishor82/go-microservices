package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/kishor82/go-microservices/currency/data"
	"github.com/kishor82/go-microservices/currency/protos/currency"
	"github.com/kishor82/go-microservices/currency/server"
)

func main() {
	log := hclog.Default()

	rates, err := data.NewRates(log)
	if err != nil {
		log.Error("Unable to generate rates", err)
		os.Exit(1)
	}

	gs := grpc.NewServer()

	cp := server.NewCurrency(rates, log)

	currency.RegisterCurrencyServer(gs, cp)

	// gRPC Server Reflection provides information about publicly-accessible gRPC services on a server,
	// and assists clients at runtime to construct RPC requests and responses without precompiled service information.
	// It is used by gRPC CLI, which can be used to introspect server protos and send/receive test RPCs.

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to Listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
