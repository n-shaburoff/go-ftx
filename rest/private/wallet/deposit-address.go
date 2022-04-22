package wallet

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

type RequestForDepositAddress struct {
	Coin    string `json:"-"`
	Network string `url:"method,omitempty"`
}

type ResponseForDepositAddress struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
}

func (req *RequestForDepositAddress) Path() string {
	return fmt.Sprintf("/wallet/deposit_address/%s", req.Coin)
}

func (req *RequestForDepositAddress) Method() string {
	return http.MethodGet
}

func (req *RequestForDepositAddress) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForDepositAddress) Payload() []byte {
	return nil
}
