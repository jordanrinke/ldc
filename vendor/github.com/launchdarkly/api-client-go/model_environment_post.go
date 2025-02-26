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

type EnvironmentPost struct {
	// The name of the new environment.
	Name string `json:"name"`
	// A project-unique key for the new environment.
	Key string `json:"key"`
	// A color swatch (as an RGB hex value with no leading '#', e.g. C8C8C8).
	Color string `json:"color"`
	// The default TTL for the new environment.
	DefaultTtl float32 `json:"defaultTtl,omitempty"`
	// Determines whether the environment is in secure mode.
	SecureMode bool `json:"secureMode,omitempty"`
	// Set to true to send detailed event information for newly created flags.
	DefaultTrackEvents bool `json:"defaultTrackEvents,omitempty"`
	// An array of tags for this environment.
	Tags []string `json:"tags,omitempty"`
	// Determines if this environment requires comments for flag and segment changes.
	RequireComments bool `json:"requireComments,omitempty"`
	// Determines if this environment requires confirmation for flag and segment changes.
	ConfirmChanges bool `json:"confirmChanges,omitempty"`
}
