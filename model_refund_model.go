/*
 * Policy Service
 *
 * Service to interact with Policies and related objects.
 *
 * API version: v2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// RefundModel struct for RefundModel
type RefundModel struct {
	// Description not provided.
	Amount float32 `json:"amount,omitempty"`
	// Description not provided.
	Option string `json:"option"`
	// Description not provided.
	Tax float32 `json:"tax,omitempty"`
}
