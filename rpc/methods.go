package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func IsWalletLoaded() (bool, error) {
	return false, nil
}

func LoadWallet(password string) (string, error) {
	body, err := json.Marshal(map[string]interface{}{
		"method": "loadwallet",
		"params": []string{password},
	})
	if err != nil {
		// Should always be able to marshal a hardcoded request
		panic(err)
	}

	return handleRequest(body)
}

func CreateWallet(password string) (string, error) {
	body, err := json.Marshal(map[string]interface{}{
		"method": "createwallet",
		"params": []string{password},
	})
	if err != nil {
		// Should always be able to marshal a hardcoded request
		panic(err)
	}

	return handleRequest(body)
}

func LoadWalletFromSeed(seed, password string) (string, error) {
	body, err := json.Marshal(map[string]interface{}{
		"method": "createwalletfromseed",
		"params": []string{seed, password},
	})
	if err != nil {
		// Should always be able to marshal a hardcoded request
		panic(err)
	}

	return handleRequest(body)
}

func TransferDUSK(amount, address string) (string, error) {
	body, err := json.Marshal(map[string]interface{}{
		"method": "transfer",
		"params": []string{amount, address},
	})
	if err != nil {
		// Should always be able to marshal a hardcoded request
		panic(err)
	}

	return handleRequest(body)
}

func BidDUSK(amount, lockTime string) (string, error) {
	body, err := json.Marshal(map[string]interface{}{
		"method": "bid",
		"params": []string{amount, lockTime},
	})
	if err != nil {
		// Should always be able to marshal a hardcoded request
		panic(err)
	}

	return handleRequest(body)
}

func StakeDUSK(amount, lockTime string) (string, error) {
	body, err := json.Marshal(map[string]interface{}{
		"method": "stake",
		"params": []string{amount, lockTime},
	})
	if err != nil {
		// Should always be able to marshal a hardcoded request
		panic(err)
	}

	return handleRequest(body)
}

func GetBalance() (string, error) {
	body, err := json.Marshal(map[string]interface{}{
		"method": "balance",
		"params": []string{},
	})
	if err != nil {
		// Should always be able to marshal a hardcoded request
		panic(err)
	}

	return handleRequest(body)
}

func GetAddress() (string, error) {
	body, err := json.Marshal(map[string]interface{}{
		"method": "address",
		"params": []string{},
	})
	if err != nil {
		// Should always be able to marshal a hardcoded request
		panic(err)
	}

	return handleRequest(body)
}

func createRequest(body io.Reader) *http.Request {
	req, err := http.NewRequest("POST", "http://"+viper.Get("rpc.address").(string), body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-type", "application/json")
	req.SetBasicAuth(viper.Get("rpc.user").(string), viper.Get("rpc.pass").(string))

	return req
}

// Create an HTTP client for the correct network type.
func createClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}

func unmarshalResponse(resp io.ReadCloser) (map[string]interface{}, error) {
	body, err := ioutil.ReadAll(resp)
	if err != nil {
		return nil, &NetworkError{err}
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, &NetworkError{err}
	}

	return result, nil
}

func handleRequest(body []byte) (string, error) {
	req := createRequest(bytes.NewBuffer(body))
	resp, err := createClient().Do(req)
	if err != nil {
		return "", &NetworkError{err}
	}
	defer resp.Body.Close()

	result, err := unmarshalResponse(resp.Body)
	if err != nil {
		return "", err
	}

	if result["error"] != "null" {
		return result["error"].(string), &MethodError{errors.New(result["error"].(string))}
	}

	return result["result"].(string), nil
}
