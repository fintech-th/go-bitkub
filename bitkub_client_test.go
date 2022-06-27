package gobitkub

import (
	"testing"
)

var (
	config     = NewConfig()
	API_KEY    = config.ApiKey
	API_SECRET = config.ApiSecret
)

func TestPlaceBidSuccess(t *testing.T) {
	t.Log("Testing Place Bid Success")
	_, err := PlaceBid(API_KEY, API_SECRET, "THB_ETH", "limit", "FN_test", 10, 30000)
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestPlaceAskByCoinSuccess(t *testing.T) {
	t.Log("Testing Place Ask By Coin Success")
	_, err := PlaceAskByCoin(API_KEY, API_SECRET, "THB_ETH", "limit", "FN_test", 0.0002, 120000)
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestPlaceAskByFiatSuccess(t *testing.T) {
	t.Log("Testing Place Ask By Fiat Success")
	_, err := PlaceAskByFiat(API_KEY, API_SECRET, "THB_ETH", "limit", "FN_test", 5, 100000)
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestPlaceBidInvalidApiSecret(t *testing.T) {
	InvalidApiSecret := "000000050104c0233e321052ca201395"
	t.Log("Testing Place Bid Invalid Api Secret")
	_, err := PlaceBid(API_KEY, InvalidApiSecret, "THB_ETH", "limit", "FN_test", 10, 30000)
	if err.Error() != "Missing / invalid signature" {
		t.Error("error must be \"Missing / invalid signature\", but get:", err)
	}
}

func TestPlaceAskByCoinInvalidAmount(t *testing.T) {
	t.Log("Testing Place Ask By Coin Invalid Amount")
	_, err := PlaceAskByCoin(API_KEY, API_SECRET, "THB_ETH", "limit", "FN_test", 0, 10000000)
	if err.Error() != "Invalid amount" {
		t.Error("error must be \"Invalid amount\", but get:", err)
	}
}

func TestPlaceAskByFiatInsufficientBalance(t *testing.T) {
	t.Log("Testing Place Ask By Fiat Insufficient Balance")
	_, err := PlaceAskByFiat(API_KEY, API_SECRET, "THB_ETH", "limit", "FN_test", 10000, 100000)
	if err.Error() != "Insufficient balance" {
		t.Error("error must be \"Insufficient balance\", but get:", err)
	}
}

func TestListOrderHistorySuccess(t *testing.T) {
	t.Log("Testing List Order History Success")
	_, err := ListOrderHistory(API_KEY, API_SECRET, "THB_ETH")
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestListOrderHistoryInvalidApiSecret(t *testing.T) {
	t.Log("Testing List Order History Invalid Api Secret")
	InvalidApiSecret := "000000050104c0233e321052ca201395"
	_, err := ListOrderHistory(API_KEY, InvalidApiSecret, "THB_ETH")
	if err.Error() != "Missing / invalid signature" {
		t.Error("error must be \"Missing / invalid signature\", but get:", err)
	}
}

func TestGetWalletSuccess(t *testing.T) {
	t.Log("Testing Get Wallet Success")
	_, err := GetWallet(API_KEY, API_SECRET)
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestGetWalletInvalidApiSecret(t *testing.T) {
	t.Log("Testing Get Wallet Invalid Api Secret")
	InvalidApiSecret := "000000050104c0233e321052ca201395"
	_, err := GetWallet(API_KEY, InvalidApiSecret)
	if err.Error() != "Missing / invalid signature" {
		t.Error("error must be \"Missing / invalid signature\", but get:", err)
	}
}

func TestGetOpenOrdersSuccess(t *testing.T) {
	t.Log("Testing Get Open Orders Success")
	_, err := GetOpenOrders(API_KEY, API_SECRET, "THB_ETH")
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestGetOpenInvalidApiSecret(t *testing.T) {
	t.Log("Testing Get Open Order Invalid Api Secret")
	InvalidApiSecret := "000000050104c0233e321052ca201395"
	_, err := GetOpenOrders(API_KEY, InvalidApiSecret, "THB_ETH")
	if err.Error() != "Missing / invalid signature" {
		t.Error("error must be \"Missing / invalid signature\", but get:", err)
	}
}

func TestGetOpenInvalidSymbol(t *testing.T) {
	t.Log("Testing Get Open Order Invalid Symbol")
	_, err := GetOpenOrders(API_KEY, API_SECRET, "THB_ETHHHH")
	if err.Error() != "Invalid symbol" {
		t.Error("error must be \"Invalid symbol\", but get:", err)
	}
}

func TestCancelOrderByHashSuccess(t *testing.T) {
	t.Log("Testing Cancel Order By Hash Success")
	res, err := PlaceAskByCoin(API_KEY, API_SECRET, "THB_ETH", "limit", "FN_test", 0.0002, 120000)
	if err != nil {
		t.Error("error must be nil:", err)
	}
	err = CancelOrderByHash(API_KEY, API_SECRET, res.Result.Hash)
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestCancelOrderByHashInvalidHash(t *testing.T) {
	t.Log("Testing Cancel Order By Hash Success")
	hash := "fwQ6dng9Uf7uE552jYZUZC13n6f"
	err := CancelOrderByHash(API_KEY, API_SECRET, hash)
	if err.Error() != "Invalid order for cancellation" {
		t.Error("error must be \"Invalid order for cancellation\", but get:", err)
	}
}

func TestGetSymbols(t *testing.T) {
	t.Log("Testing Place Bid Success")
	res, err := GetSymbols()
	if err != nil {
		t.Error("error must be nil:", err)
	}
	if len(res.Result) == 0 {
		t.Error("length of result must not be 0")
	}
}

func TestGetTickers(t *testing.T) {
	t.Log("Testing Place Bid Success")
	res, err := GetTickers()
	if err != nil {
		t.Error("error must be nil:", err)
	}
	if len(res.Result) == 0 {
		t.Error("length of result must not be 0")
	}
}
