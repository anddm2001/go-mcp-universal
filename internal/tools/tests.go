package tools

import (
	"bytes"
	"context"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterTestTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("run_tests"),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			cmd := exec.Command("go", "test", "./...", "-v", "-short")
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &out
			if err := cmd.Run(); err != nil {
				return mcp.NewToolResultText("Test errors:\n" + out.String()), nil
			}
			return mcp.NewToolResultText("✅ Tests passed:\n" + out.String()), nil
		})
}
