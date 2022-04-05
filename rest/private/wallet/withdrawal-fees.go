package wallet

import "net/http"

type RequestForWithdrawalFees struct {
	Coin    string  `json:"coin"`
	Size    float64 `json:"size"`
	Address string  `json:"address"`
	// Optionals
	Tag     string `json:"tag,omitempty"`
	Network string `json:"method,omitempty"`
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
	return ""
}

func (req *RequestForWithdrawalFees) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
