package quotes

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type RequestForCreateQuote struct {
	FromCoin string  `json:"fromCoin"`
	ToCoin   string  `json:"toCoin"`
	Size     float64 `json:"size"`
}

type ResponseForCreateQuote struct {
	QuoteId int `json:"quoteId"`
}

func (req *RequestForCreateQuote) Path() string {
	return "/otc/quotes"
}

func (req *RequestForCreateQuote) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateQuote) Query() string {
	return ""
}

func (req *RequestForCreateQuote) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
