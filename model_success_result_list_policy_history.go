/*
 * Policy Service
 *
 * Service to interact with Policies and related objects.
 *
 * API version: v2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SuccessResultListPolicyHistory struct for SuccessResultListPolicyHistory
type SuccessResultListPolicyHistory struct {
	Data []PolicyHistoryModel `json:"data"`
	Meta Metadata `json:"meta"`
}