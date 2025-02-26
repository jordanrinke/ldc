/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 5.3.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type ApprovalRequestConfigBody struct {
	// A name that describes the changes you would like to apply to a feature flag configuration
	Description string `json:"description"`
	Instructions *SemanticPatchInstruction `json:"instructions"`
	NotifyMemberIds []string `json:"notifyMemberIds"`
	// comment will be included in audit log item for change.
	Comment string `json:"comment,omitempty"`
	// Timestamp for when instructions will be executed
	ExecutionDate int64 `json:"executionDate,omitempty"`
	// ID of scheduled change to edit or delete
	OperatingOnId string `json:"operatingOnId,omitempty"`
}
