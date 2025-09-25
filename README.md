# BagelPay Go SDK

Official Go SDK for the BagelPay payment platform. Easily integrate subscription billing, one-time payments, and customer management into your Go applications.

## 🔗 Related Links

- 🌏 **BagelPay Website**: [https://bagelpay.io](https://bagelpay.io)
- 🌏 **Developer Dashboard**: [https://app.bagelpay.io/dashboard](https://app.bagelpay.io/dashboard)
- 📖 **Official Documentation**: [https://bagelpay.gitbook.io/docs](https://bagelpay.gitbook.io/docs)
- 📖 **API Documentation**: [https://bagelpay.gitbook.io/docs/apireference](https://bagelpay.gitbook.io/docs/apireference)
- 📧 **Technical Support**: support@bagelpayment.com
- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/bagelpay/bagelpay-sdk-go/issues)

## 🚀 Quick Start

### 30-Second Quick Demo

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bagelpay/bagelpay-sdk-go/src/bagelpay"
)

func main() {
	// Initialize the client
	client := bagelpay.NewClient(bagelpay.ClientConfig{
		APIKey:   "your-api-key",
		TestMode: true,
	})

	ctx := context.Background()

	// Create a product
	product, err := client.CreateProduct(ctx, bagelpay.CreateProductRequest{
		Name:              fmt.Sprintf("Premium Subscription %d", time.Now().Unix()),
		Description:       "Monthly premium subscription with unique identifier",
		Price:             29.99,
		Currency:          "USD",
		BillingType:       "subscription",
		TaxInclusive:      false,
		TaxCategory:       "digital_products",
		RecurringInterval: "monthly",
		TrialDays:         7,
	})

	if err != nil {
		log.Fatalf("Failed to create product: %v", err)
	}

	fmt.Printf("✅ Product URL: %s\n", *product.ProductURL)
}
```

## 📦 Installation Guide

```bash
go get github.com/bagelpay/bagelpay-sdk-go
```

## Core Features

### 🛍️ Product Management
- Create and manage products
- Support for both subscriptions and one-time payments
- Flexible pricing and billing intervals
- Tax configuration options

### 💳 Checkout & Payments
- Secure checkout session creation
- Customizable success/cancel URLs
- Metadata support for tracking
- Real-time payment status

### 👥 Customer Management
- Customer creation and retrieval
- Subscription management
- Payment history tracking

### 📊 Analytics & Reporting
- Transaction listing and filtering
- Subscription analytics
- Customer insights

## API Reference

### Client Initialization

```go
client := bagelpay.NewClient(bagelpay.ClientConfig{
	APIKey:     "your-api-key",        // Your BagelPay API key
	TestMode:   true,                  // Default: true
	Timeout:    30 * time.Second,      // Default: 30 seconds
	BaseURL:    "custom-url",          // Optional custom base URL
	HTTPClient: &http.Client{},        // Optional custom HTTP client
})
```

### Convenience Constructors

```go
// Test mode client (recommended for development)
testClient := bagelpay.NewTestClient("your-test-api-key")

// Live mode client (for production)
liveClient := bagelpay.NewLiveClient("your-live-api-key")

// Default client (test mode)
defaultClient := bagelpay.NewDefaultClient("your-api-key")
```

### Products

#### Create Product
```go
product, err := client.CreateProduct(ctx, bagelpay.CreateProductRequest{
	Name:              "Premium Plan",
	Description:       "Monthly premium subscription",
	Price:             29.99,
	Currency:          "USD",
	BillingType:       "subscription", // or "single_payment"
	TaxInclusive:      false,
	TaxCategory:       "digital_products",
	RecurringInterval: "monthly", // daily, weekly, monthly, 3months, 6months
	TrialDays:         7,
})
```

#### List Products
```go
products, err := client.ListProducts(ctx, pageNum, pageSize)
```

#### Get Product
```go
product, err := client.GetProduct(ctx, productID)
```

#### Update Product
```go
product, err := client.UpdateProduct(ctx, bagelpay.UpdateProductRequest{
	ProductID:         "prod_123456789",
	Name:              "Updated Premium Plan",
	Description:       "Updated description",
	Price:             39.99,
	Currency:          "USD",
	BillingType:       "subscription",
	TaxInclusive:      true,
	TaxCategory:       "digital_services",
	RecurringInterval: "monthly",
	TrialDays:         14,
})
```

#### Archive/Unarchive Product
```go
// Archive product
product, err := client.ArchiveProduct(ctx, productID)

// Unarchive product
product, err := client.UnarchiveProduct(ctx, productID)
```

### Checkout

#### Create Checkout Session
```go
checkout, err := client.CreateCheckout(ctx, bagelpay.CheckoutRequest{
	ProductID: "prod_123456789",
	Customer: &bagelpay.Customer{
		Email: "customer@example.com",
	},
	RequestID:  bagelpay.StringPtr("unique-request-id"),
	Units:      bagelpay.StringPtr("1"),
	SuccessURL: bagelpay.StringPtr("https://yoursite.com/success"),
	Metadata: map[string]interface{}{
		"order_id": "order_123",
		"user_id":  "user_456",
	},
})
```

### Transactions

#### List Transactions
```go
transactions, err := client.ListTransactions(ctx, pageNum, pageSize)
```

### Subscriptions

#### List Subscriptions
```go
subscriptions, err := client.ListSubscriptions(ctx, pageNum, pageSize)
```

#### Get Subscription
```go
subscription, err := client.GetSubscription(ctx, subscriptionID)
```

#### Cancel Subscription
```go
subscription, err := client.CancelSubscription(ctx, subscriptionID)
```

### Customers

#### List Customers
```go
customers, err := client.ListCustomers(ctx, pageNum, pageSize)
```

## Error Handling

The SDK provides specific error types for better error handling:

```go
import "github.com/bagelpay/bagelpay-sdk-go/src/bagelpay"

product, err := client.CreateProduct(ctx, productData)
if err != nil {
	switch {
	case bagelpay.IsAuthenticationError(err):
		fmt.Println("Authentication failed - check your API key")
	case bagelpay.IsValidationError(err):
		fmt.Println("Validation error - check your request data")
	case bagelpay.IsNotFoundError(err):
		fmt.Println("Resource not found")
	case bagelpay.IsRateLimitError(err):
		fmt.Println("Rate limit exceeded - please retry later")
	case bagelpay.IsServerError(err):
		fmt.Println("Server error - please try again")
	default:
		fmt.Printf("Unexpected error: %v\n", err)
	}
}
```

### Error Types

- `BagelPayError`: Base error type
- `BagelPayAPIError`: API-specific errors
- `BagelPayAuthenticationError`: Authentication failures (401)
- `BagelPayValidationError`: Request validation errors (400)
- `BagelPayNotFoundError`: Resource not found errors (404)
- `BagelPayRateLimitError`: Rate limit exceeded (429)
- `BagelPayServerError`: Server-side errors (5xx)

## Go Type Support

The SDK provides full Go type definitions and helper functions:

```go
// Helper functions for pointer types
name := bagelpay.StringPtr("Product Name")
quantity := bagelpay.IntPtr(5)
price := bagelpay.Float64Ptr(29.99)
taxInclusive := bagelpay.BoolPtr(true)

// JSON conversion utilities
jsonStr, err := bagelpay.ToJSON(product)
err = bagelpay.FromJSON(jsonStr, &product)
```

## Environment Configuration

### Test Mode
Use test mode for development and testing:

```go
client := bagelpay.NewClient(bagelpay.ClientConfig{
	APIKey:   "bagel_test_your_test_key",
	TestMode: true,
})
```

### Production Mode
For production environments:

```go
client := bagelpay.NewClient(bagelpay.ClientConfig{
	APIKey:   "bagel_live_your_live_key",
	TestMode: false,
})
```

## 🚀 Webhook Integration

```go
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const WEBHOOK_SECRET = "your_webhook_key"

func verifyWebhookSignature(payload []byte, timestamp, signature, secret string) bool {
	// Combine timestamp and payload
	signatureData := fmt.Sprintf("%s.%s", timestamp, string(payload))
	
	// Create HMAC
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(signatureData))
	expectedSignature := hex.EncodeToString(h.Sum(nil))
	
	return hmac.Equal([]byte(expectedSignature), []byte(signature))
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Get headers
	timestamp := r.Header.Get("timestamp")
	signature := r.Header.Get("Bagelpay-Signature")

	// Verify signature
	if !verifyWebhookSignature(payload, timestamp, signature, WEBHOOK_SECRET) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	// Parse the webhook event
	var event map[string]interface{}
	if err := json.Unmarshal(payload, &event); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	eventType, ok := event["event_type"].(string)
	if !ok {
		http.Error(w, "Missing event_type", http.StatusBadRequest)
		return
	}

	// Handle different event types
	switch eventType {
	case "checkout.completed":
		fmt.Println("Checkout completed:", event)
	case "checkout.failed":
		fmt.Println("Checkout failed:", event)
	case "checkout.cancel":
		fmt.Println("Checkout cancelled:", event)
	case "subscription.trialing":
		fmt.Println("Subscription trialing:", event)
	case "subscription.paid":
		fmt.Println("Subscription paid:", event)
	case "subscription.canceled":
		fmt.Println("Subscription cancelled:", event)
	case "refund.created":
		fmt.Println("Refund created:", event)
	default:
		fmt.Printf("Unhandled event type: %s\n", eventType)
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Success"}`))
}

func main() {
	http.HandleFunc("/api/webhooks", webhookHandler)
	
	fmt.Println("Webhook server starting on :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
```

## Examples

The SDK includes comprehensive examples in the `examples/` directory:

- **Basic Usage** (`examples/basic_usage/`): Complete SDK functionality demonstration
- **Checkout Payments** (`examples/checkout_payments/`): Various checkout scenarios
- **Product Management** (`examples/product_management/`): Product CRUD operations
- **Subscription & Customer Management** (`examples/subscription_customer_management/`): Subscription and customer operations

To run an example:

```bash
cd examples/basic_usage
go run main.go
```

## Support

- **Documentation**: [https://bagelpay.gitbook.io/docs/documentation](https://bagelpay.gitbook.io/docs/documentation)
- **API Reference**: [https://bagelpay.gitbook.io/docs/apireference](https://bagelpay.gitbook.io/docs/apireference)
- **Support**: [support@bagelpayment.com](mailto:support@bagelpayment.com)
- **GitHub Issues**: [Report bugs and feature requests](https://github.com/bagelpay/bagelpay-sdk-go/issues)

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

---

**Ready to get started?** [Sign up for a BagelPay account](https://bagelpay.io) and get your API keys today!