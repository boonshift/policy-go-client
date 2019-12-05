/*
 * Policy Service
 *
 * Service to interact with Policies and related objects.
 *
 * API version: v2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PolicyModel struct for PolicyModel
type PolicyModel struct {
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
	Status string `json:"status"`
	// Description not provided.
	TaxAmount float32 `json:"taxAmount"`
	// The id of the entity.
	Id string `json:"id"`
	// ID of account in IAM hierarchy.
	SfAccountId int32 `json:"sfAccountId,omitempty"`
	// The timestamp of when the entity was created.
	SfCreated string `json:"sfCreated"`
	// User ID of entity owner.
	SfOwnerId string `json:"sfOwnerId,omitempty"`
	// Path of entity in IAM hierarchy.
	SfPath string `json:"sfPath,omitempty"`
	// ID of tenant that entity belongs to.
	SfTenantId int32 `json:"sfTenantId,omitempty"`
	// The timestamp of when the entity was last updated.
	SfUpdated string `json:"sfUpdated"`
	// The current version of the entity.
	SfVersion int32 `json:"sfVersion"`
}
