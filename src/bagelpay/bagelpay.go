/*
Package bagelpay provides a Go client library for the BagelPay API.

BagelPay is a payment processing platform that allows you to accept payments,
manage products, handle subscriptions, and track transactions.

Example usage:

	package main

	import (
		"context"
		"fmt"
		"log"
		"time"

		"github.com/bagelpay/bagelpay-sdk-go/src/bagelpay"
	)

	func main() {
		// 1. Initialize the client
		client := bagelpay.NewClient(bagelpay.ClientConfig{
			APIKey:   "your-test-api-key-here",
			TestMode: true, // Use test mode
		})

		// 2. Create a payment session
		checkoutRequest := bagelpay.CheckoutRequest{
			ProductID: "prod_123456789",
			RequestID: bagelpay.StringPtr(fmt.Sprintf("req_%d", time.Now().Unix())),
			Units:     bagelpay.StringPtr("1"),
			Customer: &bagelpay.Customer{
				Email: "customer@example.com",
			},
			SuccessURL: bagelpay.StringPtr("https://yoursite.com/success"),
			Metadata: map[string]interface{}{
				"order_id": fmt.Sprintf("req_%d", time.Now().Unix()),
			},
		}

		// 3. Get payment URL
		ctx := context.Background()
		response, err := client.CreateCheckout(ctx, checkoutRequest)
		if err != nil {
			log.Fatalf("Failed to create checkout: %v", err)
		}

		fmt.Printf("Payment URL: %s\n", *response.CheckoutURL)
	}

For more examples and detailed documentation, visit:
https://github.com/bagelpay/bagelpay-sdk-go
*/
package bagelpay

import "time"

const (
	// Version represents the current version of the SDK
	Version = "1.0.3"
	// Author represents the SDK author
	Author = "andrew@gettrust.ai"
	// Email represents the support email
	Email = "support@bagelpayment.com"
)

// Default configuration values
const (
	// DefaultTestBaseURL is the default base URL for test mode
	DefaultTestBaseURL = "https://test.bagelpay.io"
	// DefaultLiveBaseURL is the default base URL for live mode
	DefaultLiveBaseURL = "https://live.bagelpay.io"
	// DefaultTimeout is the default request timeout
	DefaultTimeout = 30 * time.Second
	// DefaultUserAgent is the default user agent string
	DefaultUserAgent = "BagelPay-Go-SDK/1.0.0"
)

// NewDefaultClient creates a new BagelPay client with default configuration
// This is a convenience function for quick setup.
func NewDefaultClient(apiKey string) *BagelPayClient {
	return NewClient(ClientConfig{
		APIKey:   apiKey,
		TestMode: true,
	})
}

// NewTestClient creates a new BagelPay client configured for test mode
func NewTestClient(apiKey string) *BagelPayClient {
	return NewClient(ClientConfig{
		APIKey:   apiKey,
		TestMode: true,
	})
}

// NewLiveClient creates a new BagelPay client configured for live mode
func NewLiveClient(apiKey string) *BagelPayClient {
	return NewClient(ClientConfig{
		APIKey:   apiKey,
		TestMode: false,
	})
}
