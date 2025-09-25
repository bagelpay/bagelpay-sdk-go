// Package bagelpay provides a Go client library for the BagelPay API.
package bagelpay

import (
	"encoding/json"
)

// Customer represents customer data for checkout session
type Customer struct {
	Email string `json:"email"`
}

// CheckoutRequest represents the request model for creating a checkout session
type CheckoutRequest struct {
	ProductID  string                 `json:"product_id"`
	Customer   *Customer              `json:"customer,omitempty"`
	RequestID  *string                `json:"request_id,omitempty"`
	Units      *string                `json:"units,omitempty"`
	SuccessURL *string                `json:"success_url,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

// CheckoutResponse represents the response model for checkout session
type CheckoutResponse struct {
	Object      *string                `json:"object,omitempty"`
	Units       *int                   `json:"units,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Mode        *string                `json:"mode,omitempty"`
	PaymentID   *string                `json:"payment_id,omitempty"`
	ProductID   *string                `json:"product_id,omitempty"`
	RequestID   *string                `json:"request_id,omitempty"`
	SuccessURL  *string                `json:"success_url,omitempty"`
	CheckoutURL *string                `json:"checkout_url,omitempty"`
	CreatedAt   *string                `json:"created_at,omitempty"`
	UpdatedAt   *string                `json:"updated_at,omitempty"`
	ExpiresOn   *string                `json:"expires_on,omitempty"`
}

// CreateProductRequest represents the request model for creating a product
type CreateProductRequest struct {
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	Currency          string  `json:"currency"`
	BillingType       string  `json:"billing_type"`
	TaxInclusive      bool    `json:"tax_inclusive"`
	TaxCategory       string  `json:"tax_category"`
	RecurringInterval string  `json:"recurring_interval"`
	TrialDays         int     `json:"trial_days"`
}

// Product represents a product model
type Product struct {
	Name              *string  `json:"name,omitempty"`
	Description       *string  `json:"description,omitempty"`
	Price             *float64 `json:"price,omitempty"`
	Currency          *string  `json:"currency,omitempty"`
	Object            *string  `json:"object,omitempty"`
	Mode              *string  `json:"mode,omitempty"`
	ProductID         *string  `json:"product_id,omitempty"`
	StoreID           *string  `json:"store_id,omitempty"`
	ProductURL        *string  `json:"product_url,omitempty"`
	BillingType       *string  `json:"billing_type,omitempty"`
	BillingPeriod     *string  `json:"billing_period,omitempty"`
	TaxCategory       *string  `json:"tax_category,omitempty"`
	TaxInclusive      *bool    `json:"tax_inclusive,omitempty"`
	IsArchive         *bool    `json:"is_archive,omitempty"`
	CreatedAt         *string  `json:"created_at,omitempty"`
	UpdatedAt         *string  `json:"updated_at,omitempty"`
	TrialDays         *int     `json:"trial_days,omitempty"`
	RecurringInterval *string  `json:"recurring_interval,omitempty"`
}

// ProductListResponse represents the product list response
type ProductListResponse struct {
	Total int       `json:"total"`
	Items []Product `json:"items"`
	Code  int       `json:"code"`
	Msg   string    `json:"msg"`
}

// UpdateProductRequest represents the request model for updating a product
type UpdateProductRequest struct {
	ProductID         string  `json:"product_id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	Currency          string  `json:"currency"`
	BillingType       string  `json:"billing_type"`
	TaxInclusive      bool    `json:"tax_inclusive"`
	TaxCategory       string  `json:"tax_category"`
	RecurringInterval string  `json:"recurring_interval"`
	TrialDays         int     `json:"trial_days"`
}

// TransactionCustomer represents customer data in transaction
type TransactionCustomer struct {
	ID    *string `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
}

// Transaction represents a transaction model
type Transaction struct {
	Object         *string              `json:"object,omitempty"`
	OrderID        *string              `json:"order_id,omitempty"`
	TransactionID  *string              `json:"transaction_id,omitempty"`
	Amount         *float64             `json:"amount,omitempty"`
	AmountPaid     *float64             `json:"amount_paid,omitempty"`
	DiscountAmount *float64             `json:"discount_amount,omitempty"`
	Currency       *string              `json:"currency,omitempty"`
	TaxAmount      *float64             `json:"tax_amount,omitempty"`
	TaxCountry     *string              `json:"tax_country,omitempty"`
	RefundedAmount *float64             `json:"refunded_amount,omitempty"`
	Type           *string              `json:"type,omitempty"`
	Customer       *TransactionCustomer `json:"customer,omitempty"`
	CreatedAt      *string              `json:"created_at,omitempty"`
	UpdatedAt      *string              `json:"updated_at,omitempty"`
	Remark         *string              `json:"remark,omitempty"`
	Mode           *string              `json:"mode,omitempty"`
	Fees           *float64             `json:"fees,omitempty"`
	Tax            *float64             `json:"tax,omitempty"`
	Net            *float64             `json:"net,omitempty"`
}

// TransactionListResponse represents the transaction list response
type TransactionListResponse struct {
	Total int           `json:"total"`
	Items []Transaction `json:"items"`
	Code  int           `json:"code"`
	Msg   string        `json:"msg"`
}

// SubscriptionCustomer represents customer data in subscription
type SubscriptionCustomer struct {
	ID    *string `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
}

// Subscription represents a subscription model
type Subscription struct {
	Object             *string               `json:"object,omitempty"`
	Status             *string               `json:"status,omitempty"`
	Remark             *string               `json:"remark,omitempty"`
	Customer           *SubscriptionCustomer `json:"customer,omitempty"`
	Mode               *string               `json:"mode,omitempty"`
	Amount             *float64              `json:"amount,omitempty"`
	Last4              *string               `json:"last4,omitempty"`
	SubscriptionID     *string               `json:"subscription_id,omitempty"`
	ProductID          *string               `json:"product_id,omitempty"`
	StoreID            *string               `json:"store_id,omitempty"`
	BillingPeriodStart *string               `json:"billing_period_start,omitempty"`
	BillingPeriodEnd   *string               `json:"billing_period_end,omitempty"`
	CancelAt           *string               `json:"cancel_at,omitempty"`
	TrialStart         *string               `json:"trial_start,omitempty"`
	TrialEnd           *string               `json:"trial_end,omitempty"`
	Units              *int                  `json:"units,omitempty"`
	CreatedAt          *string               `json:"created_at,omitempty"`
	UpdatedAt          *string               `json:"updated_at,omitempty"`
	ProductName        *string               `json:"product_name,omitempty"`
	PaymentMethod      *string               `json:"payment_method,omitempty"`
	NextBillingAmount  *float64              `json:"next_billing_amount,omitempty"`
	RecurringInterval  *string               `json:"recurring_interval,omitempty"`
}

// SubscriptionListResponse represents the subscription list response
type SubscriptionListResponse struct {
	Total int            `json:"total"`
	Items []Subscription `json:"items"`
	Code  int            `json:"code"`
	Msg   string         `json:"msg"`
}

// CustomerData represents customer data model
type CustomerData struct {
	ID            *int     `json:"id,omitempty"`
	Name          *string  `json:"name,omitempty"`
	Email         *string  `json:"email,omitempty"`
	Remark        *string  `json:"remark,omitempty"`
	Subscriptions *int     `json:"subscriptions,omitempty"`
	Payments      *int     `json:"payments,omitempty"`
	StoreID       *string  `json:"store_id,omitempty"`
	TotalSpend    *float64 `json:"total_spend,omitempty"`
	CreatedAt     *string  `json:"created_at,omitempty"`
	UpdatedAt     *string  `json:"updated_at,omitempty"`
}

// CustomerListResponse represents the customer list response
type CustomerListResponse struct {
	Total int            `json:"total"`
	Items []CustomerData `json:"items"`
	Code  int            `json:"code"`
	Msg   string         `json:"msg"`
}

// APIError represents an API error response
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Error implements the error interface for APIError
func (e *APIError) Error() string {
	return e.Message
}

// APIResponse represents a generic API response wrapper
type APIResponse struct {
	Data interface{} `json:"data,omitempty"`
	*APIError
}

// Helper functions for pointer types
func StringPtr(s string) *string {
	return &s
}

func IntPtr(i int) *int {
	return &i
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func BoolPtr(b bool) *bool {
	return &b
}

// ToJSON converts a struct to JSON string
func ToJSON(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FromJSON converts JSON string to struct
func FromJSON(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}
