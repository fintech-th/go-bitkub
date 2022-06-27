package gobitkub

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	ERROR_MESSAGE = map[int]string{
		1:  "Invalid JSON payload",
		2:  "Missing X-BTK-APIKEY",
		3:  "Invalid API key",
		4:  "API pending for activation",
		5:  "IP not allowed",
		6:  "Missing / invalid signature",
		7:  "Missing timestamp",
		8:  "Invalid timestamp",
		9:  "Invalid user",
		10: "Invalid parameter",
		11: "Invalid symbol",
		12: "Invalid amount",
		13: "Invalid rate",
		14: "Improper rate",
		15: "Amount too low",
		16: "Failed to get balance",
		17: "Wallet is empty",
		18: "Insufficient balance",
		19: "Failed to insert order into db",
		20: "Failed to deduct balance",
		21: "Invalid order for cancellation",
		22: "Invalid side",
		23: "Failed to update order status",
		24: "Invalid order for lookup",
		25: "KYC level 1 is required to proceed",
		30: "Limit exceeds",
		40: "Pending withdrawal exists",
		41: "Invalid currency for withdrawal",
		42: "Address is not in whitelist",
		43: "Failed to deduct crypto",
		44: "Failed to create withdrawal record",
		45: "Nonce has to be numeric",
		46: "Invalid nonce",
		47: "Withdrawal limit exceeds",
		48: "Invalid bank account",
		49: "Bank limit exceeds",
		50: "Pending withdrawal exists",
		51: "Withdrawal is under maintenance",
		52: "Invalid permission",
		53: "Invalid internal address",
		54: "Address has been deprecated",
		55: "Cancel only mode",
		90: "Server error (please contact support)",
	}
)

func PlaceBid(api_key, api_secret, symbol, typ, client_id string, amount, rate float64) (*PlaceOrderResponseBitkub, error) {
	now := fmt.Sprint(time.Now().Unix())
	requestBody := map[string]string{
		"ts":        now,
		"sym":       symbol,
		"amt":       fmt.Sprintf("%g", amount),
		"rat":       fmt.Sprintf("%g", rate),
		"typ":       typ,
		"client_id": client_id,
	}
	b, err := hashRequest(api_key, api_secret, requestBody)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", `https://api.bitkub.com/api/market/place-bid`, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	addApiKeyToHeader(httpReq, api_key)
	body, err := doRequest(httpReq)
	if err != nil {
		return nil, err
	}
	var response PlaceOrderResponseBitkub
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorCode != 0 {
		return nil, newBitkubError(response.ErrorCode)
	}
	return &response, nil
}

func PlaceAskByCoin(api_key, api_secret, symbol, typ, client_id string, amount, rate float64) (*PlaceOrderResponseBitkub, error) {
	now := fmt.Sprint(time.Now().Unix())
	requestBody := map[string]string{
		"ts":        now,
		"sym":       symbol,
		"amt":       fmt.Sprintf("%g", amount),
		"rat":       fmt.Sprintf("%g", rate),
		"typ":       typ,
		"client_id": client_id,
	}
	b, err := hashRequest(api_key, api_secret, requestBody)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", `https://api.bitkub.com/api/market/place-ask`, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	addApiKeyToHeader(httpReq, api_key)
	body, err := doRequest(httpReq)
	if err != nil {
		return nil, err
	}

	var response PlaceOrderResponseBitkub
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorCode != 0 {
		return nil, newBitkubError(response.ErrorCode)
	}
	return &response, nil
}

func PlaceAskByFiat(api_key, api_secret, symbol, typ, client_id string, amount, rate float64) (*PlaceOrderResponseBitkub, error) {
	now := fmt.Sprint(time.Now().Unix())
	requestBody := map[string]string{
		"ts":        now,
		"sym":       symbol,
		"amt":       fmt.Sprintf("%g", amount),
		"rat":       fmt.Sprintf("%g", rate),
		"typ":       typ,
		"client_id": client_id,
	}
	b, err := hashRequest(api_key, api_secret, requestBody)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", `https://api.bitkub.com/api/market/place-ask-by-fiat`, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	addApiKeyToHeader(httpReq, api_key)
	body, err := doRequest(httpReq)
	if err != nil {
		return nil, err
	}
	var response PlaceOrderResponseBitkub
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorCode != 0 {
		return nil, newBitkubError(response.ErrorCode)
	}
	return &response, nil
}

func ListOrderHistory(api_key, api_secret, symbol, client_id string) (*OrderHistoryResponseBitkub, error) {
	now := fmt.Sprint(time.Now().Unix())
	requestBody := map[string]string{
		"ts":  now,
		"sym": symbol,
	}
	b, err := hashRequest(api_key, api_secret, requestBody)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", `https://api.bitkub.com/api/market/my-order-history`, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	addApiKeyToHeader(httpReq, api_key)
	body, err := doRequest(httpReq)
	if err != nil {
		return nil, err
	}
	var response OrderHistoryResponseBitkub
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorCode != 0 {
		return nil, newBitkubError(response.ErrorCode)
	}
	return &response, nil
}

func GetWallet(api_key, api_secret string) (*GetWalletResponseBitkub, error) {
	now := fmt.Sprint(time.Now().Unix())
	requestBody := map[string]string{
		"ts": now,
	}
	b, err := hashRequest(api_key, api_secret, requestBody)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", `https://api.bitkub.com/api/market/wallet`, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	addApiKeyToHeader(httpReq, api_key)
	body, err := doRequest(httpReq)
	if err != nil {
		return nil, err
	}
	var response GetWalletResponseBitkub
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorCode != 0 {
		return nil, newBitkubError(response.ErrorCode)
	}
	return &response, nil
}

func GetOpenOrders(api_key, api_secret, symbol string) (*OpenOrderResponseBitkub, error) {
	now := fmt.Sprint(time.Now().Unix())
	requestBody := map[string]string{
		"ts":  now,
		"sym": symbol,
	}
	b, err := hashRequest(api_key, api_secret, requestBody)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", `https://api.bitkub.com/api/market/my-open-orders`, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	addApiKeyToHeader(httpReq, api_key)
	body, err := doRequest(httpReq)
	if err != nil {
		return nil, err
	}
	var response OpenOrderResponseBitkub
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorCode != 0 {
		return nil, newBitkubError(response.ErrorCode)
	}
	return &response, nil
}

func CancelOrderByHash(api_key, api_secret, hash string) error {
	now := fmt.Sprint(time.Now().Unix())
	requestBody := map[string]string{
		"ts":   now,
		"hash": hash,
	}
	b, err := hashRequest(api_key, api_secret, requestBody)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest("POST", `https://api.bitkub.com/api/market/cancel-order`, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	addApiKeyToHeader(httpReq, api_key)
	body, err := doRequest(httpReq)
	if err != nil {
		return err
	}

	var response ErrorResponseBitkub
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	if response.ErrorCode != 0 {
		return newBitkubError(response.ErrorCode)
	}
	return nil
}

func hashRequest(api_key, api_secret string, requestBody map[string]string) ([]byte, error) {

	data, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	h := hmac.New(sha256.New, []byte(api_secret))

	// Write Data to it
	h.Write([]byte(string(data)))

	sha := hex.EncodeToString(h.Sum(nil))

	requestBody["sig"] = sha
	b, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func addApiKeyToHeader(httpReq *http.Request, api_key string) {
	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("X-BTK-APIKEY", api_key)
}

func doRequest(httpReq *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = errors.New(string(body))
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type RequestError struct {
	Code string
	Err  error
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func newBitkubError(errorCode int) *RequestError {
	return &RequestError{
		Code: "SBO-001",
		Err:  errors.New(ERROR_MESSAGE[errorCode]),
	}
}