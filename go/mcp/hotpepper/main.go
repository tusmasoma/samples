package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/tusmasoma/samples/go/mcp/hotpepper/tools"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found. Using system environment variables.")
	}
	s := server.NewMCPServer(
		"Hotpepper MCP Server",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)
	tool := mcp.NewTool("search_salon",
		mcp.WithDescription("指定された条件でサロンを検索します。"),
		mcp.WithString("keyword",
			mcp.Required(),
			mcp.Description("検索キーワード"),
		),
		mcp.WithNumber("lat",
			mcp.Description("緯度 (例: 35.6895)"),
		),
		mcp.WithNumber("lng",
			mcp.Description("経度 (例: 139.6917)"),
		),
		mcp.WithNumber("range",
			mcp.Description("検索範囲（1:300m, 2:500m, 3:1000m, 4:2000m, 5:3000m）"),
		),
	)
	s.AddTool(tool, tools.Search)
	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
	}
}
