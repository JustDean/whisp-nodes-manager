package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/JustDean/whisp-nodes-manager/internal/gateway_api"
	"github.com/JustDean/whisp-nodes-manager/internal/node_api"
	"github.com/JustDean/whisp-nodes-manager/internal/nodes_manager"
	"github.com/JustDean/whisp-nodes-manager/pkg"
)

func main() {
	manager := nodes_manager.SetupNodesManager(nodes_manager.Config{
		Secret: "aboba",
	})
	nodeApiServer, err := node_api.SetServer(
		pkg.GrpcServerConfig{
			Host: pkg.GetEnv("NODE_API_HOST", "localhost"),
			Port: pkg.GetEnv("NODE_API_PORT", "7777"),
		},
		manager,
	)
	if err != nil {
		log.Fatalf("Error setting node api server: %v", err)
	}
	gatewayApiServer, err := gateway_api.SetServer(
		pkg.GrpcServerConfig{
			Host: pkg.GetEnv("NODE_API_HOST", "localhost"),
			Port: pkg.GetEnv("NODE_API_PORT", "7000"),
		},
		manager,
	)
	if err != nil {
		log.Fatalf("Error setting node api server: %v", err)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	services := []interface{ Run(context.Context) }{manager, nodeApiServer, gatewayApiServer}
	var wg sync.WaitGroup
	for _, s := range services {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Run(ctx)
		}()
	}
	wg.Wait()
	log.Println("Shutdown complete!")
}
