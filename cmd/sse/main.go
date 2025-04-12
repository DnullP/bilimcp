package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	"github.io/dnullp/bilimcp/internel"
)

func main() {
	s := server.NewMCPServer(
		"BiliBili MCP 服务器",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	internel.RegisterServer(s)

	// Start the server
	sse := server.NewSSEServer(s, server.WithBaseURL("http://localhost:8080"))
	if err := sse.Start(":8080"); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
