package stakingwalletrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// RPCClient struct to hold HTTP client and server details
type RPCClient struct {
	client   *http.Client
	url      string
	username string
	password string
}

// NewRPCClient creates a new JSON-RPC client with authentication
func NewRPCClient(host string, port int, username, password string) (*RPCClient, error) {
	// Create HTTP client
	httpClient := &http.Client{}

	// Format the server URL
	url := fmt.Sprintf("http://%s:%d", host, port)

	return &RPCClient{
		client:   httpClient,
		url:      url,
		username: username,
		password: password,
	}, nil
}

// Call makes a JSON-RPC call to the specified method with given arguments
func (c *RPCClient) call(method string, params interface{}, result interface{}) error {
	// Create JSON-RPC request payload
	request := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      1,
	}

	// Marshal request to JSON
	body, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %v", err)
	}

	// Create HTTP POST request
	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.username, c.password)

	// Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// Parse response
	var rpcResponse struct {
		Result json.RawMessage   `json:"result"`
		Error  map[string]string `json:"error"`
		ID     int               `json:"id"`
	}

	err = json.NewDecoder(resp.Body).Decode(&rpcResponse)
	if err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	// Check for JSON-RPC error
	if rpcResponse.Error != nil {
		return fmt.Errorf("JSON-RPC error: %v", rpcResponse.Error)
	}

	// Unmarshal result into provided result interface
	if result != nil {
		err = json.Unmarshal(rpcResponse.Result, result)
		if err != nil {
			return fmt.Errorf("failed to unmarshal result: %v", err)
		}
	}

	return nil
}

// Close is a no-op for HTTP client (included for consistency)
func (c *RPCClient) Close() error {
	c.client.CloseIdleConnections()
	return nil
}

// Gets wallet information 
func (c *RPCClient) GetInfo() (info Info, err error) {
	params := []interface{}{}
	err = c.call("getinfo", params, &info)
	return info, err
}

// Get transaction by TXID to fetch transaction information
func (c *RPCClient) GetTransaction(txid string) (transaction Transaction, err error) {
	params := []interface{}{txid}
	err = c.call("gettransaction", params, &transaction)
	return transaction, err
}

// Sends coins from the wallet to a specified address
func (c *RPCClient) SendToAddress(address string, amount float64, comment string) (txid string, err error) {
	params := []interface{}{address, amount, comment}
	err = c.call("sendtoaddress", params, &txid)
	return txid, err
}

// Generates a new address with the specified label
func (c *RPCClient) GetNewAddress(label string) (address string, err error) {
	params := []interface{}{"label"}
	err = c.call("getnewaddress", params, &address)
	return address, err
}

// Validates if the address is good for this type of wallet (a bitcoin address doesn't work for a ETH wallet)
func (c *RPCClient) ValidateAddress(address string) (validate ValidateAddress, err error) {
	params := []interface{}{address}
	err = c.call("validateaddress", params, &validate)
	return validate, err
}
