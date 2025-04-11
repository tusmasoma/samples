package resty

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/tusmasoma/samples/go/mcp/hotpepper/proto"
)

const (
	SEARCH_HOTPEPPER_API_URL = "http://webservice.recruit.co.jp/hotpepper/gourmet/v1/"
)

func Search(_ context.Context, req *proto.SearchHotpepperRequest) (*proto.SearchHotpepperResponse, error) {
	client := resty.New()
	apiKey := os.Getenv("HOTPEPPER_API_KEY")
	if apiKey == "" {
		return nil, errors.New("HOTPEPPER_API_KEY is not set")
	}
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"key":     apiKey,
			"keyword": req.GetKeyword(),
			"lat":     req.GetLat(),
			"lng":     req.GetLng(),
			"range":   req.GetRangeStr(),
			"format":  "json",
		}).
		SetHeader("Accept", "application/json").
		Get(SEARCH_HOTPEPPER_API_URL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status())
	}
	var result proto.SearchHotpepperResponse
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
