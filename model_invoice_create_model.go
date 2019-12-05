/*
 * Policy Service
 *
 * Service to interact with Policies and related objects.
 *
 * API version: v2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InvoiceCreateModel struct for InvoiceCreateModel
type InvoiceCreateModel struct {
	// Description not provided.
	ContinuationLink string `json:"continuationLink"`
	// Description not provided.
	PaymentReference string `json:"paymentReference"`
	// Description not provided.
	PolicyHolderId string `json:"policyHolderId"`
	// Description not provided.
	SfAccountId int32 `json:"sfAccountId"`
	// Description not provided.
	Status string `json:"status"`
	// Description not provided.
	TotalAmount float32 `json:"totalAmount"`
}