package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterAITools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("ai_echo"),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			prompt, _ := req.Params.Arguments["prompt"].(string)
			return mcp.NewToolResultText("AI ответ (заглушка): " + prompt), nil
		})
}
