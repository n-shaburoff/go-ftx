package wallet

import (
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type RequestForWithdrawHistories struct {
	Start int64 `url:"start_time,omitempty"`
	End   int64 `url:"end_time,omitempty"`
}

type ResponseForWithdrawHistories []Withdraw

type Withdraw struct {
	Coin    string `json:"coin"`
	Address string `json:"address"`
	Tag     string `json:"tag"`
	Status  string `json:"status"`
	Txid    string `json:"txid"`

	Fee  float64 `json:"fee"`
	Size float64 `json:"size"`

	Time time.Time `json:"time"`

	ID int `json:"id"`

	Method string `json:"method"`
	Notes  string `json:"notes"`
}

func (req *RequestForWithdrawHistories) Path() string {
	return "/wallet/withdrawals"
}

func (req *RequestForWithdrawHistories) Method() string {
	return http.MethodGet
}

func (req *RequestForWithdrawHistories) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestForWithdrawHistories) Payload() []byte {
	return nil
}
