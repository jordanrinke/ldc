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

type Integration struct {
	Links *DependentFlagLinks `json:"_links,omitempty"`
	Items []IntegrationSubscription `json:"items,omitempty"`
}
