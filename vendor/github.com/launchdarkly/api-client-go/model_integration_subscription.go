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

type IntegrationSubscription struct {
	Links *HierarchicalLinks `json:"_links,omitempty"`
	// The unique resource id.
	Id string `json:"_id,omitempty"`
	// The type of integration associated with this configuration.
	Kind string `json:"kind,omitempty"`
	// The user-defined name associated with this configuration.
	Name string `json:"name,omitempty"`
	// A key-value mapping of configuration fields.
	Config *interface{} `json:"config,omitempty"`
	Statements []Statement `json:"statements,omitempty"`
	// Whether or not the integration is currently active.
	On bool `json:"on,omitempty"`
	// An array of tags for this integration configuration.
	Tags []string `json:"tags,omitempty"`
	Status *IntegrationSubscriptionStatus `json:"_status,omitempty"`
}
