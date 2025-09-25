/*
BagelPay SDK - Product Management Example

This example demonstrates product management operations:
- Creating different types of products (digital, physical, subscription)
- Listing products
- Retrieving a specific product
- Updating product information
- Archiving products

Prerequisites:
- Set your API key as an environment variable: export BAGELPAY_API_KEY="your_api_key_here"
- Install the SDK: go mod tidy

To run this example:
go run product_management.go
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
	fmt.Println("🚀 BagelPay SDK - Product Management Example")
	fmt.Println("============================================")

	// Get API key from environment variable
	apiKey := os.Getenv("BAGELPAY_API_KEY")
	if apiKey == "" {
		apiKey = "bagel_test_C6D1E83B94204A00A6F8EFD2AF05B427"
		// log.Fatal("❌ BAGELPAY_API_KEY environment variable is required")
	}

	// Initialize the BagelPay client
	fmt.Println("\n📡 Initializing BagelPay client...")
	client := bagelpay.NewTestClient(apiKey)

	ctx := context.Background()

	// Example 1: Create different types of products
	fmt.Println("\n📦 Creating different types of products...")

	digitalProduct, err := createDigitalProduct(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to create digital product: %v", err)
		return
	}
	if digitalProduct != nil && digitalProduct.ProductID != nil {
		fmt.Printf("✅ Digital product created! ID: %s\n", *digitalProduct.ProductID)
	}

	subscriptionProduct, err := createSubscriptionProduct(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to create subscription product: %v", err)
		return
	}
	if subscriptionProduct != nil && subscriptionProduct.ProductID != nil {
		fmt.Printf("✅ Subscription product created! ID: %s\n", *subscriptionProduct.ProductID)
	}

	// Example 2: List all products
	fmt.Println("\n📋 Listing all products...")
	err = listAllProducts(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to list products: %v", err)
		return
	}

	// Example 3: Retrieve a specific product
	fmt.Println("\n🔍 Retrieving specific product...")
	if digitalProduct != nil && digitalProduct.ProductID != nil {
		err = getProductDetails(ctx, client, *digitalProduct.ProductID)
		if err != nil {
			log.Printf("❌ Failed to get product: %v", err)
			return
		}
	}

	// Example 4: Update product information
	fmt.Println("\n✏️ Updating product information...")
	if digitalProduct != nil && digitalProduct.ProductID != nil {
		err = updateProductInfo(ctx, client, *digitalProduct.ProductID)
		if err != nil {
			log.Printf("❌ Failed to update product: %v", err)
			return
		}
	}

	// Example 5: Archive a product
	fmt.Println("\n🗄️ Archiving product...")
	if digitalProduct != nil && digitalProduct.ProductID != nil {
		err = archiveProduct(ctx, client, *digitalProduct.ProductID)
		if err != nil {
			log.Printf("❌ Failed to archive product: %v", err)
			return
		}
	}

	fmt.Println("\n🎉 Product management examples completed successfully!")
}

// createDigitalProduct creates a digital product
func createDigitalProduct(ctx context.Context, client *bagelpay.BagelPayClient) (*bagelpay.Product, error) {
	productRequest := bagelpay.CreateProductRequest{
		Name:              "Advanced Go Programming Course.",
		Description:       "Master advanced Go programming concepts with hands-on examples",
		Price:             49.99, // $49.99
		Currency:          "USD",
		BillingType:       "single_payment",
		TaxInclusive:      false,
		TaxCategory:       "digital_products",
		RecurringInterval: "",
		TrialDays:         0,
	}

	return client.CreateProduct(ctx, productRequest)
}

// createSubscriptionProduct creates a subscription product
func createSubscriptionProduct(ctx context.Context, client *bagelpay.BagelPayClient) (*bagelpay.Product, error) {
	productRequest := bagelpay.CreateProductRequest{
		Name:              "Pro API Access.",
		Description:       "Monthly subscription for premium API features and higher rate limits",
		Price:             29.99, // $29.99
		Currency:          "USD",
		BillingType:       "subscription",
		TaxInclusive:      false,
		TaxCategory:       "digital_products",
		RecurringInterval: "monthly",
		TrialDays:         14,
	}

	return client.CreateProduct(ctx, productRequest)
}

// listAllProducts lists all products
func listAllProducts(ctx context.Context, client *bagelpay.BagelPayClient) error {
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

		description := "N/A"
		if product.Description != nil {
			description = *product.Description
		}

		fmt.Printf("   %d. %s (ID: %s)\n", i+1, name, productID)

		// Check if it's a subscription product and add recurring interval
		priceDisplay := fmt.Sprintf("%.2f %s", price, currency)
		if product.BillingType != nil && *product.BillingType == "subscription" && product.RecurringInterval != nil && *product.RecurringInterval != "" {
			priceDisplay += fmt.Sprintf(" (%s)", *product.RecurringInterval)
		}
		fmt.Printf("      Price: %s\n", priceDisplay)
		fmt.Printf("      Description: %s\n", description)
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
		priceDisplay := fmt.Sprintf("%.2f %s", *product.Price, *product.Currency)
		// Check if it's a subscription product and add recurring interval
		if product.BillingType != nil && *product.BillingType == "subscription" && product.RecurringInterval != nil && *product.RecurringInterval != "" {
			priceDisplay += fmt.Sprintf(" (%s)", *product.RecurringInterval)
		}
		fmt.Printf("   Price: %s\n", priceDisplay)
	}
	if product.IsArchive != nil {
		fmt.Printf("   Archived: %t\n", *product.IsArchive)
	}

	return nil
}

// updateProductInfo updates product information
func updateProductInfo(ctx context.Context, client *bagelpay.BagelPayClient, productID string) error {
	updateRequest := bagelpay.UpdateProductRequest{
		ProductID:         productID,
		Name:              "Advanced Go Programming Course - Updated Edition",
		Description:       "Master advanced Go programming concepts with hands-on examples. Now includes bonus modules!",
		Price:             59.99, // Updated price: $59.99
		Currency:          "USD",
		BillingType:       "single_payment",
		TaxInclusive:      false,
		TaxCategory:       "digital_products",
		RecurringInterval: "",
		TrialDays:         0,
	}

	product, err := client.UpdateProduct(ctx, updateRequest)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Product updated successfully!\n")
	if product.Name != nil {
		fmt.Printf("   New Name: %s\n", *product.Name)
	}
	if product.Price != nil && product.Currency != nil {
		fmt.Printf("   New Price: %.2f %s\n", *product.Price, *product.Currency)
	}

	return nil
}

// archiveProduct archives a product
func archiveProduct(ctx context.Context, client *bagelpay.BagelPayClient, productID string) error {
	product, err := client.ArchiveProduct(ctx, productID)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Product archived successfully! ID: %s\n", productID)
	if product.IsArchive != nil && *product.IsArchive {
		fmt.Println("   Status: Archived")
	}
	fmt.Println("   Note: Archived products are no longer available for new purchases")

	return nil
}
