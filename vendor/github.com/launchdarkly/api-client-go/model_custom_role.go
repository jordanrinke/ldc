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

type CustomRole struct {
	Links *Links `json:"_links,omitempty"`
	// Name of the custom role.
	Name string `json:"name,omitempty"`
	// The 20-hexdigit id or the key for a custom role.
	Key string `json:"key,omitempty"`
	// Description of the custom role.
	Description string `json:"description,omitempty"`
	// The unique resource id.
	Id string `json:"_id,omitempty"`
	Policy []Policy `json:"policy,omitempty"`
}
