package main

import (
	"context"
	"fmt"
	"orch-go/config"
	"orch-go/internal/app"
	"orch-go/internal/services"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	//Init gRPC clients
	conn, err := grpc.NewClient(cfg.Core.Address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	serviceContainer := app.InitServices(conn)

	// Create a context that can be cancelled.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//FillDb(ctx, serviceContainer)
	app.CreateTestAccounts(ctx, serviceContainer)
	// Set up a channel to listen for OS signals
	//sigs := make(chan os.Signal, 1)
	//signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//
	//// Start the simulation in a goroutine
	//simDone := make(chan error, 1)
	//go func() {
	//	fmt.Println("Starting simulation...")
	//	simDone <- simulation.RunSimulation(ctx, serviceContainer)
	//}()
	//FillDb(ctx, serviceContainer)
	//
	//// Wait for either the simulation to finish or a signal to be received
	//select {
	//case err := <-simDone:
	//	if err != nil && !errors.Is(err, context.Canceled) {
	//		fmt.Printf("Simulation finished with error: %v\n", err)
	//	} else {
	//		fmt.Println("Simulation finished normally.")
	//	}
	//case sig := <-sigs:
	//	fmt.Printf("Received signal: %s. Shutting down...\n", sig)
	//	// Cancel the context to signal the simulation to stop
	//	cancel()
	//	// Wait for the simulation to acknowledge shutdown and save
	//	<-simDone
	//	fmt.Println("Shutdown complete.")
	//}
}

func FillDb(ctx context.Context, container *services.ServiceContainer) {
	var accTypeNames []string
	var feeTypeNames []string
	var transactionCategoryNames []string

	for i := 0; i < 10; i++ {
		accTypeNames = append(accTypeNames, fmt.Sprintf("Account Type %d", i))
		feeTypeNames = append(feeTypeNames, fmt.Sprintf("Fee Type %d", i))
		transactionCategoryNames = append(transactionCategoryNames, fmt.Sprintf("Transaction Category %d", i))
	}

	err := app.InitAll(ctx, container, accTypeNames, feeTypeNames, transactionCategoryNames)
	if err != nil {
		panic(err)
	}
}
