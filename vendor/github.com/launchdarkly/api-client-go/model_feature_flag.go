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

type FeatureFlag struct {
	Key string `json:"key,omitempty"`
	// Name of the feature flag.
	Name string `json:"name,omitempty"`
	// Description of the feature flag.
	Description string `json:"description,omitempty"`
	// Whether the feature flag is a boolean flag or multivariate.
	Kind string `json:"kind,omitempty"`
	// A unix epoch time in milliseconds specifying the creation time of this flag.
	CreationDate int64 `json:"creationDate,omitempty"`
	IncludeInSnippet bool `json:"includeInSnippet,omitempty"`
	// Whether or not this flag is temporary.
	Temporary bool `json:"temporary,omitempty"`
	// The ID of the member that should maintain this flag.
	MaintainerId string `json:"maintainerId,omitempty"`
	// An array of tags for this feature flag.
	Tags []string `json:"tags,omitempty"`
	// The variations for this feature flag.
	Variations []Variation `json:"variations,omitempty"`
	// An array goals from all environments associated with this feature flag
	GoalIds []string `json:"goalIds,omitempty"`
	Version int32 `json:"_version,omitempty"`
	// A mapping of keys to CustomProperty entries.
	CustomProperties map[string]CustomProperty `json:"customProperties,omitempty"`
	Links *Links `json:"_links,omitempty"`
	Maintainer *Member `json:"_maintainer,omitempty"`
	Environments map[string]FeatureFlagConfig `json:"environments,omitempty"`
	// A unix epoch time in milliseconds specifying the archived time of this flag.
	ArchivedDate int64 `json:"archivedDate,omitempty"`
	// Whether or not this flag is archived.
	Archived bool `json:"archived,omitempty"`
	ClientSideAvailability *ClientSideAvailability `json:"clientSideAvailability,omitempty"`
	Defaults *Defaults `json:"defaults,omitempty"`
}
