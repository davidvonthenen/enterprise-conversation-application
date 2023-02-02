// Copyright 2022 Symbl.ai SDK contributors. All Rights Reserved.
// SPDX-License-Identifier: MIT

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"

	server "github.com/dvonthenen/enterprise-reference-implementation/cmd/example-middleware-analyzer/server"
	analyzer "github.com/dvonthenen/enterprise-reference-implementation/pkg/middleware-analyzer"
)

func main() {
	// os hooks
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	// init
	analyzer.Init(analyzer.EnterpriseInit{
		LogLevel: analyzer.LogLevelStandard, // LogLevelStandard / LogLevelFull / LogLevelTrace
	})

	middlewareServer, err := server.New(server.ServerOptions{
		CrtFile:   "localhost.crt",
		KeyFile:   "localhost.key",
		RabbitURI: "amqp://guest:guest@localhost:5672",
	})
	if err != nil {
		fmt.Printf("server.New failed. Err: %v\n", err)
		os.Exit(1)
	}

	// init
	err = middlewareServer.Init()
	if err != nil {
		fmt.Printf("middlewareServer.Init() failed. Err: %v\n", err)
		os.Exit(1)
	}

	// start
	fmt.Printf("Starting server...\n")
	err = middlewareServer.Start()
	if err != nil {
		fmt.Printf("middlewareServer.Start() failed. Err: %v\n", err)
		os.Exit(1)
	}

	fmt.Print("Press ENTER to exit!\n\n")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	err = middlewareServer.Stop()
	if err != nil {
		fmt.Printf("middlewareServer.Stop() failed. Err: %v\n", err)
	}

	fmt.Printf("Server stopped...\n\n")
}