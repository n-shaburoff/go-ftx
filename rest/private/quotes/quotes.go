package quotes

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type RequestForQuote struct {
	FromCoin string  `json:"fromCoin"`
	ToCoin   string  `json:"toCoin"`
	Size     float64 `json:"size"`
}

type ResponseForQuote struct {
	QuoteId int `json:"quoteId"`
}

func (req *RequestForQuote) Path() string {
	return "/otc/quotes"
}

func (req *RequestForQuote) Method() string {
	return http.MethodPost
}

func (req *RequestForQuote) Query() string {
	return ""
}

func (req *RequestForQuote) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
