package tools

import (
	"context"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterEnvTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("get_env"),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			key, _ := req.Params.Arguments["name"].(string)
			val := os.Getenv(key)
			return mcp.NewToolResultText(val), nil
		})
}
