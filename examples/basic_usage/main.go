/*
BagelPay SDK - Basic Usage Example

This example demonstrates the basic usage of the BagelPay SDK for Go.
It shows how to:
1. Initialize the client
2. Create a product
3. List products
4. Get product details
5. Create a checkout session
6. List transactions

Prerequisites:
- Set your API key as an environment variable: export BAGELPAY_API_KEY="your_api_key_here"
- Install the SDK: go mod tidy

To run this example:
go run basic_usage.go
*/

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/bagelpay/bagelpay-sdk-go/src/bagelpay"
)

func main() {
	fmt.Println("🚀 BagelPay SDK - Basic Usage Example")
	fmt.Println("=====================================")

	// Get API key from environment variable
	apiKey := os.Getenv("BAGELPAY_API_KEY")
	if apiKey == "" {
		apiKey = "bagel_test_C6D1E83B94204A00A6F8EFD2AF05B427"
	}
	if apiKey == "" {
		log.Fatal("❌ BAGELPAY_API_KEY environment variable is required")
	}

	// Initialize the BagelPay client
	fmt.Println("\n📡 Initializing BagelPay client...")
	client := bagelpay.NewTestClient(apiKey)
	fmt.Printf("✅ BagelPay client initialized\n")
	fmt.Printf("   Mode: Test\n")
	fmt.Printf("   Base URL: https://test.bagelpay.io\n")
	fmt.Printf("   API Key: %s...%s\n", apiKey[:8], apiKey[len(apiKey)-4:])

	ctx := context.Background()

	// Example 1: Create a product
	fmt.Println("\n=== Example 1: Create a product ===")
	product, err := createSampleProduct(ctx, client)
	if err != nil {
		fmt.Printf("Error creating product: %v\n", err)
	} else {
		fmt.Printf("✓ Product created successfully!\n")
		if product.ProductID != nil {
			fmt.Printf("  Product ID: %s\n", *product.ProductID)
		} else {
			fmt.Printf("  Product ID: Not available\n")
		}
		if product.BillingType != nil {
			fmt.Printf("  Billing Type: %s\n", *product.BillingType)
		}
		if product.Name != nil {
			fmt.Printf("  Name: %s\n", *product.Name)
		}
		if product.Price != nil {
			fmt.Printf("  Price: $%.2f\n", *product.Price)
		}
		if product.Currency != nil {
			fmt.Printf("  Currency: %s\n", *product.Currency)
		}
		if product.RecurringInterval != nil && *product.RecurringInterval != "" {
			fmt.Printf("  Recurring Interval: %s\n", *product.RecurringInterval)
		}
		if product.ProductURL != nil {
			fmt.Printf("  Product URL: %s\n", *product.ProductURL)
		}
	}

	// Example 2: List products
	fmt.Println("\n📋 Listing products...")
	err = listProducts(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to list products: %v", err)
		return
	}

	// Example 3: Get product details
	productID := ""
	if product.ProductID != nil {
		productID = *product.ProductID
		fmt.Printf("\n🔍 Getting details for product %s...\n", productID)
		err = getProductDetails(ctx, client, productID)
		if err != nil {
			log.Printf("❌ Failed to get product details: %v", err)
			return
		}
	}

	// Example 4: Create a checkout session
	fmt.Println("\n=== Example 4: Create a checkout session ===")
	if product != nil && product.ProductID != nil {
		checkout, err := createCheckoutSession(ctx, client, *product.ProductID)
		if err != nil {
			fmt.Printf("Error creating checkout: %v\n", err)
		} else {
			fmt.Printf("✓ Checkout session created successfully!\n")
			if checkout.PaymentID != nil {
				fmt.Printf("  Payment ID: %s\n", *checkout.PaymentID)
			} else {
				fmt.Printf("  Payment ID: Not available\n")
			}
			if checkout.Status != nil {
				fmt.Printf("  Status: %s\n", *checkout.Status)
			}
			if checkout.CheckoutURL != nil {
				fmt.Printf("  Checkout URL: %s\n", *checkout.CheckoutURL)
			}
			if checkout.ExpiresOn != nil {
				fmt.Printf("  Expires On: %s\n", *checkout.ExpiresOn)
			}
			if checkout.SuccessURL != nil {
				fmt.Printf("  Success URL: %s\n", *checkout.SuccessURL)
			}
		}
	} else {
		fmt.Println("Skipping checkout creation - no product ID available")
	}

	// Example 5: List customers
	fmt.Println("\n👥 Listing customers...")
	err = listCustomers(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to list customers: %v", err)
		return
	}

	// Example 6: Update product
	fmt.Println("\n✏️ Updating product...")
	err = updateProduct(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to update product: %v", err)
		return
	}

	// Example 7: Archive product
	fmt.Println("\n📦 Archiving product...")
	err = archiveProduct(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to archive product: %v", err)
		return
	}

	// Example 8: Unarchive product
	fmt.Println("\n📦 Unarchiving product...")
	err = unarchiveProduct(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to unarchive product: %v", err)
		return
	}

	// Example 9: List subscriptions
	fmt.Println("\n📋 Listing subscriptions...")
	err = listSubscriptions(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to list subscriptions: %v", err)
		return
	}

	// Example 10: Get subscription details
	fmt.Println("\n🔍 Getting subscription details...")
	err = getSubscriptionDetails(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to get subscription details: %v", err)
		return
	}

	// Example 11: Cancel subscription
	fmt.Println("\n❌ Cancelling subscription...")
	err = cancelSubscription(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to cancel subscription: %v", err)
		return
	}

	// Example 12: List transactions
	fmt.Println("\n💰 Listing recent transactions...")
	err = listRecentTransactions(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to list transactions: %v", err)
		return
	}

	fmt.Println("\n🎉 Basic usage example completed successfully!")
}

// createSampleProduct creates a sample product with randomized data
func createSampleProduct(ctx context.Context, client *bagelpay.BagelPayClient) (*bagelpay.Product, error) {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Generate random data similar to TypeScript version
	randomNum := rand.Intn(9000) + 1000
	price := rand.Float64()*(1024.5-50.5) + 50.5

	billingTypes := []string{"subscription", "single_payment"}
	taxCategories := []string{"digital_products", "saas_services", "ebooks"}
	recurringIntervals := []string{"daily", "weekly", "monthly", "3months", "6months"}
	trialDaysOptions := []int{0, 1, 7}

	billingType := billingTypes[rand.Intn(len(billingTypes))]
	taxCategory := taxCategories[rand.Intn(len(taxCategories))]
	recurringInterval := ""
	trialDays := 0

	// Set recurring interval and trial days only for subscription products
	if billingType == "subscription" {
		recurringInterval = recurringIntervals[rand.Intn(len(recurringIntervals))]
		trialDays = trialDaysOptions[rand.Intn(len(trialDaysOptions))]
	}

	productRequest := bagelpay.CreateProductRequest{
		Name:              fmt.Sprintf("Product_GO_%d", randomNum),
		Description:       fmt.Sprintf("Description_of_product_%d", randomNum),
		Price:             price,
		Currency:          "USD",
		BillingType:       billingType,
		TaxInclusive:      false,
		TaxCategory:       taxCategory,
		RecurringInterval: recurringInterval,
		TrialDays:         trialDays,
	}

	return client.CreateProduct(ctx, productRequest)
}

// createCheckoutSession creates a sample checkout session with detailed customer info
func createCheckoutSession(ctx context.Context, client *bagelpay.BagelPayClient, productID string) (*bagelpay.CheckoutResponse, error) {
	// Generate random data for customer
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(9000) + 1000

	customer := &bagelpay.Customer{
		Email: fmt.Sprintf("customer_%d@example.com", randomNum),
	}

	checkoutRequest := bagelpay.CheckoutRequest{
		ProductID:  productID,
		RequestID:  bagelpay.StringPtr(fmt.Sprintf("req_go_%d", time.Now().Unix())),
		Units:      bagelpay.StringPtr("1"),
		Customer:   customer,
		SuccessURL: bagelpay.StringPtr("https://example.com/success"),
		Metadata: map[string]interface{}{
			"order_id":    fmt.Sprintf("order_%d", randomNum),
			"source":      "go_sdk",
			"customer_id": fmt.Sprintf("cust_%d", randomNum),
		},
	}

	return client.CreateCheckout(ctx, checkoutRequest)
}

// listProducts lists all products
func listProducts(ctx context.Context, client *bagelpay.BagelPayClient) error {
	response, err := client.ListProducts(ctx, 1, 5)
	if err != nil {
		return err
	}

	if len(response.Items) == 0 {
		fmt.Println("📝 No products found")
		return nil
	}

	fmt.Printf("📝 Found %d product(s):\n", len(response.Items))
	for i, product := range response.Items {
		name := "N/A"
		if product.Name != nil {
			name = *product.Name
		}

		productID := "N/A"
		if product.ProductID != nil {
			productID = *product.ProductID
		}

		price := 0.0
		if product.Price != nil {
			price = *product.Price
		}

		currency := "N/A"
		if product.Currency != nil {
			currency = *product.Currency
		}

		// Check if it's a subscription product and get billing cycle
		priceDisplay := fmt.Sprintf("%.2f %s", price, currency)
		if product.BillingType != nil && *product.BillingType == "subscription" {
			if product.RecurringInterval != nil && *product.RecurringInterval != "" {
				priceDisplay += fmt.Sprintf(" / %s", *product.RecurringInterval)
			}
		}

		fmt.Printf("   %d. %s (ID: %s)\n", i+1, name, productID)
		fmt.Printf("      Price: %s\n", priceDisplay)
		fmt.Println()
	}

	return nil
}

// getProductDetails retrieves and displays product details
func getProductDetails(ctx context.Context, client *bagelpay.BagelPayClient, productID string) error {
	product, err := client.GetProduct(ctx, productID)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Product Details:\n")

	if product.ProductID != nil {
		fmt.Printf("   ID: %s\n", *product.ProductID)
	}
	if product.Name != nil {
		fmt.Printf("   Name: %s\n", *product.Name)
	}
	if product.Description != nil {
		fmt.Printf("   Description: %s\n", *product.Description)
	}
	if product.BillingType != nil {
		fmt.Printf("   Billing Type: %s\n", *product.BillingType)
	}
	if product.Price != nil && product.Currency != nil {
		fmt.Printf("   Price: %.2f %s\n", *product.Price, *product.Currency)
	}
	if product.RecurringInterval != nil && *product.RecurringInterval != "" {
		fmt.Printf("   Recurring Interval: %s\n", *product.RecurringInterval)
	}

	return nil
}

// listCustomers lists customers
func listCustomers(ctx context.Context, client *bagelpay.BagelPayClient) error {
	response, err := client.ListCustomers(ctx, 1, 10) // pageNum=1, pageSize=10
	if err != nil {
		return err
	}

	customers := response.Items

	if len(customers) == 0 {
		fmt.Println("📝 No customers found")
		return nil
	}

	fmt.Printf("✅ Found %d customers\n", len(response.Items))
	for i, customer := range customers {
		id := "N/A"
		if customer.ID != nil {
			id = fmt.Sprintf("%d", *customer.ID)
		}

		name := "N/A"
		if customer.Name != nil {
			name = *customer.Name
		}

		email := "N/A"
		if customer.Email != nil {
			email = *customer.Email
		}

		totalSpend := 0.0
		if customer.TotalSpend != nil {
			totalSpend = *customer.TotalSpend
		}

		subscriptions := 0
		if customer.Subscriptions != nil {
			subscriptions = *customer.Subscriptions
		}

		payments := 0
		if customer.Payments != nil {
			payments = *customer.Payments
		}

		fmt.Printf("   %d. ID: %s, Name: %s, Email: %s\n",
			i+1, id, name, email)
		fmt.Printf("      Total Spend: $%.2f, Subscriptions: %d, Payments: %d\n",
			totalSpend, subscriptions, payments)
	}

	return nil
}

// listRecentTransactions lists recent transactions
func listRecentTransactions(ctx context.Context, client *bagelpay.BagelPayClient) error {
	response, err := client.ListTransactions(ctx, 1, 10) // pageNum=1, pageSize=10
	if err != nil {
		return err
	}

	transactions := response.Items

	if len(transactions) == 0 {
		fmt.Println("📝 No transactions found")
		return nil
	}

	fmt.Printf("✅ Found %d transactions\n", len(response.Items))
	for i, transaction := range transactions {
		id := "N/A"
		if transaction.TransactionID != nil {
			id = *transaction.TransactionID
		}

		amount := 0.0
		if transaction.Amount != nil {
			amount = *transaction.Amount
		}

		currency := "N/A"
		if transaction.Currency != nil {
			currency = *transaction.Currency
		}

		txType := "N/A"
		if transaction.Type != nil {
			txType = *transaction.Type
		}

		fmt.Printf("   %d. ID: %s, Amount: %.2f %s, Type: %s\n",
			i+1, id, amount, currency, txType)
	}

	return nil
}

// Example 6: Update Product
func updateProduct(ctx context.Context, client *bagelpay.BagelPayClient) error {
	fmt.Println("\n=== Example 6: Update Product ===")

	// First, get a product to update (using the first product from our list)
	response, err := client.ListProducts(ctx, 1, 1)
	if err != nil {
		fmt.Printf("Error listing products: %v\n", err)
		return err
	}

	if len(response.Items) == 0 {
		fmt.Println("No products available to update")
		return nil
	}

	productToUpdate := response.Items[0]
	if productToUpdate.ProductID == nil {
		fmt.Println("Product ID is nil, cannot update")
		return nil
	}

	// Create update request with modified data
	updateRequest := bagelpay.UpdateProductRequest{
		ProductID:         *productToUpdate.ProductID,
		Name:              "Updated Product Name",
		Description:       "This product has been updated via the Go SDK",
		Price:             29.99,
		Currency:          "USD",
		BillingType:       "single_payment",
		TaxInclusive:      true,
		TaxCategory:       "digital_products",
		RecurringInterval: "",
		TrialDays:         0,
	}

	// Update the product
	updatedProduct, err := client.UpdateProduct(ctx, updateRequest)
	if err != nil {
		fmt.Printf("Error updating product: %v\n", err)
		return err
	}

	// Display updated product details
	fmt.Printf("Product updated successfully!\n")

	productID := "N/A"
	if updatedProduct.ProductID != nil {
		productID = *updatedProduct.ProductID
	}
	fmt.Printf("Product ID: %s\n", productID)

	name := "N/A"
	if updatedProduct.Name != nil {
		name = *updatedProduct.Name
	}
	fmt.Printf("Name: %s\n", name)

	description := "N/A"
	if updatedProduct.Description != nil {
		description = *updatedProduct.Description
	}
	fmt.Printf("Description: %s\n", description)

	price := 0.0
	if updatedProduct.Price != nil {
		price = *updatedProduct.Price
	}
	fmt.Printf("Price: $%.2f\n", price)

	currency := "N/A"
	if updatedProduct.Currency != nil {
		currency = *updatedProduct.Currency
	}
	fmt.Printf("Currency: %s\n", currency)

	billingType := "N/A"
	if updatedProduct.BillingType != nil {
		billingType = *updatedProduct.BillingType
	}
	fmt.Printf("Billing Type: %s\n", billingType)

	taxInclusive := false
	if updatedProduct.TaxInclusive != nil {
		taxInclusive = *updatedProduct.TaxInclusive
	}
	fmt.Printf("Tax Inclusive: %t\n", taxInclusive)

	updatedAt := "N/A"
	if updatedProduct.UpdatedAt != nil {
		updatedAt = *updatedProduct.UpdatedAt
	}
	fmt.Printf("Updated At: %s\n", updatedAt)

	return nil
}

// Example 7: Archive Product
func archiveProduct(ctx context.Context, client *bagelpay.BagelPayClient) error {
	fmt.Println("\n=== Example 7: Archive Product ===")

	// First, get a product to archive (using the first product from our list)
	response, err := client.ListProducts(ctx, 1, 1)
	if err != nil {
		fmt.Printf("Error listing products: %v\n", err)
		return err
	}

	if len(response.Items) == 0 {
		fmt.Println("No products available to archive")
		return nil
	}

	productToArchive := response.Items[0]
	if productToArchive.ProductID == nil {
		fmt.Println("Product ID is nil, cannot archive")
		return nil
	}

	// Check if product is already archived
	if productToArchive.IsArchive != nil && *productToArchive.IsArchive {
		fmt.Printf("Product %s is already archived\n", *productToArchive.ProductID)
		return nil
	}

	// Archive the product
	archivedProduct, err := client.ArchiveProduct(ctx, *productToArchive.ProductID)
	if err != nil {
		fmt.Printf("Error archiving product: %v\n", err)
		return err
	}

	// Display archived product details
	fmt.Printf("Product archived successfully!\n")

	productID := "N/A"
	if archivedProduct.ProductID != nil {
		productID = *archivedProduct.ProductID
	}
	fmt.Printf("Product ID: %s\n", productID)

	name := "N/A"
	if archivedProduct.Name != nil {
		name = *archivedProduct.Name
	}
	fmt.Printf("Name: %s\n", name)

	isArchive := false
	if archivedProduct.IsArchive != nil {
		isArchive = *archivedProduct.IsArchive
	}
	fmt.Printf("Is Archived: %t\n", isArchive)

	updatedAt := "N/A"
	if archivedProduct.UpdatedAt != nil {
		updatedAt = *archivedProduct.UpdatedAt
	}
	fmt.Printf("Updated At: %s\n", updatedAt)

	return nil
}

// Example 8: Unarchive Product
func unarchiveProduct(ctx context.Context, client *bagelpay.BagelPayClient) error {
	fmt.Println("\n=== Example 8: Unarchive Product ===")

	// First, get an archived product to unarchive
	response, err := client.ListProducts(ctx, 1, 10)
	if err != nil {
		fmt.Printf("Error listing products: %v\n", err)
		return err
	}

	if len(response.Items) == 0 {
		fmt.Println("No products available")
		return nil
	}

	// Find an archived product
	var productToUnarchive *bagelpay.Product
	for _, product := range response.Items {
		if product.IsArchive != nil && *product.IsArchive {
			productToUnarchive = &product
			break
		}
	}

	if productToUnarchive == nil {
		fmt.Println("No archived products found to unarchive")
		return nil
	}

	if productToUnarchive.ProductID == nil {
		fmt.Println("Product ID is nil, cannot unarchive")
		return nil
	}

	// Unarchive the product
	unarchivedProduct, err := client.UnarchiveProduct(ctx, *productToUnarchive.ProductID)
	if err != nil {
		fmt.Printf("Error unarchiving product: %v\n", err)
		return err
	}

	// Display unarchived product details
	fmt.Printf("Product unarchived successfully!\n")

	productID := "N/A"
	if unarchivedProduct.ProductID != nil {
		productID = *unarchivedProduct.ProductID
	}
	fmt.Printf("Product ID: %s\n", productID)

	name := "N/A"
	if unarchivedProduct.Name != nil {
		name = *unarchivedProduct.Name
	}
	fmt.Printf("Name: %s\n", name)

	isArchive := false
	if unarchivedProduct.IsArchive != nil {
		isArchive = *unarchivedProduct.IsArchive
	}
	fmt.Printf("Is Archived: %t\n", isArchive)

	updatedAt := "N/A"
	if unarchivedProduct.UpdatedAt != nil {
		updatedAt = *unarchivedProduct.UpdatedAt
	}
	fmt.Printf("Updated At: %s\n", updatedAt)

	return nil
}

// Example 9: List Subscriptions
func listSubscriptions(ctx context.Context, client *bagelpay.BagelPayClient) error {
	fmt.Println("\n=== Example 9: List Subscriptions ===")

	// List subscriptions with pagination
	response, err := client.ListSubscriptions(ctx, 1, 3)
	if err != nil {
		fmt.Printf("Error listing subscriptions: %v\n", err)
		return err
	}

	fmt.Printf("Total subscriptions: %d\n", response.Total)
	fmt.Printf("Showing %d subscriptions:\n", len(response.Items))

	if len(response.Items) == 0 {
		fmt.Println("No subscriptions found")
		return nil
	}

	for i, subscription := range response.Items {
		fmt.Printf("\n--- Subscription %d ---\n", i+1)

		subscriptionID := "N/A"
		if subscription.SubscriptionID != nil {
			subscriptionID = *subscription.SubscriptionID
		}
		fmt.Printf("Subscription ID: %s\n", subscriptionID)

		status := "N/A"
		if subscription.Status != nil {
			status = *subscription.Status
		}
		fmt.Printf("Status: %s\n", status)

		productID := "N/A"
		if subscription.ProductID != nil {
			productID = *subscription.ProductID
		}
		fmt.Printf("Product ID: %s\n", productID)

		productName := "N/A"
		if subscription.ProductName != nil {
			productName = *subscription.ProductName
		}
		fmt.Printf("Product Name: %s\n", productName)

		amount := 0.0
		if subscription.Amount != nil {
			amount = *subscription.Amount
		}
		fmt.Printf("Amount: $%.2f\n", amount)

		if subscription.Customer != nil {
			customerEmail := "N/A"
			if subscription.Customer.Email != nil {
				customerEmail = *subscription.Customer.Email
			}
			fmt.Printf("Customer Email: %s\n", customerEmail)
		}

		recurringInterval := "N/A"
		if subscription.RecurringInterval != nil {
			recurringInterval = *subscription.RecurringInterval
		}
		fmt.Printf("Recurring Interval: %s\n", recurringInterval)

		createdAt := "N/A"
		if subscription.CreatedAt != nil {
			createdAt = *subscription.CreatedAt
		}
		fmt.Printf("Created At: %s\n", createdAt)
	}

	return nil
}

// Example 10: Get Subscription Details
func getSubscriptionDetails(ctx context.Context, client *bagelpay.BagelPayClient) error {
	fmt.Println("\n=== Example 10: Get Subscription Details ===")

	// First, get a subscription ID from the list
	response, err := client.ListSubscriptions(ctx, 1, 1)
	if err != nil {
		fmt.Printf("Error listing subscriptions: %v\n", err)
		return err
	}

	if len(response.Items) == 0 {
		fmt.Println("No subscriptions available to get details for")
		return nil
	}

	subscriptionToGet := response.Items[0]
	if subscriptionToGet.SubscriptionID == nil {
		fmt.Println("Subscription ID is nil, cannot get details")
		return nil
	}

	// Get subscription details
	subscription, err := client.GetSubscription(ctx, *subscriptionToGet.SubscriptionID)
	if err != nil {
		fmt.Printf("Error getting subscription details: %v\n", err)
		return err
	}

	// Display subscription details
	fmt.Printf("Subscription details retrieved successfully!\n")

	subscriptionID := "N/A"
	if subscription.SubscriptionID != nil {
		subscriptionID = *subscription.SubscriptionID
	}
	fmt.Printf("Subscription ID: %s\n", subscriptionID)

	status := "N/A"
	if subscription.Status != nil {
		status = *subscription.Status
	}
	fmt.Printf("Status: %s\n", status)

	productID := "N/A"
	if subscription.ProductID != nil {
		productID = *subscription.ProductID
	}
	fmt.Printf("Product ID: %s\n", productID)

	productName := "N/A"
	if subscription.ProductName != nil {
		productName = *subscription.ProductName
	}
	fmt.Printf("Product Name: %s\n", productName)

	amount := 0.0
	if subscription.Amount != nil {
		amount = *subscription.Amount
	}
	fmt.Printf("Amount: $%.2f\n", amount)

	if subscription.Customer != nil {
		customerEmail := "N/A"
		if subscription.Customer.Email != nil {
			customerEmail = *subscription.Customer.Email
		}
		fmt.Printf("Customer Email: %s\n", customerEmail)
	}

	billingPeriodStart := "N/A"
	if subscription.BillingPeriodStart != nil {
		billingPeriodStart = *subscription.BillingPeriodStart
	}
	fmt.Printf("Billing Period Start: %s\n", billingPeriodStart)

	billingPeriodEnd := "N/A"
	if subscription.BillingPeriodEnd != nil {
		billingPeriodEnd = *subscription.BillingPeriodEnd
	}
	fmt.Printf("Billing Period End: %s\n", billingPeriodEnd)

	recurringInterval := "N/A"
	if subscription.RecurringInterval != nil {
		recurringInterval = *subscription.RecurringInterval
	}
	fmt.Printf("Recurring Interval: %s\n", recurringInterval)

	createdAt := "N/A"
	if subscription.CreatedAt != nil {
		createdAt = *subscription.CreatedAt
	}
	fmt.Printf("Created At: %s\n", createdAt)

	return nil
}

// Example 11: Cancel Subscription
func cancelSubscription(ctx context.Context, client *bagelpay.BagelPayClient) error {
	fmt.Println("\n=== Example 11: Cancel Subscription ===")

	// First, get an active subscription to cancel
	response, err := client.ListSubscriptions(ctx, 1, 10)
	if err != nil {
		fmt.Printf("Error listing subscriptions: %v\n", err)
		return err
	}

	if len(response.Items) == 0 {
		fmt.Println("No subscriptions available to cancel")
		return nil
	}

	// Find an active subscription
	var subscriptionToCancel *bagelpay.Subscription
	for _, subscription := range response.Items {
		if subscription.Status != nil && *subscription.Status == "active" {
			subscriptionToCancel = &subscription
			break
		}
	}

	if subscriptionToCancel == nil {
		fmt.Println("No active subscriptions found to cancel")
		return nil
	}

	if subscriptionToCancel.SubscriptionID == nil {
		fmt.Println("Subscription ID is nil, cannot cancel")
		return nil
	}

	// Cancel the subscription
	cancelledSubscription, err := client.CancelSubscription(ctx, *subscriptionToCancel.SubscriptionID)
	if err != nil {
		fmt.Printf("Error cancelling subscription: %v\n", err)
		return err
	}

	// Display cancelled subscription details
	fmt.Printf("Subscription cancelled successfully!\n")

	subscriptionID := "N/A"
	if cancelledSubscription.SubscriptionID != nil {
		subscriptionID = *cancelledSubscription.SubscriptionID
	}
	fmt.Printf("Subscription ID: %s\n", subscriptionID)

	status := "N/A"
	if cancelledSubscription.Status != nil {
		status = *cancelledSubscription.Status
	}
	fmt.Printf("Status: %s\n", status)

	productName := "N/A"
	if cancelledSubscription.ProductName != nil {
		productName = *cancelledSubscription.ProductName
	}
	fmt.Printf("Product Name: %s\n", productName)

	cancelAt := "N/A"
	if cancelledSubscription.CancelAt != nil {
		cancelAt = *cancelledSubscription.CancelAt
	}
	fmt.Printf("Cancel At: %s\n", cancelAt)

	if cancelledSubscription.Customer != nil {
		customerEmail := "N/A"
		if cancelledSubscription.Customer.Email != nil {
			customerEmail = *cancelledSubscription.Customer.Email
		}
		fmt.Printf("Customer Email: %s\n", customerEmail)
	}

	updatedAt := "N/A"
	if cancelledSubscription.UpdatedAt != nil {
		updatedAt = *cancelledSubscription.UpdatedAt
	}
	fmt.Printf("Updated At: %s\n", updatedAt)

	return nil
}
