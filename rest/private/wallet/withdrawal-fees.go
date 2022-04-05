package wallet

import (
	"github.com/google/go-querystring/query"
	"net/http"
)

type RequestForWithdrawalFees struct {
	Coin    string  `url:"coin"`
	Size    float64 `url:"size"`
	Address string  `url:"address"`
	// Optionals
	Tag     string `url:"tag,omitempty"`
	Network string `url:"method,omitempty"`
}

type ResponseForWithdrawalFees struct {
	Method    string  `json:"method"`
	Address   string  `json:"address"`
	Fee       float64 `json:"fee"`
	Congested bool    `json:"congested"`
}

func (req *RequestForWithdrawalFees) Path() string {
	return "/wallet/withdrawal_fee"
}

func (req *RequestForWithdrawalFees) Method() string {
	return http.MethodGet
}

func (req *RequestForWithdrawalFees) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForWithdrawalFees) Payload() []byte {
	return nil
}
