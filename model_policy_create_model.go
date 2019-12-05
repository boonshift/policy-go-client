/*
 * Policy Service
 *
 * Service to interact with Policies and related objects.
 *
 * API version: v2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PolicyCreateModel struct for PolicyCreateModel
type PolicyCreateModel struct {
	// Description not provided.
	Config string `json:"config"`
	// Description not provided.
	CoverageEndDate string `json:"coverageEndDate"`
	// Description not provided.
	CoverageStartDate string `json:"coverageStartDate"`
	// Description not provided.
	FullPremium float32 `json:"fullPremium"`
	// Description not provided.
	LifecycleType string `json:"lifecycleType"`
	// Description not provided.
	PayablePremium float32 `json:"payablePremium"`
	// Description not provided.
	PolicyBundleId string `json:"policyBundleId"`
	// Description not provided.
	PolicyHolderId string `json:"policyHolderId"`
	// Description not provided.
	PolicyNumber string `json:"policyNumber"`
	// Description not provided.
	ProductId string `json:"productId"`
	// Description not provided.
	RenewsPolicyId string `json:"renewsPolicyId,omitempty"`
	// Description not provided.
	SfAccountId int32 `json:"sfAccountId"`
	// Description not provided.
	Status string `json:"status"`
	// Description not provided.
	TaxAmount float32 `json:"taxAmount"`
}
