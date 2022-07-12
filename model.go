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

type ListOrderHistoryRequest struct {
	Sig       string `json:"sig,omitempty"`
	Timestamp string `json:"ts"`
	Symbol    string `json:"sym"`
	Start     int    `json:"start,omitempty"`
	End       int    `json:"end,omitempty"`
	Page      int    `json:"p,omitempty"`
	Limit     int    `json:"lmt,omitempty"`
}

type OrderHistoryResponseBitkub struct {
	ErrorCode   int                  `json:"error"`
	Result      []OrderHistoryResult `json:"result"`
	Pagignation Pagination           `json:"pagination"`
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
	Rate          JsonFloat64 `json:"rate"`
	Fee           JsonFloat64 `json:"fee"`
	Credit        JsonFloat64 `json:"credit"`
	Amount        JsonFloat64 `json:"amount"`
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

type GetSymbolsResponseBitkub struct {
	ErrorCode int      `json:"error"`
	Result    []Symbol `json:"result"`
}

type Symbol struct {
	Id     int    `json:"id"`
	Symbol string `json:"symbol"`
	Info   string `json:"info"`
}

type GetTickersResponseBitkub struct {
	Result map[string]Ticker `json:"result"`
}

type Ticker struct {
	Id            int     `json:"id"`
	Last          float64 `json:"last"`
	LowestAsk     float64 `json:"lowestAsk"`
	HighestBid    float64 `json:"highestBid"`
	PercentChange float64 `json:"percentChange"`
	BaseVolume    float64 `json:"baseVolume"`
	QuoteVolume   float64 `json:"quoteVolume"`
	IsFrozen      int     `json:"isFrozen"`
	High24hr      float64 `json:"high24hr"`
	Low24hr       float64 `json:"low24hr"`
}

type GetFiatDepositHistoryResponseBitkub struct {
	ErrorCode int                  `json:"error"`
	Result    []FiatDepositHistory `json:"result"`
}

type FiatDepositHistory struct {
	Id       string  `json:"txn_id"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
	Time     int     `json:"time"`
}

type GetFiatWithdrawHistoryResponseBitkub struct {
	ErrorCode int                   `json:"error"`
	Result    []FiatWithdrawHistory `json:"result"`
}

type FiatWithdrawHistory struct {
	Id       string  `json:"txn_id"`
	Currency string  `json:"curruncy"`
	Amount   float64 `json:"amount"`
	Fee      float64 `json:"fee"`
	Status   string  `json:"status"`
	Time     int     `json:"time"`
}

type GetCryptoDepositHistoryResponseBitkub struct {
	ErrorCode int                    `json:"error"`
	Result    []CryptoDepositHistory `json:"result"`
}

type CryptoDepositHistory struct {
	Hash          string  `json:"hash"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	FromAddress   string  `json:"from_address"`
	ToAddress     string  `json:"to_address"`
	Confirmations int     `json:"confirmations"`
	Status        string  `json:"status"`
	Time          int     `json:"time"`
}

type GetCryptoWithdrawHistoryResponseBitkub struct {
	ErrorCode int                     `json:"error"`
	Result    []CryptoWithdrawHistory `json:"result"`
}

type CryptoWithdrawHistory struct {
	Id       string      `json:"txn_id"`
	Hash     string      `json:"hash"`
	Currency string      `json:"currency"`
	Amount   JsonFloat64 `json:"amount"`
	Fee      float64     `json:"fee"`
	Address  string      `json:"address"`
	Stauts   string      `json:"status"`
	Time     int         `json:"time"`
}

type Pagination struct {
	Page int `json:"page"`
	Last int `json:"last"`
	Next int `json:"next"`
	Prev int `json:"prev"`
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
