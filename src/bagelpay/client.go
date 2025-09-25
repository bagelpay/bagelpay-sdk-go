package bagelpay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ClientConfig represents configuration options for BagelPayClient
type ClientConfig struct {
	// APIKey for authentication
	APIKey string
	// TestMode determines whether to use test mode (default: true)
	TestMode bool
	// BaseURL is an optional custom base URL (overrides TestMode)
	BaseURL string
	// Timeout is the request timeout duration (default: 30 seconds)
	Timeout time.Duration
	// HTTPClient is an optional custom HTTP client
	HTTPClient *http.Client
}

// BagelPayClient provides access to the BagelPay API endpoints
type BagelPayClient struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new BagelPay API client
func NewClient(config ClientConfig) *BagelPayClient {
	// Determine base URL based on test mode
	baseURL := config.BaseURL
	if baseURL == "" {
		if config.TestMode {
			baseURL = "https://test.bagelpay.io"
		} else {
			baseURL = "https://live.bagelpay.io"
		}
	}
	baseURL = strings.TrimSuffix(baseURL, "/")

	// Set default timeout
	timeout := config.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	// Use provided HTTP client or create a new one
	httpClient := config.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: timeout,
		}
	}

	return &BagelPayClient{
		baseURL:    baseURL,
		apiKey:     config.APIKey,
		httpClient: httpClient,
	}
}

// makeRequest makes an HTTP request to the API
func (c *BagelPayClient) makeRequest(ctx context.Context, method, endpoint string, data interface{}, params map[string]string) (*http.Response, error) {
	// Build URL
	u, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return nil, NewBagelPayError("invalid URL", err)
	}

	// Add query parameters
	if params != nil {
		q := u.Query()
		for key, value := range params {
			if value != "" {
				q.Add(key, value)
			}
		}
		u.RawQuery = q.Encode()
	}

	// Prepare request body
	var body io.Reader
	if data != nil && (method == "POST" || method == "PUT" || method == "PATCH") {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, NewBagelPayError("failed to marshal request data", err)
		}
		body = bytes.NewBuffer(jsonData)
	}

	// Create request
	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, NewBagelPayError("failed to create request", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "BagelPay-Go-SDK/1.0.0")
	req.Header.Set("x-api-key", c.apiKey)

	// Make request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, NewBagelPayError("request failed", err)
	}

	return resp, nil
}

// handleResponse processes the HTTP response and handles errors
func (c *BagelPayClient) handleResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NewBagelPayError("failed to read response body", err)
	}

	// Check for API errors
	if resp.StatusCode >= 400 {
		var apiError APIError
		if err := json.Unmarshal(body, &apiError); err != nil {
			// If we can't parse the error, create a generic one
			apiError = APIError{
				Code:    resp.StatusCode,
				Message: fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(body)),
			}
		}

		// Return specific error types based on status code
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			return NewBagelPayAuthenticationErrorSimple(apiError.Message, nil)
		case http.StatusBadRequest:
			return NewBagelPayValidationErrorSimple(apiError.Message, nil)
		case http.StatusNotFound:
			return NewBagelPayNotFoundErrorSimple(apiError.Message, nil)
		case http.StatusTooManyRequests:
			return NewBagelPayRateLimitErrorSimple(apiError.Message, nil)
		default:
			if resp.StatusCode >= 500 {
				return NewBagelPayServerErrorSimple(resp.StatusCode, apiError.Message, nil)
			}
			return NewBagelPayAPIError(resp.StatusCode, &apiError, nil)
		}
	}

	// Parse successful response
	if result != nil {
		if err := json.Unmarshal(body, result); err != nil {
			return NewBagelPayError("failed to parse response", err)
		}
	}

	return nil
}

// CreateCheckout creates a new checkout session
func (c *BagelPayClient) CreateCheckout(ctx context.Context, request CheckoutRequest) (*CheckoutResponse, error) {
	resp, err := c.makeRequest(ctx, "POST", "/api/payments/checkouts", request, nil)
	if err != nil {
		return nil, err
	}

	var apiResp struct {
		Data CheckoutResponse `json:"data"`
	}
	if err := c.handleResponse(resp, &apiResp); err != nil {
		return nil, err
	}

	return &apiResp.Data, nil
}

// CreateProduct creates a new product
func (c *BagelPayClient) CreateProduct(ctx context.Context, request CreateProductRequest) (*Product, error) {
	resp, err := c.makeRequest(ctx, "POST", "/api/products/create", request, nil)
	if err != nil {
		return nil, err
	}

	var apiResp struct {
		Data Product `json:"data"`
	}
	if err := c.handleResponse(resp, &apiResp); err != nil {
		return nil, err
	}

	return &apiResp.Data, nil
}

// GetProduct retrieves a product by ID
func (c *BagelPayClient) GetProduct(ctx context.Context, productID string) (*Product, error) {
	endpoint := fmt.Sprintf("/api/products/%s", productID)
	resp, err := c.makeRequest(ctx, "GET", endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	var apiResp struct {
		Data Product `json:"data"`
	}
	if err := c.handleResponse(resp, &apiResp); err != nil {
		return nil, err
	}

	return &apiResp.Data, nil
}

// ListProducts retrieves a list of products
func (c *BagelPayClient) ListProducts(ctx context.Context, pageNum, pageSize int) (*ProductListResponse, error) {
	params := make(map[string]string)
	if pageSize > 0 {
		params["pageSize"] = strconv.Itoa(pageSize)
	}
	if pageNum > 0 {
		params["pageNum"] = strconv.Itoa(pageNum)
	}

	resp, err := c.makeRequest(ctx, "GET", "/api/products/list", nil, params)
	if err != nil {
		return nil, err
	}

	var result ProductListResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateProduct updates an existing product
func (c *BagelPayClient) UpdateProduct(ctx context.Context, request UpdateProductRequest) (*Product, error) {
	resp, err := c.makeRequest(ctx, "POST", "/api/products/update", request, nil)
	if err != nil {
		return nil, err
	}

	var apiResp struct {
		Data Product `json:"data"`
	}
	if err := c.handleResponse(resp, &apiResp); err != nil {
		return nil, err
	}

	return &apiResp.Data, nil
}

// ArchiveProduct archives a product by ID
func (c *BagelPayClient) ArchiveProduct(ctx context.Context, productID string) (*Product, error) {
	endpoint := fmt.Sprintf("/api/products/%s/archive", productID)
	resp, err := c.makeRequest(ctx, "POST", endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	var apiResp struct {
		Data Product `json:"data"`
	}
	if err := c.handleResponse(resp, &apiResp); err != nil {
		return nil, err
	}

	return &apiResp.Data, nil
}

// UnarchiveProduct unarchives a product by ID
func (c *BagelPayClient) UnarchiveProduct(ctx context.Context, productID string) (*Product, error) {
	endpoint := fmt.Sprintf("/api/products/%s/unarchive", productID)
	resp, err := c.makeRequest(ctx, "POST", endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	var apiResp struct {
		Data Product `json:"data"`
	}
	if err := c.handleResponse(resp, &apiResp); err != nil {
		return nil, err
	}

	return &apiResp.Data, nil
}

// ListTransactions retrieves a list of transactions
func (c *BagelPayClient) ListTransactions(ctx context.Context, pageNum, pageSize int) (*TransactionListResponse, error) {
	params := make(map[string]string)
	if pageSize > 0 {
		params["pageSize"] = strconv.Itoa(pageSize)
	}
	if pageNum > 0 {
		params["pageNum"] = strconv.Itoa(pageNum)
	}

	resp, err := c.makeRequest(ctx, "GET", "/api/transactions/list", nil, params)
	if err != nil {
		return nil, err
	}

	var result TransactionListResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ListSubscriptions retrieves a list of subscriptions
func (c *BagelPayClient) ListSubscriptions(ctx context.Context, pageNum, pageSize int) (*SubscriptionListResponse, error) {
	params := make(map[string]string)
	if pageSize > 0 {
		params["pageSize"] = strconv.Itoa(pageSize)
	}
	if pageNum > 0 {
		params["pageNum"] = strconv.Itoa(pageNum)
	}

	resp, err := c.makeRequest(ctx, "GET", "/api/subscriptions/list", nil, params)
	if err != nil {
		return nil, err
	}

	var result SubscriptionListResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetSubscription retrieves a subscription by ID
func (c *BagelPayClient) GetSubscription(ctx context.Context, subscriptionID string) (*Subscription, error) {
	endpoint := fmt.Sprintf("/api/subscriptions/%s", subscriptionID)
	resp, err := c.makeRequest(ctx, "GET", endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	var apiResp struct {
		Data Subscription `json:"data"`
	}
	if err := c.handleResponse(resp, &apiResp); err != nil {
		return nil, err
	}

	return &apiResp.Data, nil
}

// CancelSubscription cancels a subscription by ID
func (c *BagelPayClient) CancelSubscription(ctx context.Context, subscriptionID string) (*Subscription, error) {
	endpoint := fmt.Sprintf("/api/subscriptions/%s/cancel", subscriptionID)
	resp, err := c.makeRequest(ctx, "POST", endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	var apiResp struct {
		Data Subscription `json:"data"`
	}
	if err := c.handleResponse(resp, &apiResp); err != nil {
		return nil, err
	}

	return &apiResp.Data, nil
}

// ListCustomers retrieves a list of customers
func (c *BagelPayClient) ListCustomers(ctx context.Context, pageNum, pageSize int) (*CustomerListResponse, error) {
	params := make(map[string]string)
	if pageSize > 0 {
		params["pageSize"] = strconv.Itoa(pageSize)
	}
	if pageNum > 0 {
		params["pageNum"] = strconv.Itoa(pageNum)
	}

	resp, err := c.makeRequest(ctx, "GET", "/api/customers/list", nil, params)
	if err != nil {
		return nil, err
	}

	var result CustomerListResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
