package tools

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/tusmasoma/samples/go/mcp/hotpepper/proto"
	"github.com/tusmasoma/samples/go/mcp/hotpepper/resty"
)

func Search(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	keyword, ok := req.Params.Arguments["keyword"].(string)
	if !ok || keyword == "" {
		return nil, errors.New("keyword is required and must be a string")
	}
	lat := fmt.Sprintf("%f", req.Params.Arguments["lat"])
	lng := fmt.Sprintf("%f", req.Params.Arguments["lng"])
	rangeStr := fmt.Sprintf("%d", int(req.Params.Arguments["range"].(float64)))
	searchReq := &proto.SearchHotpepperRequest{
		Keyword:  keyword,
		Lat:      lat,
		Lng:      lng,
		RangeStr: rangeStr,
	}
	resp, err := resty.Search(ctx, searchReq)
	if err != nil {
		return nil, err
	}
	var contents []mcp.Content
	for _, shop := range resp.Results.Shops {
		var subGenreName string
		if shop.SubGenre != nil {
			subGenreName = shop.SubGenre.Name
		}
		text := fmt.Sprintf(
			`📍 %s
			%s
			📌 住所: %s
			🚉 最寄駅: %s
			💰 予算: %s
			🍽 ジャンル: %s / %s
			🕒 営業時間: %s ～ %s
			🪑 席数: %d
			🌐 URL: %s`,
			shop.Name,
			shop.Catch,
			shop.Address,
			shop.StationName,
			shop.Budget.Name,
			shop.Genre.Name,
			subGenreName,
			shop.Open,
			shop.Close,
			shop.Capacity,
			shop.Urls.Pc,
		)
		content := mcp.TextContent{
			Type: "text",
			Text: text,
		}
		contents = append(contents, content)
	}
	return &mcp.CallToolResult{
		Content: contents,
		IsError: false,
	}, nil
}
