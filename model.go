package gobitkub

import "encoding/json"

type PlaceOrderResponseBitkub struct {
	ErrorCode int              `json:"error"`
	Result    PlaceOrderResult `json:"result"`
}

type PlaceOrderResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    *PlaceOrderResult `json:"data"`
}

type PlaceOrderResult struct {
	Id              int     `json:"id"`
	Hash            string  `json:"hash"`
	OrderType       string  `json:"typ"`
	Amount          float64 `json:"amt"`
	Rate            float64 `json:"rat"`
	Fee             float64 `json:"fee"`
	FeeCreditUsed   float64 `json:"cre"`
	AmountToReceive float64 `json:"rec"`
	Timestamp       int64   `json:"ts"`
}

type PlaceBidBody struct {
	InvestorID string  `json:"investor_id"`
	ClientID   string  `json:"client_id"`
	ApiKey     string  `json:"api_key"`
	ApiSecret  string  `json:"api_secret"`
	Symbol     string  `json:"symbol"`
	ThbAmount  float64 `json:"thb_amount"`
	Price      float64 `json:"price"`
	OrderType  string  `json:"order_type"`
}

type PlaceAskByCoinBody struct {
	InvestorID string  `json:"investor_id"`
	ClientID   string  `json:"client_id"`
	ApiKey     string  `json:"api_key"`
	ApiSecret  string  `json:"api_secret"`
	Symbol     string  `json:"symbol"`
	CoinAmount float64 `json:"coin_amount"`
	Price      float64 `json:"price"`
	OrderType  string  `json:"order_type"`
}

type PlaceAskByFiatBody struct {
	InvestorID string  `json:"investor_id"`
	ClientID   string  `json:"client_id"`
	ApiKey     string  `json:"api_key"`
	ApiSecret  string  `json:"api_secret"`
	Symbol     string  `json:"symbol"`
	ThbAmount  float64 `json:"thb_amount"`
	Price      float64 `json:"price"`
	OrderType  string  `json:"order_type"`
}

type OrderHistoryBody struct {
	ClientID  string `json:"client_id"`
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
	Symbol    string `json:"symbol"`
}

type OrderHistoryResponseBitkub struct {
	ErrorCode int                  `json:"error"`
	Result    []OrderHistoryResult `json:"result"`
}

type OrderHistoryResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Data    []OrderHistoryResult `json:"data"`
}

type OrderHistoryResult struct {
	TxnId         string      `json:"txn_id"`
	OrderId       int         `json:"order_id"`
	Hash          string      `json:"hash"`
	ParentOrderId int         `json:"parent_order_id"`
	SuperOrderId  int         `json:"super_order_id"`
	TakenByMe     bool        `json:"taken_by_me"`
	IsMaker       bool        `json:"is_maker"`
	Side          string      `json:"side"`
	Type          string      `json:"type"`
	Rate          JsonFloat64 `json:"rate,string"`
	Fee           JsonFloat64 `json:"fee,string"`
	Credit        JsonFloat64 `json:"credit,string"`
	Amount        JsonFloat64 `json:"amount,string"`
	Timestamp     int         `json:"ts"`
	ClientId      string      `json:"client_id"`
}

type GetWalletBody struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
}

type GetWalletResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    map[string]float64 `json:"data"`
}

type GetWalletResponseBitkub struct {
	ErrorCode int                `json:"error"`
	Result    map[string]float64 `json:"result"`
}

type OpenOrderResponseBitkub struct {
	ErrorCode int               `json:"error"`
	Result    []OpenOrderResult `json:"result"`
}

type OpenOrderResult struct {
	Id            int     `json:"id"`
	Hash          string  `json:"hash"`
	Side          string  `json:"side"`
	Type          string  `json:"type"`
	Rate          float64 `json:"rate,string"`
	Fee           float64 `json:"fee,string"`
	Credit        float64 `json:"credit,string"`
	Amount        float64 `json:"amount,string"`
	Receive       float64 `json:"receive,string"`
	ParentOrderId int     `json:"parent_order_id"`
	SuperOrderId  int     `json:"super_order_id"`
	Timestamp     int     `json:"ts"`
	ClientID      string  `json:"ci"`
}

type CancelOrdersBySymbolsBody struct {
	ApiKey    string   `json:"api_key"`
	ApiSecret string   `json:"api_secret"`
	Symbols   []string `json:"symbols"`
}

type CancelOrderResult struct {
	Side      string  `json:"side"`
	Rate      float64 `json:"rate,string"`
	Amount    float64 `json:"amount"`
	IsSuccess bool    `json:"is_success"`
	Error     string  `json:"error"`
}

type CancelOrderBySymbolResult struct {
	Symbol string              `json:"symbol"`
	Result []CancelOrderResult `json:"result"`
	Error  string              `json:"error"`
}

type CancelOrdersBySymbolsResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Data    []CancelOrderBySymbolResult `json:"data"`
}

type ErrorResponseBitkub struct {
	ErrorCode int `json:"error"`
}

type GinErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// for case that sometimes quotes numbers and sometimes doesn't
type JsonFloat64 float64

func (f JsonFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(f))
}

func (f *JsonFloat64) UnmarshalJSON(data []byte) error {
	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	var tmp float64
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	*f = JsonFloat64(tmp)
	return nil
}
