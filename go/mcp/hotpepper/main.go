package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found. Using system environment variables.")
	}
	s := server.NewMCPServer(
		"Hotpepper MCP Server",
		"1.0.0",
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
	s.AddTool(tool, searchToolHandlerFunc)
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func searchToolHandlerFunc(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	keyword, ok := req.Params.Arguments["keyword"].(string)
	if !ok || keyword == "" {
		return nil, errors.New("keyword is required and must be a string")
	}
	lat := fmt.Sprintf("%f", req.Params.Arguments["lat"])
	lng := fmt.Sprintf("%f", req.Params.Arguments["lng"])
	rangeStr := fmt.Sprintf("%d", int(req.Params.Arguments["range"].(float64)))
	searchReq := &SearchHotpepperRequest{
		keyword:  keyword,
		lat:      lat,
		lng:      lng,
		rangeStr: rangeStr,
	}
	resp, err := search(ctx, searchReq)
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
			shop.Urls.PC,
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

func search(_ context.Context, req *SearchHotpepperRequest) (*SearchHotpepperResponse, error) {
	client := resty.New()
	apiKey := os.Getenv("HOTPEPPER_API_KEY")
	if apiKey == "" {
		return nil, errors.New("HOTPEPPER_API_KEY is not set")
	}
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"key":     apiKey,
			"keyword": req.keyword,
			"lat":     req.lat,
			"lng":     req.lng,
			"range":   req.rangeStr,
			"format":  "json",
		}).
		SetHeader("Accept", "application/json").
		Get("http://webservice.recruit.co.jp/hotpepper/gourmet/v1/")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status())
	}
	var result SearchHotpepperResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if result.Results.ResultsAvailable == 0 {
		return nil, errors.New("no results found")
	}
	if len(result.Results.Shops) == 0 {
		return nil, errors.New("no shops found")
	}
	return &result, nil
}

type SearchHotpepperRequest struct {
	keyword  string
	lat      string
	lng      string
	rangeStr string
}

type SearchHotpepperResponse struct {
	Results struct {
		API              string `json:"api_version"`
		ResultsAvailable int    `json:"results_available"`
		ResultsReturned  string `json:"results_returned"`
		ResultsStart     int    `json:"results_start"`
		Shops            []Shop `json:"shop"`
	} `json:"results"`
}

type Shop struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	NameKana     string     `json:"name_kana"`
	Address      string     `json:"address"`
	StationName  string     `json:"station_name"`
	LogoImage    string     `json:"logo_image"`
	Catch        string     `json:"catch"`
	Access       string     `json:"access"`
	MobileAccess string     `json:"mobile_access"`
	Lat          float64    `json:"lat"`
	Lng          float64    `json:"lng"`
	Budget       Budget     `json:"budget"`
	BudgetMemo   string     `json:"budget_memo"`
	Genre        Genre      `json:"genre"`
	SubGenre     *Genre     `json:"sub_genre,omitempty"`
	Capacity     int        `json:"capacity"`
	Open         string     `json:"open"`
	Close        string     `json:"close"`
	Photo        ShopPhoto  `json:"photo"`
	Urls         URLs       `json:"urls"`
	CouponUrls   CouponURLs `json:"coupon_urls"`

	// 各種設備・オプション
	Course         string `json:"course"`
	FreeDrink      string `json:"free_drink"`
	FreeFood       string `json:"free_food"`
	PrivateRoom    string `json:"private_room"`
	Horigotatsu    string `json:"horigotatsu"`
	Tatami         string `json:"tatami"`
	Card           string `json:"card"`
	NonSmoking     string `json:"non_smoking"`
	Charter        string `json:"charter"`
	Ktai           string `json:"ktai"`
	Parking        string `json:"parking"`
	BarrierFree    string `json:"barrier_free"`
	OtherMemo      string `json:"other_memo"`
	Sommelier      string `json:"sommelier"`
	OpenAir        string `json:"open_air"`
	Show           string `json:"show"`
	Equipment      string `json:"equipment"`
	Karaoke        string `json:"karaoke"`
	Band           string `json:"band"`
	TV             string `json:"tv"`
	English        string `json:"english"`
	Pet            string `json:"pet"`
	Child          string `json:"child"`
	Lunch          string `json:"lunch"`
	Midnight       string `json:"midnight"`
	ShopDetailMemo string `json:"shop_detail_memo"`

	// エリア情報
	LargeServiceArea Area `json:"large_service_area"`
	ServiceArea      Area `json:"service_area"`
	LargeArea        Area `json:"large_area"`
	MiddleArea       Area `json:"middle_area"`
	SmallArea        Area `json:"small_area"`
}

type Area struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Genre struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Catch string `json:"catch"`
}

type Budget struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Average string `json:"average"`
}

type URLs struct {
	PC string `json:"pc"`
}

type CouponURLs struct {
	PC string `json:"pc"`
	SP string `json:"sp"`
}

type ShopPhoto struct {
	PC struct {
		L string `json:"l"`
		M string `json:"m"`
		S string `json:"s"`
	} `json:"pc"`
	Mobile struct {
		L string `json:"l"`
		S string `json:"s"`
	} `json:"mobile"`
}
