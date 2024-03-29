/*
 * Policy Service
 *
 * Service to interact with Policies and related objects.
 *
 * API version: v2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ReviewPolicyResponseModel struct for ReviewPolicyResponseModel
type ReviewPolicyResponseModel struct {
	Bundle PolicyBundleModel `json:"bundle"`
	Policy PolicyModel `json:"policy"`
}
