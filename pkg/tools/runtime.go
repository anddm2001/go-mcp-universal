package tools

import (
	"context"
	"fmt"
	"runtime"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterRuntimeTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("dump_goroutines"),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			buf := make([]byte, 1<<20)
			n := runtime.Stack(buf, true)
			return mcp.NewToolResultText(string(buf[:n])), nil
		})

	s.AddTool(mcp.NewTool("num_goroutines"),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			n := runtime.NumGoroutine()
			return mcp.NewToolResultText(fmt.Sprintf("%d goroutines", n)), nil
		})
}
