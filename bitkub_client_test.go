package gobitkub

import (
	"fmt"
	"testing"
)

var (
	config     = NewConfig()
	API_KEY    = config.ApiKey
	API_SECRET = config.ApiSecret
)

func TestPlaceBidMarketSuccess(t *testing.T) {
	t.Log("Testing Place Bid Success")
	res, err := PlaceBid(API_KEY, API_SECRET, "THB_ETH", "market", "FN_test", 20, 30000)
	fmt.Printf("%+v\n", res)
	var f float64
	fmt.Println(float64(res.Result.Rate) + f)
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestPlaceBidLimitSuccess(t *testing.T) {
	t.Log("Testing Place Bid Success")
	res, err := PlaceBid(API_KEY, API_SECRET, "THB_ETH", "market", "FN_test", 20, 30000)
	fmt.Printf("%+v\n", res)
	var f float64
	fmt.Println(float64(res.Result.Rate) + f)
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestPlaceAskByCoinLimitSuccess(t *testing.T) {
	t.Log("Testing Place Ask By Coin Success")
	_, err := PlaceAskByCoin(API_KEY, API_SECRET, "THB_ETH", "limit", "FN_test", 0.0002, 120000)
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestPlaceAskByCoinMarketSuccess(t *testing.T) {
	t.Log("Testing Place Ask By Coin Success")
	_, err := PlaceAskByCoin(API_KEY, API_SECRET, "THB_ETH", "market", "FN_test", 0.0002, 120000)
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestPlaceAskByFiatLimitSuccess(t *testing.T) {
	t.Log("Testing Place Ask By Fiat Success")
	_, err := PlaceAskByFiat(API_KEY, API_SECRET, "THB_ETH", "limit", "FN_test", 10, 100000)
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestPlaceAskByFiatMarketSuccess(t *testing.T) {
	t.Log("Testing Place Ask By Fiat Success")
	_, err := PlaceAskByFiat(API_KEY, API_SECRET, "THB_ETH", "market", "FN_test", 10, 100000)
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

func TestListOrderHistorySuccessWithTimestamp(t *testing.T) {
	t.Log("Testing List Order History Success With Timestamp")
	req := ListOrderHistoryRequest{
		Symbol: "THB_ETH",
		Start:  1640332716,
		End:    1653905896,
	}
	_, err := ListOrderHistory(API_KEY, API_SECRET, req)
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestListOrderHistorySuccessWithPageLimit(t *testing.T) {
	t.Log("Testing List Order History Success With Page Limit")
	req := ListOrderHistoryRequest{
		Symbol: "THB_ETH",
		Page:   1,
		Limit:  5,
	}
	_, err := ListOrderHistory(API_KEY, API_SECRET, req)
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestListOrderHistorySuccessWithoutTimestampAndPageLimit(t *testing.T) {
	t.Log("Testing List Order History Success Without Timestamp And Page Limit")
	req := ListOrderHistoryRequest{
		Symbol: "THB_ETH",
	}
	_, err := ListOrderHistory(API_KEY, API_SECRET, req)
	if err != nil {
		t.Error("error must be nil:", err)
	}
}

func TestListOrderHistoryInvalidApiSecret(t *testing.T) {
	t.Log("Testing List Order History Invalid Api Secret")
	InvalidApiSecret := "000000050104c0233e321052ca201395"
	_, err := ListOrderHistory(API_KEY, InvalidApiSecret, ListOrderHistoryRequest{})
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

func TestGetFiatDepositHistory(t *testing.T) {
	t.Log("Testing Get Fiat Depostit")
	_, err := GetFiatDepositHistory(API_KEY, API_SECRET)
	if err != nil {
		t.Error("error mustt be nil:", err)
	}
}

func TestGetFiatWithdrawHistory(t *testing.T) {
	t.Log("Testing Get Fiat Withdraw")
	_, err := GetFiatWithdrawHistory(API_KEY, API_SECRET)
	if err != nil {
		t.Error("error mustt be nil:", err)
	}
}

func TestGetCryptotDepositHistory(t *testing.T) {
	t.Log("Testing Get Crypto Depostit")
	_, err := GetCryptoDepositHistory(API_KEY, API_SECRET)
	if err != nil {
		t.Error("error mustt be nil:", err)
	}
}

func TestGetCryptotWtihdrawHistory(t *testing.T) {
	t.Log("Testing Get Crypto Withdraw")
	_, err := GetCryptoWithdrawHistory(API_KEY, API_SECRET)
	if err != nil {
		t.Error("error mustt be nil:", err)
	}
}

func TestGetBalanceSuccess(t *testing.T) {
	t.Log("Testing Place Bid Success")
	res, err := GetBalances(API_KEY, API_SECRET)
	fmt.Printf("%+v\n", res)
	if err != nil {
		t.Error("error must be nil:", err)
	}
}
