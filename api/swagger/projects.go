/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Projects struct {

	Links *Links `json:"_links,omitempty"`

	// The unique resource id.
	Id string `json:"_id,omitempty"`

	Key string `json:"key,omitempty"`

	Name string `json:"name,omitempty"`

	Items []Project `json:"items,omitempty"`
}