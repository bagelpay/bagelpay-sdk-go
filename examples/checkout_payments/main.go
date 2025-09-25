/*
BagelPay SDK - Checkout Payments Example

This example demonstrates various checkout payment scenarios:
- Simple one-time payment checkout
- Subscription checkout
- Checkout with customer information
- Checkout with metadata and custom URLs

Prerequisites:
- Set your API key as an environment variable: export BAGELPAY_API_KEY="your_api_key_here"
- Install the SDK: go mod tidy

To run this example:
go run checkout_payments.go
*/

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bagelpay/bagelpay-sdk-go/src/bagelpay"
)

func main() {
	fmt.Println("🚀 BagelPay SDK - Checkout Payments Example")
	fmt.Println("==========================================")

	// Get API key from environment variable
	apiKey := os.Getenv("BAGELPAY_API_KEY")
	if apiKey == "" {
		// log.Fatal("❌ BAGELPAY_API_KEY environment variable is required")
		apiKey = "bagel_test_C6D1E83B94204A00A6F8EFD2AF05B427"
	}

	// Initialize the BagelPay client
	fmt.Println("\n📡 Initializing BagelPay client...")
	client := bagelpay.NewTestClient(apiKey)

	ctx := context.Background()

	// Create a sample product first
	fmt.Println("\n📦 Creating sample products...")
	digitalProduct, err := createDigitalProduct(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to create digital product: %v", err)
		return
	}
	if digitalProduct != nil && digitalProduct.ProductID != nil {
		fmt.Printf("✅ Digital product created! ID: %s\n", *digitalProduct.ProductID)
	} else {
		fmt.Println("✅ Digital product created, but ID not available")
	}

	subscriptionProduct, err := createSubscriptionProduct(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to create subscription product: %v", err)
		return
	}
	if subscriptionProduct != nil && subscriptionProduct.ProductID != nil {
		fmt.Printf("✅ Subscription product created! ID: %s\n", *subscriptionProduct.ProductID)
	} else {
		fmt.Println("✅ Subscription product created, but ID not available")
	}

	// Example 1: Simple one-time payment checkout
	fmt.Println("\n💳 Creating simple checkout session...")
	if digitalProduct != nil && digitalProduct.ProductID != nil {
		if err := createSimpleCheckout(ctx, client, *digitalProduct.ProductID); err != nil {
			log.Printf("❌ Failed to create simple checkout: %v", err)
		}
	}

	// Example 2: Subscription checkout
	fmt.Println("\n🔄 Creating subscription checkout session...")
	if subscriptionProduct != nil && subscriptionProduct.ProductID != nil {
		if err := createSubscriptionCheckout(ctx, client, *subscriptionProduct.ProductID); err != nil {
			log.Printf("❌ Failed to create subscription checkout: %v", err)
		}
	}

	// Example 3: Checkout with customer information
	fmt.Println("\n👤 Creating checkout with customer info...")
	if digitalProduct != nil && digitalProduct.ProductID != nil {
		if err := createCheckoutWithCustomer(ctx, client, *digitalProduct.ProductID); err != nil {
			log.Printf("❌ Failed to create checkout with customer: %v", err)
		}
	}

	// Example 4: Checkout with metadata and custom URLs
	fmt.Println("\n🏷️ Creating checkout with metadata...")
	if digitalProduct != nil && digitalProduct.ProductID != nil {
		if err := createCheckoutWithMetadata(ctx, client, *digitalProduct.ProductID); err != nil {
			log.Printf("❌ Failed to create checkout with metadata: %v", err)
		}
	}

	fmt.Println("\n🎉 Checkout payments examples completed successfully!")
}

// createDigitalProduct creates a sample digital product
func createDigitalProduct(ctx context.Context, client *bagelpay.BagelPayClient) (*bagelpay.Product, error) {
	productRequest := bagelpay.CreateProductRequest{
		Name:              "Premium E-book-single",
		Description:       "A comprehensive guide to payment processing",
		Price:             19.99, // $19.99
		Currency:          "USD",
		BillingType:       "single_payment",
		TaxInclusive:      false,
		TaxCategory:       "digital_products",
		RecurringInterval: "",
		TrialDays:         0,
	}

	return client.CreateProduct(ctx, productRequest)
}

// createSubscriptionProduct creates a sample subscription product
func createSubscriptionProduct(ctx context.Context, client *bagelpay.BagelPayClient) (*bagelpay.Product, error) {
	productRequest := bagelpay.CreateProductRequest{
		Name:              "Monthly Newsletter Subscription",
		Description:       "Monthly premium newsletter with industry insights",
		Price:             9.99, // $9.99
		Currency:          "USD",
		BillingType:       "subscription",
		TaxInclusive:      false,
		TaxCategory:       "digital_products",
		RecurringInterval: "monthly",
		TrialDays:         7,
	}

	return client.CreateProduct(ctx, productRequest)
}

// createSimpleCheckout creates a simple one-time payment checkout
func createSimpleCheckout(ctx context.Context, client *bagelpay.BagelPayClient, productID string) error {
	successURL := "https://example.com/success"
	checkoutRequest := bagelpay.CheckoutRequest{
		ProductID:  productID,
		SuccessURL: &successURL,
	}

	checkout, err := client.CreateCheckout(ctx, checkoutRequest)
	if err != nil {
		return err
	}

	if checkout.CheckoutURL != nil {
		fmt.Printf("✅ Simple checkout created! URL: %s\n", *checkout.CheckoutURL)
	}
	if checkout.PaymentID != nil {
		fmt.Printf("   Payment ID: %s\n", *checkout.PaymentID)
	}
	return nil
}

// createSubscriptionCheckout creates a subscription checkout
func createSubscriptionCheckout(ctx context.Context, client *bagelpay.BagelPayClient, productID string) error {
	successURL := "https://example.com/subscription/success"
	metadata := map[string]interface{}{
		"subscription_type": "monthly",
		"trial_period":      "7_days",
	}
	checkoutRequest := bagelpay.CheckoutRequest{
		ProductID:  productID,
		SuccessURL: &successURL,
		Metadata:   metadata,
	}

	checkout, err := client.CreateCheckout(ctx, checkoutRequest)
	if err != nil {
		return err
	}

	if checkout.CheckoutURL != nil {
		fmt.Printf("✅ Subscription checkout created! URL: %s\n", *checkout.CheckoutURL)
	}
	if checkout.PaymentID != nil {
		fmt.Printf("   Payment ID: %s\n", *checkout.PaymentID)
	}
	return nil
}

// createCheckoutWithCustomer creates a checkout with customer information
func createCheckoutWithCustomer(ctx context.Context, client *bagelpay.BagelPayClient, productID string) error {
	successURL := "https://example.com/success"
	customer := &bagelpay.Customer{
		Email: "john.doe@example.com",
	}
	metadata := map[string]interface{}{
		"customer_type": "premium",
		"referral_code": "FRIEND2023",
	}
	checkoutRequest := bagelpay.CheckoutRequest{
		ProductID:  productID,
		SuccessURL: &successURL,
		Customer:   customer,
		Metadata:   metadata,
	}

	checkout, err := client.CreateCheckout(ctx, checkoutRequest)
	if err != nil {
		return err
	}

	if checkout.CheckoutURL != nil {
		fmt.Printf("✅ Customer checkout created! URL: %s\n", *checkout.CheckoutURL)
	}
	if checkout.PaymentID != nil {
		fmt.Printf("   Payment ID: %s\n", *checkout.PaymentID)
	}
	fmt.Printf("   Customer: %s\n", customer.Email)
	return nil
}

// createCheckoutWithMetadata creates a checkout with extensive metadata
func createCheckoutWithMetadata(ctx context.Context, client *bagelpay.BagelPayClient, productID string) error {
	successURL := "https://mystore.com/order/success?session_id={CHECKOUT_SESSION_ID}"
	metadata := map[string]interface{}{
		"order_id":         "ORD-2023-001",
		"campaign":         "summer_sale",
		"discount_code":    "SAVE20",
		"affiliate_id":     "AFF123",
		"source":           "website",
		"customer_segment": "returning",
	}
	checkoutRequest := bagelpay.CheckoutRequest{
		ProductID:  productID,
		SuccessURL: &successURL,
		Metadata:   metadata,
	}

	checkout, err := client.CreateCheckout(ctx, checkoutRequest)
	if err != nil {
		return err
	}

	if checkout.CheckoutURL != nil {
		fmt.Printf("✅ Metadata checkout created! URL: %s\n", *checkout.CheckoutURL)
	}
	if checkout.PaymentID != nil {
		fmt.Printf("   Payment ID: %s\n", *checkout.PaymentID)
	}
	fmt.Printf("   Order ID: %s\n", metadata["order_id"])
	return nil
}
