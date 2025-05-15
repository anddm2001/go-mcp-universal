package main

import (
	"github.com/anddm2001/go-mcp-universal/internal/tools"
	"log"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer("go-mcp-universal", "0.1.0")

	tools.RegisterRuntimeTools(s)
	tools.RegisterEnvTools(s)
	tools.RegisterTestTools(s)
	tools.RegisterPprofTools(s)
	tools.RegisterAITools(s)

	log.Println("MCP server started (stdio mode)...")
	server.ServeStdio(s)
}
