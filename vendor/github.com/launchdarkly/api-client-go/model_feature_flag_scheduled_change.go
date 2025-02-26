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

type FeatureFlagScheduledChange struct {
	// A unix epoch time in milliseconds specifying the date the scheduled changes will be applied
	ExecutionDate int32 `json:"executionDate,omitempty"`
	Version int32 `json:"_version,omitempty"`
	Id string `json:"_id,omitempty"`
	Instructions *SemanticPatchInstruction `json:"instructions,omitempty"`
}
