/*
BagelPay SDK - Subscription & Customer Management Example

This example demonstrates subscription and customer management operations:
- Listing subscriptions
- Retrieving specific subscription details
- Canceling subscriptions
- Listing customers
- Retrieving customer details

Prerequisites:
- Set your API key as an environment variable: export BAGELPAY_API_KEY="your_api_key_here"
- Install the SDK: go mod tidy
- Have some existing subscriptions and customers in your account

To run this example:
go run subscription_customer_management.go
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
	fmt.Println("🚀 BagelPay SDK - Subscription & Customer Management Example")
	fmt.Println("===========================================================")

	// Get API key from environment variable
	apiKey := os.Getenv("BAGELPAY_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ BAGELPAY_API_KEY environment variable is required")
	}

	// Initialize the BagelPay client
	fmt.Println("\n📡 Initializing BagelPay client...")
	client := bagelpay.NewTestClient(apiKey)

	ctx := context.Background()

	// Example 1: List all subscriptions
	fmt.Println("\n🔄 Listing all subscriptions...")
	subscriptions, err := listAllSubscriptions(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to list subscriptions: %v", err)
		return
	}

	// Example 2: Get specific subscription details (if any exist)
	if len(subscriptions) > 0 && subscriptions[0].SubscriptionID != nil {
		fmt.Println("\n🔍 Getting subscription details...")
		err = getSubscriptionDetails(ctx, client, *subscriptions[0].SubscriptionID)
		if err != nil {
			log.Printf("❌ Failed to get subscription details: %v", err)
			return
		}

		// Example 3: Cancel a subscription (commented out to avoid accidental cancellation)
		fmt.Println("\n⚠️  Subscription cancellation example (not executed)...")
		fmt.Printf("   To cancel subscription %s, uncomment the cancellation code\n", *subscriptions[0].SubscriptionID)
		// err = cancelSubscription(ctx, client, *subscriptions[0].SubscriptionID)
		// if err != nil {
		// 	log.Printf("❌ Failed to cancel subscription: %v", err)
		// 	return
		// }
	}

	// Example 4: List all customers
	fmt.Println("\n👥 Listing all customers...")
	_, err = listAllCustomers(ctx, client)
	if err != nil {
		log.Printf("❌ Failed to list customers: %v", err)
		return
	}

	// Example 5: Customer details are already displayed in the list above
	// Note: Individual customer retrieval is not available in the current SDK

	fmt.Println("\n🎉 Subscription & customer management examples completed successfully!")
}

// listAllSubscriptions lists all subscriptions
func listAllSubscriptions(ctx context.Context, client *bagelpay.BagelPayClient) ([]*bagelpay.Subscription, error) {
	response, err := client.ListSubscriptions(ctx, 1, 5) // pageNum=1, pageSize=50
	if err != nil {
		return nil, err
	}

	// Convert []Subscription to []*Subscription for consistency
	var subscriptions []*bagelpay.Subscription
	for i := range response.Items {
		subscriptions = append(subscriptions, &response.Items[i])
	}

	if len(subscriptions) == 0 {
		fmt.Println("📝 No subscriptions found")
		fmt.Println("   💡 Tip: Create some subscription products and complete checkout sessions to see subscriptions here")
		return subscriptions, nil
	}

	fmt.Printf("📝 Found %d subscription(s):\n", len(subscriptions))
	for i, subscription := range subscriptions {
		if subscription.SubscriptionID != nil {
			fmt.Printf("   %d. ID: %s\n", i+1, *subscription.SubscriptionID)
		}
		if subscription.Status != nil {
			fmt.Printf("      Status: %s\n", *subscription.Status)
		}
		if subscription.Customer != nil && subscription.Customer.Email != nil {
			fmt.Printf("      Customer: %s\n", *subscription.Customer.Email)
		}
		if subscription.ProductID != nil && *subscription.ProductID != "" {
			fmt.Printf("      Product: %s\n", *subscription.ProductID)
		}
		if subscription.NextBillingAmount != nil {
			fmt.Printf("      Next Billing: $%.2f\n", *subscription.NextBillingAmount)
		}
		if subscription.CreatedAt != nil {
			fmt.Printf("      Created: %s\n", *subscription.CreatedAt)
		}
		fmt.Println()
	}

	return subscriptions, nil
}

// getSubscriptionDetails retrieves and displays subscription details
func getSubscriptionDetails(ctx context.Context, client *bagelpay.BagelPayClient, subscriptionID string) error {
	subscription, err := client.GetSubscription(ctx, subscriptionID)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Subscription Details:\n")
	if subscription.SubscriptionID != nil {
		fmt.Printf("   ID: %s\n", *subscription.SubscriptionID)
	}
	if subscription.Status != nil {
		fmt.Printf("   Status: %s\n", *subscription.Status)
	}
	if subscription.Customer != nil && subscription.Customer.Email != nil {
		fmt.Printf("   Customer Email: %s\n", *subscription.Customer.Email)
	}

	if subscription.ProductID != nil && *subscription.ProductID != "" {
		fmt.Printf("   Product ID: %s\n", *subscription.ProductID)
	}

	if subscription.NextBillingAmount != nil {
		fmt.Printf("   Next Billing Amount: $%.2f\n", *subscription.NextBillingAmount)
	}
	if subscription.CreatedAt != nil {
		fmt.Printf("   Created: %s\n", *subscription.CreatedAt)
	}

	if subscription.BillingPeriodStart != nil && *subscription.BillingPeriodStart != "" {
		fmt.Printf("   Billing Period Start: %s\n", *subscription.BillingPeriodStart)
	}

	if subscription.BillingPeriodEnd != nil && *subscription.BillingPeriodEnd != "" {
		fmt.Printf("   Billing Period End: %s\n", *subscription.BillingPeriodEnd)
	}

	if subscription.CancelAt != nil && *subscription.CancelAt != "" {
		fmt.Printf("   Cancel At: %s\n", *subscription.CancelAt)
	}

	return nil
}

// cancelSubscription cancels a subscription
func cancelSubscription(ctx context.Context, client *bagelpay.BagelPayClient, subscriptionID string) error {
	subscription, err := client.CancelSubscription(ctx, subscriptionID)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Subscription canceled successfully!\n")
	if subscription.SubscriptionID != nil {
		fmt.Printf("   ID: %s\n", *subscription.SubscriptionID)
	}
	if subscription.Status != nil {
		fmt.Printf("   Status: %s\n", *subscription.Status)
	}
	if subscription.CancelAt != nil {
		fmt.Printf("   Cancel At: %s\n", *subscription.CancelAt)
	}
	fmt.Println("   Note: The subscription will remain active until the end of the current billing period")

	return nil
}

// listAllCustomers lists all customers
func listAllCustomers(ctx context.Context, client *bagelpay.BagelPayClient) ([]*bagelpay.CustomerData, error) {
	response, err := client.ListCustomers(ctx, 1, 5) // pageNum=1, pageSize=50
	if err != nil {
		return nil, err
	}

	// Convert []CustomerData to []*CustomerData for consistency
	var customers []*bagelpay.CustomerData
	for i := range response.Items {
		customers = append(customers, &response.Items[i])
	}

	if len(customers) == 0 {
		fmt.Println("📝 No customers found")
		fmt.Println("   💡 Tip: Complete some checkout sessions to see customers here")
		return customers, nil
	}

	fmt.Printf("📝 Found %d customer(s):\n", len(customers))
	for i, customer := range customers {
		if customer.ID != nil {
			fmt.Printf("   %d. ID: %d\n", i+1, *customer.ID)
		}
		if customer.Email != nil {
			fmt.Printf("      Email: %s\n", *customer.Email)
		}

		if customer.Name != nil && *customer.Name != "" {
			fmt.Printf("      Name: %s\n", *customer.Name)
		}

		if customer.CreatedAt != nil {
			fmt.Printf("      Created: %s\n", *customer.CreatedAt)
		}
		fmt.Println()
	}

	return customers, nil
}
