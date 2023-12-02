package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/TrustedPay/tp-term/pkg/tpterm"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("unix", "/tmp/tp-term.sock")
	if err != nil {
		logrus.Fatalf("failed to listen on /tmp/tp-term.sock: %v", err)
	}

	opts := []grpc.ServerOption{}

	t := tpterm.NewTPTerm()

	grpcServer := grpc.NewServer(opts...)
	tpterm.RegisterTPTermServer(grpcServer, t)

	// Handle common process-killing signals so we can gracefully shut down:
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(c chan os.Signal) {
		// Wait for a SIGINT or SIGKILL:
		<-c
		println()
		logrus.Warn("Shutting down...")
		// Stop listening (and unlink the socket if unix type):
		lis.Close()
		// And we're done:
		os.Exit(0)
	}(sigc)

	logrus.Infof("TP Term listening on %s", lis.Addr())
	grpcServer.Serve(lis)
}
