/*
 * Policy Service
 *
 * Service to interact with Policies and related objects.
 *
 * API version: v2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SuccessResultListInvoiceLine struct for SuccessResultListInvoiceLine
type SuccessResultListInvoiceLine struct {
	Data []InvoiceLineModel `json:"data"`
	Meta Metadata `json:"meta"`
}
