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

type IntegrationSubscriptionStatus struct {
	SuccessCount int32 `json:"successCount,omitempty"`
	// A unix epoch time in milliseconds specifying the last time this integration was successfully used.
	LastSuccess int64 `json:"lastSuccess,omitempty"`
	ErrorCount int32 `json:"errorCount,omitempty"`
}
