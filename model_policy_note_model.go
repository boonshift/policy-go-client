/*
 * Policy Service
 *
 * Service to interact with Policies and related objects.
 *
 * API version: v2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PolicyNoteModel struct for PolicyNoteModel
type PolicyNoteModel struct {
	// Description not provided.
	CreatedBy string `json:"createdBy"`
	// Description not provided.
	Note string `json:"note"`
	// Description not provided.
	PolicyId string `json:"policyId"`
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
