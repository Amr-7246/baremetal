// The main file responsble for Loading configuration, initializing the logger, starting the server, and listening for "kill" signals (Graceful Shutdown). 
// It shouldn't know how the TCP connection works, it just knows how to start the process

package main

import (
	"os"
	"log/slog"
	"tcp-server/server"
)

func main(){
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	port := os.Getenv("SERVER_PORT")
	if (port == "") { port = "8080" }

	srv := server.New(":"+port, logger)
	logger.Info("Starting application...")
	if err := srv.Start(); err != nil {
		logger.Error("Application failed", "error", err)
		os.Exit(1) 
	}
}