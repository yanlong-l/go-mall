package library

import (
	"context"
	"encoding/json"
	"github.com/yanlong-l/go-mall/common/logger"
	"github.com/yanlong-l/go-mall/common/util/httptool"
)

// 对接 ipwhois.io 的Lib
// Documentation: https://ipwhois.io/documentation

type WhoisLib struct {
	ctx context.Context
}

func NewWhoisLib(ctx context.Context) *WhoisLib {
	return &WhoisLib{ctx: ctx}
}

type WhoisIpDetail struct {
	Ip            string  `json:"ip"`
	Success       bool    `json:"success"`
	Type          string  `json:"type"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continent_code"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"country_code"`
	Region        string  `json:"region"`
	RegionCode    string  `json:"region_code"`
	City          string  `json:"city"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	IsEu          bool    `json:"is_eu"`
	Postal        string  `json:"postal"`
	CallingCode   string  `json:"calling_code"`
	Capital       string  `json:"capital"`
	Borders       string  `json:"borders"`
}

func (whois *WhoisLib) GetHostIpDetail() (*WhoisIpDetail, error) {
	httpStatusCode, respBody, err := httptool.Get(
		whois.ctx, "https://ipwho.is",
		httptool.WithHeaders(map[string]string{
			"User-Agent": "curl/7.77.0",
		}),
	)
	if err != nil {
		logger.Error(context.Background(), "whois request error", "err", err, "httpStatusCode", httpStatusCode)
		return nil, err
	}
	reply := new(WhoisIpDetail)
	json.Unmarshal(respBody, reply)

	return reply, nil
}
