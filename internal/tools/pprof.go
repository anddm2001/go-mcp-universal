package tools

import (
	"bytes"
	"context"
	"runtime/pprof"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterPprofTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("pprof_goroutine"),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var buf bytes.Buffer
			if err := pprof.Lookup("goroutine").WriteTo(&buf, 2); err != nil {
				return nil, err
			}
			return mcp.NewToolResultText(buf.String()), nil
		})

	s.AddTool(mcp.NewTool("pprof_heap"),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var buf bytes.Buffer
			if err := pprof.Lookup("heap").WriteTo(&buf, 2); err != nil {
				return nil, err
			}
			return mcp.NewToolResultText(buf.String()), nil
		})
}
